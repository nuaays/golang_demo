package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"hash/crc32"
	"os"
)

func isGzipFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}

	defer file.Close()

	magic := make([]byte, 2)
	numbytes, err := file.Read(magic)
	if numbytes != 2 || err != nil {
		return false
	}

	return (magic[0] == 0x1f && magic[1] == 0x8b)

}

func GeneralHash(filename string, size int) uint32 {
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("Failed to open the '%s'", filename)
		return 0
	}
	defer fd.Close()

	buf := make([]byte, size)

	if isGzipFile(filename) {
		gw, err := gzip.NewReader(fd)
		if err != nil {
			panic(err)
		}
		defer gw.Close()

		for {
			n, _ := gw.Read(buf)
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

	}

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
	crc32_code := crc32.ChecksumIEEE(buf)
	return crc32_code
}

func main() {
	filename := flag.String("f", "/home/wanghe/work/VerificationCode/hash/hash.go", "file name ")
	flag.Parse()

	fmt.Println(*filename)
	crc32 := GeneralHash(*filename, 2048)
	fmt.Println(crc32)

}
