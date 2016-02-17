package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("Usage: md5 filename\n")
		return
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("os.Open: %s\n", err)
		return
	}
	hasher := md5.New()
	if _, err := io.Copy(hasher, file); err != nil {
		fmt.Errorf("io.Copy: %s\n", err)
		return
	}

	sum := hasher.Sum(nil)

	sumString := hex.EncodeToString(sum)

	fmt.Println(sumString)
}
