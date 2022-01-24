package dto

import "addnewer.com/godemo/internel/service/dto"

type UserParams struct {
	ID   uint64 `form:"id" binding:"required"`
	Name string `form:"name"`
}

func (p *UserParams) ServiceUserListParam() *dto.UserListParams {
	return &dto.UserListParams{
		ID:   p.ID,
		Name: p.Name,
	}
}

type UserResponse struct {
	ID    uint64
	Name  string
	Email string
}
