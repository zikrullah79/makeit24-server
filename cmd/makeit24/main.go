package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "zikrullah79/makeit24-server/pkg/ScoreServices"
	"zikrullah79/makeit24-server/pkg/mongo"
	"zikrullah79/makeit24-server/pkg/mongo/model"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedScoreServicesServer
}

func (s *server) GetAllScore(in *pb.GetAllScoreRequest, stream pb.ScoreServices_GetAllScoreServer) error {
	d, err := mongo.GetAllScores()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range d {
		if err := stream.Send(&pb.Score{Name: v.Name, Point: fmt.Sprintf("%v", v.Point)}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) AddNewScore(ctx context.Context, in *pb.AddNewScoreRequest) (*pb.ScoreResponse, error) {
	_, err := mongo.InsertScore(model.InitScore(in.Score, in.Name))
	if err != nil {
		log.Fatal(err)
	}
	return &pb.ScoreResponse{Message: "Succesfully Add Scores"}, nil
}

var (
	NWorker = flag.Int("n", 4, "Number Of Worker")
	Address = flag.String("http", "127.0.0.1:2001", "Address and port to listen for HTTP Request")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Print(err)
	}
	grpcServer := grpc.NewServer()
	err = mongo.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterScoreServicesServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	// worker.StartDivider(*NWorker)
	// http.HandleFunc("/card", worker.Handler)
	// http.HandleFunc("/score", mongo.GetAllScoresHandler)
	// http.HandleFunc("/score/{id}", mongo.GetOneByIdHandler)
	// fmt.Printf("Server Running on %v \n", *Address)

	// i, err := mongo.InsertScores(&model.Score{Point: 50.0, Name: "udin"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(i)
	// i, err := mongo.GetAllScores()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(i)
	// if err := http.ListenAndServe(*Address, nil); err != nil {
	// 	fmt.Println(err)
	// }
}
