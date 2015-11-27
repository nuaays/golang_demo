package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
	"net"
	//	"os"
)

const (
	MAX_RECV_MESSAGE_BYTES = 4 * 1024 * 1024
)

func DoMessage(outBytes []byte) {
	if outBytes != nil {
		if outBytes != nil {
			out.data = append(out.data, outBytes...)
			out.cursor = pack.QueueCursor
			msgCounter++
		}

		if len(pack.MsgBytes) >= 0 {
			msgCounter++
			msgCounter2++
			fmt.Println("msg counter = ", msgCounter)
			if msgCounter2 == 5 {
				b, _ := json.Marshal(o.CheckpointInfo)
				o.WriteCheckPointInfo(b)
				msgCounter2 = uint32(0)
			}
			if err = block.UnMarshal(pack.MsgBytes); err != nil {
				fmt.Println(err)
			}
			// block.SeeBigAllString(block.GetEventGroupS(), block.GetEventGroupCount())
			groupsbytes = pack.MsgBytes[20:]
			groupsbytesLen = uint32(len(groupsbytes))
			compress_block_len = block.GetBlockSize()
			total_size += int(compress_block_len)
			//区块索引
			go func() {
				eventgroup_count = block.GetEventGroupCount()
				host_id = block.GetHostId()
				receive_timestamp = block.GetTimestamp()
				row_index = fmt.Sprintf("%d,%d,%d,%d,%d\n", offset, offset_compress, eventgroup_count, host_id, receive_timestamp)
				fmt.Println(row_index)
				o.rowChan <- row_index
			}()
			//tokenserver数据
			go func() {
				block.TokenBlockOffset = uint64(offset)
				block.TokenBlockSize = uint32(28) + groupsbytesLen
				tokenBytes, e = block.MarshalTokenServer()
				o.tokenChan <- tokenBytes
			}()
			offset += uint64(len(groupsbytes))            //未压缩的偏移, 为用tokenserver提供的全局偏移
			offset_compress += uint64(compress_block_len) //压缩后的偏移,记录本地文件位置

		}
	}
}


func WriteBlockIndex(row chan string) {
	for {
		select {
		case ch := <-row:
			fmt.Println("ch = ", ch)
			if _, err := o.indexfile.WriteString(ch); err != nil {
				fmt.Println("write index file error")
			}
		}
	}
}

// 当程序启动时从文件系统中读取上一次记录的文件偏移位置
// 这个文件记录在一个默认名为checkpoint.info的文件中
func  ReadOffsetFromFile(path string) error {
	var (
		fd  *os.File
		err error
	)

	if fd, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		fmt.Println("Failed to open checkpoint.info")
		return err
	}
	defer fd.Close()

	buf := make([]byte, 1024)
	if _, err := fd.Read(buf); err != nil {
		fmt.Println("err: Failed to read checkpoint.info cotent")
		return err
	}

	if err = json.Unmarshal(buf, o.CheckpointInfo); err != nil {
		fmt.Println("Failed to Unmarshal checkpointInfo: ", err)
		return err
	}

	//o.CheckpointInfo.Offset, err = strconv.ParseUint(string(buf), 10, 64)
	buf = buf[:0]
	return nil

}

// 写入CheckpointInfo
func WriteCheckPointInfo(b []byte) error {
	var (
		fd  *os.File
		err error
	)
	if fd, err = os.OpenFile(o.StorageOutputConfig.CheckPointInfoPath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666); err != nil {
		fmt.Println("Failed to write checkpoint information.")
		return err
	}

	defer fd.Close()
	if _, err = fd.Write(b); err != nil {
		return err
	}

	return nil
}

//Runs in a separate goroutine, waits for buffered data on the committer
// channel, writes it out to the filesystem, and puts the now empty buffer on
// the return channel for reuse.
func  committer() {
	initBatch := newOutBatch()
	o.backChan <- initBatch
	var out *outBatch
	var err error

	ok := true
	hupChan := make(chan interface{})
	notify.Start(RELOAD, hupChan)
	var fileInfo os.FileInfo
	for ok {
		select {
		case out, ok = <-o.batchChan:
			if !ok {
				// Channel is closed => we're shutting down, exit cleanly.
				o.file.Close()
				close(o.closing)
				break
			}
			var e error
			fileInfo, e = os.Stat(o.path)
			if e == nil {
				o.CurrentFileSize = fileInfo.Size()
			}
			if o.CurrentFileSize >= o.StorageOutputConfig.FileMaxSize {
				o.file.Close()
				t := gostrftime.Format("%Y-%m-%d_%H%M%S", time.Now())
				filename := fmt.Sprint(o.StorageOutputConfig.NamePattern, "-", t, o.StorageOutputConfig.Suffix)
				path := filepath.Join(o.StorageOutputConfig.LogDir, filename)
				o.file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, o.perm)
				if err != nil {
					or.LogError(err)
					o.file.Close()
				}
				o.indexfile.Close()
				filename = fmt.Sprint(o.StorageOutputConfig.NamePattern, "-", t, o.StorageOutputConfig.SuffixIndex)
				o.indexpath = filepath.Join(o.StorageOutputConfig.LogDir, filename)
				o.indexfile, err = os.OpenFile(o.indexpath, os.O_APPEND|os.O_CREATE|os.O_RDWR, o.perm)
				if err != nil {
//					or.LogError(err)
					o.indexfile.Close()
				}
				o.path = path
				o.CurrentFileSize = int64(0)
			}

			n, err := file.Write(out.data)

			if err != nil {
//				or.LogError(fmt.Errorf("Can't write to %s: %s", o.path, err))
			} else if n != len(out.data) {
//				or.LogError(fmt.Errorf("data loss - truncated output for %s", o.path))
//				or.UpdateCursor(out.cursor)
			} else {
//				o.file.Sync()
//				or.UpdateCursor(out.cursor)
			}

			out.data = out.data[:0]
			o.backChan <- out
		case <-hupChan:
			o.file.Close()
			if err = o.openFile(); err != nil {
				close(o.closing)
				err = fmt.Errorf("unable to reopen file '%s': %s", o.path, err)
				errChan <- err
				ok = false
				break
			}
		case rotateTime := <-o.rotateChan:
			o.file.Close()
			o.path = gostrftime.Strftime(o.StorageOutputConfig.Path, rotateTime)
			if err = o.openFile(); err != nil {
				close(o.closing)
				err = fmt.Errorf("unable to open rotated file '%s': %s", o.path, err)
				errChan <- err
				ok = false
				break
			}
		}
	}
}


type ClientTokenServer struct {
	Address string
	Ok      bool
	Stop    chan bool
	Conn    net.Conn
	err     error
}

func NewClientToken(address string) *ClientTokenServer {
	if len(address) == 0 {
		fmt.Println("Invalid Ip address for token server")
		return nil
	}
	return &ClientTokenServer{Address: address}
}

func (self *ClientTokenServer) Run(tokenChan chan []byte) {
	if self.Conn, self.err = net.Dial("tcp", self.Address); self.err != nil {
		self.Ok = false
		fmt.Println("Failed to connect:  ", self.err.Error())
		return
	}
	fmt.Println("Connected to structure server ...")

	OK := true
	for OK {
		select {
		case ch := <-tokenChan:
			n, err := self.Conn.Write(ch)
			if err != nil {
				fmt.Printf("n = %d:%s\n", n, err)
			} else {
				ch = nil
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, MAX_RECV_MESSAGE_BYTES)
	go o.committer(or, errChan)
	rowChan = make(chan string, 50)
	go o.WriteBlockIndex(o.rowChan)
	o.tokenChan = make(chan []byte, 100)
	o.ClientTokenServer.Address = o.StorageOutputConfig.Address
	go o.ClientTokenServer.Run(o.tokenChan)
	return o.receiver(or, errChan)
	for {
		_, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		DoMessage(buf)

	}
}
func main() {
	ln, err := net.Listen("tcp", ":5565")
	if err != nil {
		fmt.Println("failed to listen ")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
