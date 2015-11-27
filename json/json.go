package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Stream struct {
	StreamId    int    `json:"stream_id"`
	SubStreamId int    `json:"sub_stream_id"`
	Hash        int    `json:"hash"`
	FileName    string `json:"file_name"`
}

type StreamTable struct {
	T []Stream `json:"table"`
}

func main() {

	stream := Stream{0, 0, 0, "/var/heka"}
	b, err := json.Marshal(stream)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	var table StreamTable
	table.T = make([]Stream, 0, 2048)
	stream = Stream{1, 1, 1, "/var/log"}
	table.T = append(table.T, stream)
	b, err = json.Marshal(table)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(b))

	fd, err := os.OpenFile("/var/heka/StreamTable", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0660)
	if err != nil {
		fmt.Println("failed to open")
	}
	_, err = fd.Write(b)
	if err != nil {
		fmt.Println("failed to write")
	}

}