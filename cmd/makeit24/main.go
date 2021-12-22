package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	pb "zikrullah79/makeit24-server/pkg/ScoreServices"
	"zikrullah79/makeit24-server/pkg/mongo"
	"zikrullah79/makeit24-server/pkg/mongo/model"
	"zikrullah79/makeit24-server/pkg/variable"
	"zikrullah79/makeit24-server/pkg/worker"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedScoreServicesServer
}

func (s *server) GetAllScore(in *pb.GetAllScoreRequest, stream pb.ScoreServices_GetAllScoreServer) error {
	res := make(chan interface{})
	req := &worker.WorkRequest{WorkType: variable.AllScores, Result: res}

	worker.WorkQueue <- *req

	result := <-req.Result
	close(req.Result)
	if req.Error != nil {
		return req.Error
	}

	for _, v := range result.([]model.Score) {
		if err := stream.Send(&pb.Score{Name: v.Name, Point: v.Point}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) AddNewScore(ctx context.Context, in *pb.AddNewScoreRequest) (*pb.ScoreResponse, error) {
	if in.Name == "" {
		return nil, errors.New("BAD REQUEST")
	}
	res := make(chan interface{})
	req := &worker.WorkRequest{WorkType: variable.PostScore, Data: model.InitScore(in.Score, in.Name), Result: res}

	worker.WorkQueue <- *req

	result := <-req.Result
	close(req.Result)
	if req.Error != nil {
		return nil, req.Error
	}

	return &pb.ScoreResponse{Message: fmt.Sprintf("Succesfully Add Scores, id : %v", result)}, nil
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
	worker.StartDivider(*NWorker)
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
