package main

import (
	"bufio"
	"ekyu.moe/leb128"
	"encoding/binary"
	"errors"
	"io"
	"log"
)

func ReadLong(buffer io.Reader) (uint64, error) {
	integer := uint64(0)
	e := binary.Read(buffer, binary.LittleEndian, &integer)
	return integer, e
}

func ReadInteger(buffer io.Reader) (uint32, error) {
	integer := uint32(0)
	e := binary.Read(buffer, binary.LittleEndian, &integer)
	return integer, e
}

func ReadShort(buffer io.Reader) (uint16, error) {
	short := uint16(0)
	e := binary.Read(buffer, binary.LittleEndian, &short)
	return short, e
}

func ReadBoolean(buffer *bufio.Reader) (bool, error) {
	next, e := buffer.ReadByte()
	if e != nil {
		return false, e
	}
	return next == 0x1, nil
}

// Get the value of a variable length integer
func ReadUleb(reader *bufio.Reader) (uint64, error) {
	next, e := reader.Peek(10)
	if e != nil {
		return 0, e
	}
	value, length := leb128.DecodeUleb128(next)
	_, e = reader.Discard(int(length))
	if e != nil {
		return 0, e
	}
	return value, nil
}

// Get the string value from bufio.Reader
func ReadString(reader *bufio.Reader) (string, error) {
	b, e := reader.ReadByte()
	if e != nil {
		return "", e
	}
	if b == 0x00 {
		return "", nil
	}
	if b != 0x0b {
		return "", errors.New("could not find string")
	}
	length, e := ReadUleb(reader)
	if e != nil {
		return "", e
	}
	valueArray := make([]byte, length)
	_, e = reader.Read(valueArray)
	if e != nil {
		return "", e
	}
	return string(valueArray), nil
}

// Log an error if it's not nil
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
