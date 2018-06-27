package client

import (
	"google.golang.org/grpc"
)

var DataSync *grpc.ClientConn

func InitGrpcClient(url string) {
	DataSync, _ = grpc.Dial(url, grpc.WithInsecure())
}
