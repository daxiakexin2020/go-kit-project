package Services

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

//定义请求与响应的结构
type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Result string `json:"result"`
}

//定义基本结构
//response interface{}   接口类型
//request interface{}  请求
func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error){

		//断言UserRequest 此处可以定义许多请求类型，如何去传，使用Transport去定义与控制
        r:=request.(UserRequest)

        //r.Uid从请求中获取uid，接着获取响应结果
        result:=userService.GetName(r.Uid)

        //使用UserService中的方法
        return UserResponse{Result:result},nil
	}
}
