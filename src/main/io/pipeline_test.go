package io

import (
	"fmt"
	"io"
	"testing"
	"time"
)

func TestPipeline(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		for {
			fmt.Println("write to pipe ")
			// Write to the PipeWriter blocks until it has satisfied one or more Reads from the PipeReader that fully consume the written data.
			// The data is copied directly from the Write to the corresponding Read (or Reads);
			// there is no internal buffering.
			pipeWriter.Write([]byte("hello world"))
			fmt.Println("[write] before sleep ")
			time.Sleep(time.Second * 10)
		}
	}()

	go func() {
		buf := make([]byte, 16)
		pipeReader.Read(buf)
		fmt.Println("read from writer ", string(buf))
		fmt.Println("[read] before sleep ")
		time.Sleep(time.Second * 30)
	}()

	time.Sleep(3600 * time.Second)
}