package controller

import (
	friendModel "splitwise/friend/model"
	groupModel "splitwise/group/model"
	groupService "splitwise/group/service"
	userModel "splitwise/user/model"

	"github.com/kataras/iris/v12"
)

var allUsers map[int]userModel.User
var allFriends map[int][]friendModel.Friend
var allGroups map[int][]groupModel.Group

func AllApis(app *iris.Application, allGroups1 map[int][]groupModel.Group,
	allFriends1 map[int][]friendModel.Friend,
	allUsers1 map[int]userModel.User) {
	//
	allUsers = allUsers1
	allFriends = allFriends1
	allGroups = allGroups1
	groupAPI := app.Party("/group")
	{
		groupAPI.Get("/", allGroupList)
		groupAPI.Post("/createGroup", createGroup)
		// delete group
		// add member
		// remove member
	}
}

func allGroupList(ctx iris.Context) {
	result := groupService.GetAllGroupsList(allGroups, allUsers)
	ctx.JSON(result)
}

func createGroup(ctx iris.Context) {
	var groups groupModel.Group
	ctx.ReadJSON(&groups)
	groupService.CreateGroup(groups, allGroups)
}
