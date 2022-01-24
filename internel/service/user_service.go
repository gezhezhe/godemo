package service

import "addnewer.com/godemo/internel/service/dto"

type UserService interface {
	UserList(dto.UserListParams) *dto.UserList
}

type userService struct {
	UserService
}

func (s *userService) UserList(params dto.UserListParams) *dto.UserList {
	return &dto.UserList{}
}
