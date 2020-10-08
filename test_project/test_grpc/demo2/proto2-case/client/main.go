package main

import (
	"context"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"github.com/golang/protobuf/proto"
	"golang_project/test_project/test_grpc/demo2/proto2-case/pb"
	"google.golang.org/grpc"
	"io"
)

func main() {
	conn, e := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	// say hello
	r, e := c.SayHello(context.Background(), &pb.HelloRequest{Username: proto.String("ft")})
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(*r.Message)

	// chat
	chatClient, e := c.Chat(context.Background())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	//go func(){
	//	time.Sleep(10 * time.Second)
	//	chatClient.CloseSend()
	//}()
	go func() {
		for {
			stream, e := chatClient.Recv()
			if e == io.EOF {
				fmt.Println("EOF")
				return
			}
			if e != nil {
				fmt.Println(errorx.Wrap(e).Error())
				return
			}
			fmt.Println("receive from server:", stream.Stream)
		}
	}()
	chatClient.Send(&pb.ClientStream{
		Stream: newBytes(10, 9, 8, 7),
	})
	select {
	//case <-time.After(20 * time.Second):
	}
}

func newBytes(a ...byte) []byte {
	return a
}
