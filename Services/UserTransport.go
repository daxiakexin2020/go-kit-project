package Services

import (
	"context"
	"net/http"
	"strconv"
	"errors"
	"encoding/json"
)


//对外部的请求进行转码，判断使用哪个Request进行处理
func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {

	//开始判断参数
	//例如；类似Url的形式访问的  http://127.0.0.1:xx/?uid=101
	if r.URL.Query().Get("uid") != "" {

		//拿到请求参数
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))

		//TODO 决定使用哪个Request进行处理 ,实际开发，肯定会定义许多个Request
		return UserRequest{
			Uid: uid,
		}, nil
	}

	return nil, errors.New("参数错误")

}

//编码响应，判断使用哪个response进行响应
func EncodeUserResponse(ctx context.Context,w http.ResponseWriter,response interface{}) error {
	 return json.NewEncoder(w).Encode(response)
}
