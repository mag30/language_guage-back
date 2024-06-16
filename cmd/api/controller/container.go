package controller

import "github.com/mag30/project-backend/cmd/service"

type Container struct {
	AuthController         *AuthController
	UserController         *UserController
	EntranceTestController *EntranceTestController
	TestController         *TestController
}

func NewContainer(
	authService *service.AuthService,
	userService *service.UserService,
	entranceTestService *service.EntranceTestService,
	testService *service.TestService) *Container {
	return &Container{
		AuthController:         NewAuthController(authService),
		UserController:         NewUserController(userService),
		EntranceTestController: NewEntranceTestController(entranceTestService),
		TestController:         NewTestController(testService),
	}
}
