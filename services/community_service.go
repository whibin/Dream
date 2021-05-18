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
	dreams, err := models.GetCommentsByDream(id)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return dreams, true
}
