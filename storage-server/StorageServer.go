package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const KB = 1024
const MB = 1024 * 1024

var debug = false

type StorageServerConfig struct {
	
}
type StorageServer struct {
	conn        net.Conn
	ln          net.Listener
	address     string
	protocol    string
	BaseDir     string
	bw          *bufio.Writer
	fd          *os.File
	Logfile     string
	Duaration   int
}

func NewStorageServer(proto string, addr string, base_dir string) *StorageServer {
	return &StorageServer{
		protocol:  proto,
		address:   addr,
		BaseDir:   base_dir,
		Logfile:   "logfile",
		Duaration: 100,
	}
}

func (self *StorageServer) Run() {
	var (
		err error
	)
	self.ln, err = net.Listen(self.protocol, self.address)
	if err != nil {
		fmt.Println("Failed to listen : ", self.address)
		return
	}
	fmt.Println("Start listening ", self.address)
	 
	defer self.fd.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	 
	ok := true
	for ok {
		self.conn, err = self.ln.Accept()
		if err != nil {
			fmt.Println("conn error: ", err)
			continue
		}
		go self.handleConnection()
	}
}

func (self *StorageServer) handleConnection() error {
	fmt.Println("WriterServer Handle connection")
	max_size := uint64(30)
	keepAlivePeroid := uint64(30)
	keepAliveDuration := time.Duration(keepAlivePeroid) * time.Second
	defer self.conn.Close()
	input := make(chan interface{}, max_size)
	output := make(chan []interface{}, 1)
	
	ok := true
	for ok {
		buf := make([]byte, 4)
		_, err := self.conn.Read(buf[0:])
		if err != nil {
			fmt.Println("error: ", err)
			return err
		}
		buf_len := binary.BigEndian.Uint32(buf[0:])
		if buf_len == 0 {
			break
		}
		content_buf := make([]byte, buf_len-4)
		_, err = self.conn.Read(content_buf[0:])
		if err != nil {
			fmt.Println("error: ", err)
			return err
		}
		fmt.Println(content_buf)
	}
	return nil
}
