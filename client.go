package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting client")

	const ServerKey = "vl6wu0A@XP?}Or/&BR#LSxn>A+}L)p44/W[wXL3<"
	const Path = "/kv/foo" //TODO sending long keys fails !!
	const Payload = "{\"name\":\"dave\", \"age\":30}"
	const ReqEndpoint = "tcp://127.0.0.1:5555"
	const toc = "dasdasdas"

	zestC := ZestClient{}
	zestC.Connect(ReqEndpoint, ServerKey)
	err := zestC.Post(ReqEndpoint, toc, Path, Payload)
	if err != nil {
		log(err.Error())
	}
}