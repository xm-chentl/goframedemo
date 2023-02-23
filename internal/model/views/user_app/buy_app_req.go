package userapp

type BuyAppReq struct {
	AppID    int
	AppName  string
	Duration int64
	UserID   int
	UserName string
}
