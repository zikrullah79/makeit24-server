package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"zikrullah79/makeit24-server/pkg/mongo"
	"zikrullah79/makeit24-server/pkg/worker"

	"github.com/joho/godotenv"
)

var (
	NWorker = flag.Int("n", 4, "Number Of Worker")
	Address = flag.String("http", "127.0.0.1:2001", "Address and port to listen for HTTP Request")
)

func main() {
	flag.Parse()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(os.Getenv("MONGO_HOST"))
	worker.StartDivider(*NWorker)

	http.HandleFunc("/card", worker.Handler)
	fmt.Printf("Server Running on %v \n", *Address)
	err = mongo.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	// i, err := mongo.InsertScores(&model.Score{Point: 50.0, Name: "udin"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(i)
	i, err := mongo.GetAllScores()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
	if err := http.ListenAndServe(*Address, nil); err != nil {
		fmt.Println(err)
	}
}
