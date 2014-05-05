// +build ignore

package main

import (
	"fmt"
	"github.com/vbatts/d2r/sum"
	"os"
)

func main() {
	jsonfile := os.Args[1]
	tarfile := os.Args[2]
	tar_fh, err := os.Open(tarfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json_fh, err := os.Open(jsonfile)
	if err != nil {
		return "", err
	}
	defer json_fh.Close()

	sum, err := SumTarLayer(tar_fh, json_fh)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(sum)
}
