package mock

type LoginUser struct {
	ID   int
	Name string
}

var (
	CurrentUser = LoginUser{
		ID:   100001,
		Name: "客户00001号",
	}
)
