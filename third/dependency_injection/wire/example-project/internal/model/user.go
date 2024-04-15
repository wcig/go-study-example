package model

type CreateUserReq struct {
	Name string `json:"name"`
}

type CreateUserRes struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}
