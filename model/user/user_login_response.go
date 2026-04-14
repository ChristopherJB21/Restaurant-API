package user

import "github.com/google/uuid"

type UserLoginResponse struct {
	IDUser   uuid.UUID `json:"IDUser"`
	UserName string    `json:"UserName"`
	Token    string    `json:"Token"`
}

func ToUserLoginResponse(item User, token string) UserLoginResponse {
	return UserLoginResponse{
		IDUser:   item.ID,
		UserName: item.UserName,
		Token:    token,
	}
}
