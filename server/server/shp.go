/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012-2015
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Mike Trinkala (trink@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

// Extensions to make Message more useable in our current code outside the scope
// of protocol buffers.  See message.pb.go for the actually message definition.

/*

Internal message representation.

*/
//package EventGroupPacket

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

const (
	EventGroupProtoLen = uint32(28)
	BlockProtoLen      = uint32(20)
	//BlockProtoLen      = uint32(20 + 8)
	//MaxBlock = uint32(2097152) //2M
	MaxBlock = uint32(2 * 2097152) //4M
	//defualt false
	//DEBUG = false
	DEBUG = true
)

type Event struct {
	len     uint32
	payload []byte
}

func (e *Event) GetPayload() []byte {
	return e.payload
}

type EventGroup struct {
	eventgroup_size  uint32 // 字节大小
	streamid         uint32
	offset_in_origin uint64 // 该 event group 在原始日志文件的偏移量
	agent_time       uint64 // 该 event_group 在 agent 所在机器读取到时，agent 读到的系统时间的时间戳
	event_num        uint32
	eventsBytes      []byte
	events           []Event
}

func (eg *EventGroup) SetEventsBytes(data []byte) {
	eg.eventsBytes = data
}

func (eg *EventGroup) GetEventGroupSize() uint32 {
	return eg.eventgroup_size
}

func (eg *EventGroup) SetEventGroupSize(eventgroup_size uint32) {
	if eventgroup_size != 0 {
		eg.eventgroup_size = eventgroup_size
	}
}

func (eg *EventGroup) GetStreamId() uint32 {
	return eg.streamid
}

func (eg *EventGroup) SetStreamId(streamid uint32) {
	if streamid != 0 {
		eg.streamid = streamid
	}
}

func (eg *EventGroup) GetOffsetInOrigin() uint64 {
	return eg.offset_in_origin
}

func (eg *EventGroup) SetOffsetInOrigin(offset_in_origin uint64) {
	if offset_in_origin != 0 {
		eg.offset_in_origin = offset_in_origin
	}
}

func (eg *EventGroup) GetAgentTime() uint64 {
	return eg.agent_time
}

func (eg *EventGroup) SetAgentTime(agent_time uint64) {
	if agent_time != 0 {
		eg.agent_time = agent_time
	}
}

func (eg *EventGroup) GetEventNum() uint32 {
	return eg.event_num
}

func (eg *EventGroup) SetEventNum(event_num uint32) {
	if event_num != 0 {
		eg.event_num = event_num
	}
}

func (eg *EventGroup) GetEventS() []Event {
	return eg.events
}

func (eg *EventGroup) AddEvent(payload []byte) {
	if len(payload) != 0 {
		e := Event{
			len:     uint32(len(payload)),
			payload: payload,
		}
		eg.events = append(eg.events, e)
		eg.event_num += 1
		eg.SetEventGroupSize(eg.GetEventGroupSize() + uint32(len(payload)) + 4)
	}
}

func (eg *EventGroup) ReSet() {
	eg.eventgroup_size = EventGroupProtoLen
	eg.streamid = 0
	eg.offset_in_origin = 0
	eg.agent_time = uint64(time.Now().UnixNano())
	eg.event_num = 0
	eg.events = nil
	eg.eventsBytes = nil
}

func NewEventGroup() *EventGroup {
	return &EventGroup{
		eventgroup_size:  EventGroupProtoLen,
		streamid:         0,
		offset_in_origin: 0,
		agent_time:       uint64(time.Now().UnixNano()),
		event_num:        0,
		events:           make([]Event, 0, 64),
	}
}

func (eg *EventGroup) Marshal() (b []byte, err error) {
	if eg == nil {
		fmt.Printf("Invalid EventGroup Object")
		return nil, nil
	}

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, eg.eventgroup_size)
	if err != nil {
		log.Fatal("write eg.eventgroup_size error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, eg.streamid)
	if err != nil {
		log.Fatal("write eg.streamid error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, eg.offset_in_origin)
	if err != nil {
		log.Fatal("write eg.offset_in_origin error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, eg.agent_time)
	if err != nil {
		log.Fatal("write eg.agent_time error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, eg.event_num)
	if err != nil {
		log.Fatal("write eg.event_num  error: ", err)
		return nil, nil
	}

	if eg.eventsBytes != nil {
		buf.Write(eg.eventsBytes)
	}

	for _, v := range eg.events {
		err = binary.Write(buf, binary.BigEndian, v.len)
		if err != nil {
			log.Fatal("write eg.body field len error: ", err)
			return nil, nil
		}

		buf.Write(v.payload)
	}

	return buf.Bytes(), nil
}

func (eg *EventGroup) UnMarshal(data []byte) error {
	data_len := len(data)
	if data_len < 0 {
		err := fmt.Errorf("Unmarshal data is invalid")
		return err
	}

	if eg == nil {
		err := fmt.Errorf("Invalid EventGroup object")
		return err
	}

	var offset uint32 = 0
	eg.eventgroup_size = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	if uint32(data_len) < eg.eventgroup_size {
		return fmt.Errorf("EventGroup UnMarshal error")
	}
	eg.streamid = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	eg.offset_in_origin = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8

	eg.agent_time = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8

	eg.event_num = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	eg.events = make([]Event, 0, eg.event_num)
	var v Event
	for i := 0; i < int(eg.event_num); i++ {
		v.len = binary.BigEndian.Uint32(data[offset : offset+4])
		offset += 4
		v.payload = data[offset : offset+v.len]
		offset += v.len

		eg.events = append(eg.events, v)
	}

	return nil
}

type Block struct {
	blocksize        uint32 // block 的总长度(block_head+gzip（groups）后数据长度)
	offset           uint64
	uncompass_len    uint64 //groups gzip前的长度
	host_id          uint32
	timestamp        uint64 // block 创建的时间戳
	eventgroup_count uint32
	groups           []byte //此处数据为 gzip eventgroup[]的2进制

	//tokenblock
	TokenBlockOffset uint64
	TokenBlockSize   uint32
	//groups []EventGroup // 此处数据 gzip
}

func NewBlock() *Block {
	return &Block{
		blocksize:        BlockProtoLen,
		offset:           uint64(0),
		uncompass_len:    0,
		host_id:          0,
		timestamp:        uint64(time.Now().UnixNano()),
		eventgroup_count: 0,
		groups:           make([]byte, 0, MaxBlock),
		//groups: make([]byte, 0),
	}
}

func (b *Block) ReSet() {
	b.blocksize = BlockProtoLen
	b.uncompass_len = 0
	b.host_id = 0
	b.timestamp = uint64(time.Now().UnixNano())
	b.eventgroup_count = 0
	b.groups = nil
	//b.groups = make([]byte, 0, MaxBlock)
}

func (b *Block) AddEventGroup(eg []byte) {
	if len(eg) != 0 {
		b.groups = append(b.groups, eg...)
		b.blocksize += uint32(len(eg))
		b.eventgroup_count += 1
	}
}

func (b *Block) GetBlockSize() uint32 {
	return b.blocksize
}

func (b *Block) SetBlockSize(blocksize uint32) {
	if blocksize != 0 {
		b.blocksize = blocksize
	}
}

func (b *Block) GetHostId() uint32 {
	return b.host_id
}

func (b *Block) SetHostId(host_id uint32) {
	if host_id != 0 {
		b.host_id = host_id
	}
}
func (b *Block) SetOffset(offset uint64) {
	if offset != 0 {
		b.offset = offset
	}
}

func (b *Block) GetTimestamp() uint64 {
	return b.timestamp
}

func (b *Block) SetTimestamp(timestamp uint64) {
	if timestamp != 0 {
		b.timestamp = timestamp
	}
}

func (b *Block) GetEventGroupCount() uint32 {
	return b.eventgroup_count
}

func (b *Block) SetEventGroupCount(eventgroup_count uint32) {
	if eventgroup_count != 0 {
		b.eventgroup_count = eventgroup_count
	}
}

func (b *Block) GetEventGroupS() []byte {
	return b.groups
}

func (b *Block) Marshal() (bb []byte, err error) {
	if b == nil {
		fmt.Printf("Invalid CommandPacket Object")
		return nil, nil
	}

	//gzip
	var gz_data []byte
	if gz_data, err = GzipEncode(b.groups); err != nil {
		return nil, err
	}

	b.blocksize = uint32(len(gz_data)) + BlockProtoLen
	b.uncompass_len = uint64(len(b.groups))

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, b.blocksize)
	if err != nil {
		log.Fatal("write b.host_len error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.host_id)
	if err != nil {
		log.Fatal("write b.host_len error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.timestamp)
	if err != nil {
		log.Fatal("write b.host_len error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.eventgroup_count)
	if err != nil {
		log.Fatal("write b.EGPack_Num error: ", err)
		return nil, nil
	}

	//gzip
	buf.Write(gz_data)
	return buf.Bytes(), nil
}

func (b *Block) SetTokenBlockOffset(offset uint64) {
	b.TokenBlockOffset = offset
}

func (b *Block) SetTokenBlockSize(size uint32) {
	b.TokenBlockSize = size
}

// token marshal : |4B           |8B       |   4B       |           8B     |           4B                 |...
// token marshal : block size | offset | host id | timestamp | eventgroup count | groups bytes
func (b *Block) MarshalTokenServer() ([]byte, error) {
	if b == nil {
		fmt.Printf("Invalid CommandPacket Object")
		return nil, nil
	}

	var err error

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, b.TokenBlockSize)
	if err != nil {
		log.Fatal("write b.host_len error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.TokenBlockOffset)
	if err != nil {
		log.Fatal("write b.offset error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.host_id)
	if err != nil {
		log.Fatal("write b.host_len error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.timestamp)
	if err != nil {
		log.Fatal("write b.host_len error: ", err)
		return nil, nil
	}

	err = binary.Write(buf, binary.BigEndian, b.eventgroup_count)
	if err != nil {
		log.Fatal("write b.EGPack_Num error: ", err)
		return nil, nil
	}

	buf.Write(b.groups)
	return buf.Bytes(), nil
}

func (b *Block) UnMarshalTokenServer(data []byte) error {
	data_len := len(data)
	if data_len < 0 {
		err := fmt.Errorf("Unmarshal data is invalid")
		return err
	}

	if b == nil {
		err := fmt.Errorf("Invalid EventGroupPacket object")
		return err
	}

	var offset uint32 = 0
	b.blocksize = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	b.offset = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8

	b.host_id = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	b.timestamp = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8

	b.eventgroup_count = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	fmt.Println(b.eventgroup_count)
	//ungzip
	b.groups = b.groups[offset:]
	return nil

}

func (p *Block) SeeBigAllString(groups []byte, eventgroup_count uint32) {
	var offset uint32 = 0
	v := NewEventGroup()
	fmt.Println("see allstring")
	for i := uint32(0); i < eventgroup_count; i++ {
		v.UnMarshal(groups[offset:])
		offset += v.GetEventGroupSize()
		fmt.Println("=========eventgroup_count: " + strconv.Itoa(int(eventgroup_count)) + "__" + strconv.Itoa(int(i)) + "===v.GetEventS Num:" + strconv.Itoa(len(v.GetEventS())))
		for _, k := range v.GetEventS() {
			fmt.Println(string(k.GetPayload()))
		}
	}
}

func (p *Block) SeeTokenBytes(groups []byte, eventgroup_count uint32) {
	var offset uint32 = 0
	v := NewEventGroup()
	fmt.Println("see allstring")
	for i := uint32(0); i < eventgroup_count; i++ {
		v.UnMarshal(groups[offset:])
		offset += v.GetEventGroupSize()
		fmt.Println("=========eventgroup_count: " + strconv.Itoa(int(eventgroup_count)) + "__" + strconv.Itoa(int(i)) + "===v.GetEventS Num:" + strconv.Itoa(len(v.GetEventS())))
		for _, k := range v.GetEventS() {
			fmt.Println(string(k.GetPayload()))
		}
	}
}

func (b *Block) UnMarshal(data []byte) error {
	data_len := len(data)
	if data_len < 0 {
		err := fmt.Errorf("Unmarshal data is invalid")
		return err
	}

	if b == nil {
		err := fmt.Errorf("Invalid EventGroupPacket object")
		return err
	}

	var offset uint32 = 0
	b.blocksize = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	//	b.uncompass_len = binary.BigEndian.Uint64(data[offset : offset+8])
	//	offset += 8

	b.host_id = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	b.timestamp = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8

	b.eventgroup_count = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4

	//ungzip
	b.groups = b.groups[:0]
	if ungz_data, err := GzipDecode(data[offset:]); err == nil {
		b.groups = ungz_data
		return nil
	} else {
		return err
	}
}

func GzipEncode(in []byte) ([]byte, error) {
	var (
		buffer bytes.Buffer
		out    []byte
		err    error
	)
	writer := gzip.NewWriter(&buffer)
	_, err = writer.Write(in)
	if err != nil {
		writer.Close()
		return out, err
	}
	err = writer.Close()
	if err != nil {
		return out, err
	}

	return buffer.Bytes(), nil
}

func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		var out []byte
		return out, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
