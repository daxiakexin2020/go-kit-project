package main

import (
	httptansport "github.com/go-kit/kit/transport/http"
	."go-kit-project/Services"
	"net/http"
)
func main()  {

	//创建ednPoint
	user:= UserService{}
	endPoint:=GenUserEndpoint(user)

	//服务
	serverHandler:=httptansport.NewServer(endPoint,DecodeUserRequest,EncodeUserResponse)

	http.ListenAndServe(":8080",serverHandler)

}
