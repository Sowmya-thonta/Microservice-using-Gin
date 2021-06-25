package services

import (
	"fmt"
	"microservices/domain/httperrors"
	"microservices/domain/users"
)

var (
	UsersService          = usersService{}
	registeredUsers       = map[int64]*users.User{}
	currentUserId   int64 = 1
)

type usersService struct {
}

func (service usersService) Create(user users.User) (*users.User, *httperrors.HttpError) {

	if user.FirstName == "" {
		return nil, httperrors.NewBadRequestError("Invalid User First Name")
	}

	if user.LastName == "" {
		return nil, httperrors.NewBadRequestError("Invalid User Last Name")
	}

	user.ID = currentUserId
	currentUserId++
	registeredUsers[user.ID] = &user
	return &user, nil
}

func (service usersService) Get(userId int64) (*users.User, *httperrors.HttpError) {

	if user := registeredUsers[userId]; user != nil {
		return user, nil
	}
	return nil, httperrors.NewNotFoundError(fmt.Sprintf("User with id:%d Not found!!", userId))

}
