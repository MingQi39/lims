package dto

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	AccessToken string `json:"access_token"`
	UserID      uint64 `json:"user_id"`
	UserName    string `json:"user_name"`
}
