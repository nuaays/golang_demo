package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//	"net"
)

func main() {
	//	conn, err := net.Dial("tcp", ":12800")
	//	if err != nil {
	//		fmt.Errorf("%s", err.Error())
	//		//		return
	//	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint64(0x1234567887654321))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.Bytes())

	//	_, err = conn.Write(buf.Bytes())
	//	if err != nil {
	//		fmt.Println("erroror ")
	//	}

}
