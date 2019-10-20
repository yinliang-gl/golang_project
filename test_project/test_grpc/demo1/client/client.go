package main

import (
	"golang.org/x/net/context"
	"golang_project/test_project/test_grpc/demo1/test_proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	waiterClient := test_proto.NewWaiterClient(conn)

	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}

	// 调用gRPC接口
	tr, err := waiterClient.DoMD5(context.Background(), &test_proto.Req{JsonStr: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.BackJson)
}
