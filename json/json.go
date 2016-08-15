package main

import (
	"encoding/json"
	"fmt"
	"time"
//	"os"
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

type Meta struct {
	Count      uint64 `json: "count"`
	Start      uint64 `json: "start"`
	TotalCount uint64 `json: "total_count"`
}

type Body struct {
	Content string `json: "content"`
	Meta    `json: "meta"`
}

type TestJSON struct {
				CreatedAt time.Time `json: "created_at"`
				UpdatedAt time.Time `json: "updated_at"`
}
func main() {
/*
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
*/

	var body Body
	body.Meta.Count = 10
	body.Meta.Start = 0
	body.Meta.TotalCount = 100

	body.Content = "hello world "
	b, _ := json.Marshal(body)
	fmt.Println(string(b))


	js := TestJSON{CreatedAt: time.Now(), UpdatedAt: time.Now()}

	b, _ = json.Marshal(js)
	fmt.Println(string(b))

	var a TestJSON
	json.Unmarshal(b, &a)
	fmt.Println(b)

//	StreamId    int    `json:"stream_id"`
//	SubStreamId int    `json:"sub_stream_id"`
//	Hash        int    `json:"hash"`
//	FileName    string `json:"file_name"`

stream := Stream{StreamId: 0, SubStreamId: 1, Hash: 3, FileName: "ttttttt"}
b, _ = json.Marshal(stream)
fmt.Println(string(b))


}
