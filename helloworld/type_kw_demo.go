package main

import "fmt"

// 定义一个 string 变量类型
type ss string

type (
	Client interface {
		request()
	}
)

type GrpcClient struct {
}

func (grpc GrpcClient) request() {
	fmt.Println("request to ....")
}

func main() {
	var s ss

	s = "hello world"
	fmt.Println(s)

	var client Client

	client = GrpcClient{}

	client.request()
}
