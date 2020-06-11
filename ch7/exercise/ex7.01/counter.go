package main

import (
	"bufio"
	"bytes"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanWords)
	var count int
	for s.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	var count int
	for s.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {

}
