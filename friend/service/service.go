package service

import (
	friendModel "splitwise/friend/model"
	userModel "splitwise/user/model"
)

func AllFriendsList(allFriends map[int][]friendModel.Friend, allUsers map[int]userModel.User) map[string][]userModel.User {
	result := make(map[string][]userModel.User)
	for index, friends := range allFriends {
		var userFriends []userModel.User
		for _, eachFriend := range friends {
			userFriends = append(userFriends, allUsers[eachFriend.UserId])
		}
		result[allUsers[index].UserName] = userFriends
	}
	return result
}

func MakeFriend(friends []friendModel.Friend, allFriends map[int][]friendModel.Friend) {
	allFriends[friends[0].UserId] = append(allFriends[friends[0].UserId], friends[1])
	allFriends[friends[1].UserId] = append(allFriends[friends[1].UserId], friends[0])
}

func RemoveFriend(friends []friendModel.Friend, allFriends map[int][]friendModel.Friend) {

	userId1 := friends[0].UserId
	userId2 := friends[1].UserId
	friendsOfUser1 := []friendModel.Friend{}
	friendsOfUser2 := []friendModel.Friend{}
	for _, value := range allFriends[userId1] {
		if value == friends[1] {
			continue
		} else {
			friendsOfUser1 = append(friendsOfUser1, value)
		}
	}

	for _, value := range allFriends[userId2] {
		if value == friends[0] {
			continue
		} else {
			friendsOfUser2 = append(friendsOfUser2, value)
		}
	}
	allFriends[userId1] = friendsOfUser1
	allFriends[userId2] = friendsOfUser2
}
