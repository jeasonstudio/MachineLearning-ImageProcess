package main

import (
	"MachineLearning-ImageProcess/minstFront"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("./datasheet/train-images.idx3-ubyte", os.O_RDONLY, os.ModePerm)
	if err != nil {
		defer file.Close()
		os.Exit(0)
	}
	file.Seek(8, 0)
	fmt.Println("Success Open File")
	var buffer bytes.Buffer
	io.CopyN(&buffer, file, 8)
	_bytes := buffer.Bytes()
	var magic []byte = []byte{0x12, 0x31, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22}
	for _, byte := range _bytes {
		fmt.Printf("%02X ", byte)
	}
	fmt.Println()
	if bytes.Compare(magic, _bytes) == 0 {
		fmt.Println("Equal")
	}

	a := minstFront.MinstLabelArr{}
	fmt.Println(a)
}
