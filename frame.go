package gows

type Frame struct {
	FIN      byte
	RSV      [3]byte
	Opcode   byte
	Mask     byte
	Length   uint64
	Mask_key uint32
	Payload  []byte
}

func NewFrame() (frame *Frame) {

	return
}

func Pack(data []byte) (buf []byte) {
	buf = make([]byte, 2)
	fin := 0
	rsv := []byte{0, 0, 0}
	opcode := TEXT
	buf[0] = byte(fin << 7)
	for i, v := range rsv {
		buf[0] |= byte(v << byte(6-i))
	}
	buf[0] |= byte(opcode)
	var mask byte = 1 // check if sender is server or client
	buf[1] = mask << 7
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

	if mask == 1 {
		buf = append(buf, 0, 0, 0, 0)
		for i := 0; i < 4; i++ {
			buf[idx+i] = 0xff // mask_key here
		}
		idx += 4
	}

	buf = append(buf, data...)
	return
}

func Parse(buf []byte) (frame *Frame) {
	frame.FIN = buf[0] & 0x80
	frame.RSV[0], frame.RSV[1], frame.RSV[2] = buf[0]&0x40, buf[0]&0x20, buf[0]&0x10
	frame.Opcode = buf[0] & 0x0f
	frame.Mask = buf[1] & 0x80
	frame.Length = uint64(buf[1] & 0x7f)
	idx := 2
	if frame.Length == 126 {
		frame.Length = uint64(byte(buf[2]<<8) | buf[3])
		idx += 2
	} else if frame.Length == 127 {
		frame.Length = uint64(byte(buf[2]<<56) | byte(buf[3]<<48) |
			byte(buf[4]<<40) | byte(buf[5]<<32) | byte(buf[6]<<24) |
			byte(buf[7]<<16) | byte(buf[8]<<8) | buf[9])
		idx += 8
	}
	if frame.Mask == 1 {
		frame.Mask_key = uint32(byte(buf[idx]<<24) | byte(buf[idx]<<16) |
			byte(buf[idx]<<8) | buf[idx])
		idx += 4
	}
	frame.Payload = make([]byte, frame.Length/8)
	for i := 0; uint64(i) < frame.Length/8; i++ {
		frame.Payload[i] = buf[idx+i]
	}
	if frame.Mask == 1 {
		for i, v := range frame.Payload {
			frame.Payload[i] = v ^ byte(frame.Mask_key>>byte(3-(i%4)*8))
		}
	}

	return
}
