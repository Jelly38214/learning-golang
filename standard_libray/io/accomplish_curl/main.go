package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(os.Args[2])

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	// 使用MultiWriter做桥接,当向dest写入时,同时写入os.Stdout和file
	dest := io.MultiWriter(os.Stdout, file)

	// 读出响应的内容,并写到两个目的地
	io.Copy(dest, r.Body)

	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}

}
