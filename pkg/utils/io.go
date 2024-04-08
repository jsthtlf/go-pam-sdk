package utils

import (
	"encoding/binary"
	"io"
)

func ReadBytes(len int, r io.Reader) ([]byte, error) {
	b := make([]byte, len)
	length, err := io.ReadFull(r, b)
	if err != nil {
		return nil, err
	}
	return b[:length], err
}

func ReadByte(r io.Reader) (byte, error) {
	b, err := ReadBytes(1, r)
	if err != nil {
		return 0, err
	}
	return b[0], err
}

func ReadUint8(r io.Reader) (uint8, error) {
	b, err := ReadBytes(1, r)
	if err != nil {
		return 0, err
	}
	return uint8(b[0]), err
}

func ReadUint16LE(r io.Reader) (uint16, error) {
	b := make([]byte, 2)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return 0, nil
	}
	return binary.LittleEndian.Uint16(b), nil
}

func ReadUint16BE(r io.Reader) (uint16, error) {
	b := make([]byte, 2)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return 0, nil
	}
	return binary.BigEndian.Uint16(b), nil
}

func ReadUint32LE(r io.Reader) (uint32, error) {
	b := make([]byte, 4)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return 0, nil
	}
	return binary.LittleEndian.Uint32(b), nil
}

func ReadUint32BE(r io.Reader) (uint32, error) {
	b := make([]byte, 4)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return 0, nil
	}
	return binary.BigEndian.Uint32(b), nil
}

func WriteByte(data byte, w io.Writer) (int, error) {
	b := make([]byte, 1)
	b[0] = data
	return w.Write(b)
}

func WriteBytes(data []byte, w io.Writer) (int, error) {
	return w.Write(data)
}

func WriteUint8(data uint8, w io.Writer) (int, error) {
	b := make([]byte, 1)
	b[0] = data
	return w.Write(b)
}

func WriteUint16BE(data uint16, w io.Writer) (int, error) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, data)
	return w.Write(b)
}

func WriteUint16LE(data uint16, w io.Writer) (int, error) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, data)
	return w.Write(b)
}

func WriteUint32LE(data uint32, w io.Writer) (int, error) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, data)
	return w.Write(b)
}

func WriteUint32BE(data uint32, w io.Writer) (int, error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, data)
	return w.Write(b)
}

func PutUint16BE(data uint16) (uint8, uint8) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, data)
	return b[0], b[1]
}

func Uint16BE(d0, d1 uint8) uint16 {
	b := make([]byte, 2)
	b[0] = d0
	b[1] = d1

	return binary.BigEndian.Uint16(b)
}

func PutUint16LE(data uint16) (uint8, uint8) {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, data)
	return b[0], b[1]
}

func Uint16LE(d0, d1 uint8) uint16 {
	b := make([]byte, 2)
	b[0] = d0
	b[1] = d1

	return binary.LittleEndian.Uint16(b)
}
