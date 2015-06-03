package gows

import (
	"math/rand"
)

type Frame struct {
	FIN      byte
	RSV      [3]byte
	opc      Opcode
	Mask     byte
	Length   uint64
	Mask_key [4]byte
	Payload  []byte
}

func NewFrame() (frame *Frame) {
	frame = &Frame{0, [3]byte{0, 0, 0}, CONTINUE, 0, 0, [4]byte{0, 0, 0, 0}, []byte{}}
	return
}

func Pack(data []byte, opc Opcode, isClient bool) (buf []byte) {
	buf = make([]byte, 2)
	fin := 1
	rsv := []byte{0, 0, 0}
	buf[0] = byte(fin << 7)
	for i, v := range rsv {
		buf[0] |= byte(v << byte(6-i))
	}
	buf[0] |= byte(opc)
	if isClient {
		buf[1] |= 0x80
	}
	datalen := len(data)

	var idx int
	if datalen >= 126 {
		if datalen <= 0xffff {
			buf[1] |= 126
			buf = append(buf, 0, 0)
			for i := 0; i < 2; i++ {
				buf[2+i] = byte(datalen>>byte(1-i)*8) & 0xff
			}
			idx = 3
		} else {
			buf[1] |= 126
			buf = append(buf, 0, 0, 0, 0, 0, 0, 0, 0)
			for i := 0; i < 8; i++ {
				buf[2+i] = byte(datalen>>byte(7-i)*8) & 0xff
			}
			idx = 9
		}
	} else {
		buf[1] |= byte(datalen)
		idx = 2
	}

	if isClient {
		key := rand.Uint32()
		mask_key := []byte{byte(key >> 24), byte(key >> 16), byte(key >> 8), byte(key)}
		buf = append(buf, mask_key...)
		for i, v := range data {
			data[i] = v ^ mask_key[i%4]
		}
		idx += 4
	}

	buf = append(buf, data...)
	return
}

func Parse(conn *Connection) (frame *Frame, err error) {
	frame = NewFrame()
	buf, err := conn.Read(2)
	frame.FIN = (buf[0] & 0x80) >> 7
	frame.RSV[0], frame.RSV[1], frame.RSV[2] = buf[0]&0x40, buf[0]&0x20, buf[0]&0x10
	frame.opc = Opcode(buf[0] & 0x0f)
	frame.Mask = (buf[1] & 0x80) >> 7
	frame.Length = uint64(buf[1] & 0x7f)
	var ext []byte
	if frame.Length == 126 {
		ext, err = conn.Read(2)
		frame.Length = uint64(byte(ext[0]<<8) | ext[1])
	} else if frame.Length == 127 {
		ext, err = conn.Read(8)
		frame.Length = uint64(byte(ext[0]<<56) | byte(ext[1]<<48) |
			byte(ext[2]<<40) | byte(ext[3]<<32) | byte(ext[4]<<24) |
			byte(ext[5]<<16) | byte(ext[6]<<8) | ext[7])
	}
	if frame.Mask == 1 {
		mask, _ := conn.Read(4)
		frame.Mask_key = [4]byte{mask[0], mask[1], mask[2], mask[3]}
	}
	frame.Payload, err = conn.Read(uint32(frame.Length))
	if frame.Mask == 1 {
		for i, v := range frame.Payload {
			frame.Payload[i] = v ^ frame.Mask_key[i%4]
		}
	}
	return
}
