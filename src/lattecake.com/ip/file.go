package main

import (
	"os"
	"log"
)

func main() {
	fileName := "/Users/LatteCake/Downloads/general.yirendai.com_access.log"
	fl, err := os.Open(fileName)
	if err != nil {
		log.Println(fileName, err)
		return
	}

	defer fl.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}

}
