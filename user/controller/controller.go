package controller

import (
	userModel "splitwise/user/model"
	userService "splitwise/user/service"

	"github.com/kataras/iris/v12"
)

var allUsers map[int]userModel.User

func AllApis(app *iris.Application, allUsr map[int]userModel.User) {
	allUsers = allUsr
	userAPI := app.Party("/user")
	{
		userAPI.Get("/", getAllUsers)
		userAPI.Post("/", createUser)
		userAPI.Get("/{userId}", getUser)
	}
}

func getAllUsers(ctx iris.Context) {
	ctx.JSON(allUsers)
}

func createUser(ctx iris.Context) {
	var user userModel.User
	ctx.ReadJSON(&user)
	userService.CreateUser(user, allUsers)
}

func getUser(ctx iris.Context) {
	userId, _ := ctx.Params().GetInt("userId")
	result := userService.GetUser(userId, allUsers)
	ctx.JSON(result)
}
