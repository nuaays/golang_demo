package main

import (
	"fmt"
	"hash/crc32"
	"os"
)

func GeneralHash(filename string, size int) uint32 {
	if len(filename) <= 0 {
		fmt.Errorf("Invalid file")
		return 0
	}
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("Failed to open the '%s'", filename)
		return 0
	}

	defer fd.Close()

	buf := make([]byte, size)
	for {
		n, _ := fd.Read(buf)
		if 0 == n {
			break
		}
		if size == n {
			break
		}

		if n < size {
			break
		}
	}
	fmt.Println(string(buf))
	crc32_code := crc32.ChecksumIEEE(buf)
	return crc32_code
}
