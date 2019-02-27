package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buff := make([]byte, size)

	n, err := f.Read(buff)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buff[0:n]

}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("<buffer size> <filename>")
		return
	}

	buffSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	file := os.Args[2]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for {
		readData := readSize(f, buffSize)
		if readData != nil {
			fmt.Print(string(readData))
		} else {
			break
		}
	}
}
