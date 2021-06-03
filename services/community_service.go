package services

import (
	"Dream/models"
	"fmt"
)

func AddComment(c models.Chat) bool {
	err := c.AddComment()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func DeleteComment(id string) bool {
	err := models.DeleteComment(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetCommentsByDream(id string) ([]models.Chat, bool) {
	chats, err := models.GetCommentsByDream(id)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	var chats2 []models.Chat
	for _, chat := range chats {
		chat.MainNickname = models.GetNickname(chat.MainId)
		chat.SendNickname = models.GetNickname(chat.SendId)
		chats2 = append(chats2, chat)
	}
	return chats2, true
}
