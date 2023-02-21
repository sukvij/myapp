package service

import (
	groupModel "splitwise/group/model"
	userModel "splitwise/user/model"
)

func GetAllGroupsList(allGroups map[int][]groupModel.Group, allUsers map[int]userModel.User) map[int][]groupModel.Group {
	return allGroups
}

func CreateGroup(group groupModel.Group, allGroups map[int][]groupModel.Group) {
	allGroups[group.GroupAdmin] = append(allGroups[group.GroupAdmin], group)
}
