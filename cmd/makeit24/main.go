package main

import (
	"flag"
	"fmt"
	"net/http"
	"zikrullah79/makeit24-server/pkg/worker"
)

var (
	NWorker = flag.Int("n", 4, "Number Of Worker")
	Address = flag.String("http", "127.0.0.1:2001", "Address and port to listen for HTTP Request")
)

func main() {
	flag.Parse()
	worker.StartDivider(*NWorker)

	http.HandleFunc("/card", worker.Handler)
	fmt.Printf("Server Running on %v", *Address)
	if err := http.ListenAndServe(*Address, nil); err != nil {
		fmt.Println(err)
	}
}
