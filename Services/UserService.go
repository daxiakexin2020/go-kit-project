package Services

//定义接口
type IUserService interface {
	GetName(userid int) string
}

//具体实现类
type UserService struct {
}

//具体实现方法
func (this UserService) GetName(userid int) string {

	if userid == 101 {
		return "zhangzhou"
	}
	return "guest"
}
