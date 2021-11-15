package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Process loads environment in given file and calls consul
func Process(filename string) error {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := strings.TrimSpace(string(bytes))
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		index := strings.Index(line, "=")
		if index < 0 {
			return fmt.Errorf("bad environment line: '%s'", line)
		}
		key := strings.TrimSpace(line[:index])
		value := strings.TrimSpace(line[index+1:])
		_, err = exec.Command("/opt/bin/consul", "kv", "put", key, value).CombinedOutput()
		if err != nil {
			return fmt.Errorf("calling consul: %v", err)
		}
	}
	return nil
}

func main() {
	for _, filename := range os.Args[1:] {
		if err := Process(filename); err != nil {
			fmt.Printf("Error calling consul: %v\n", err)
			os.Exit(1)
		}
	}
}
