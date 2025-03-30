package commands

import (
	"log"

	"github.com/mattermost/mattermost-server/v6/model"
)

func SendMessage(client *model.Client4, channelID, message string) {
	post := &model.Post{
		ChannelId: channelID,
		Message:   message,
	}

	if _, _, err := client.CreatePost(post); err != nil {
		log.Printf("[ERROR] Ошибка при отправке сообщения: %v", err)
	}
}

func sendError(client *model.Client4, channelID, msg string) {
	log.Printf("[ERROR] %s", msg)
	SendMessage(client, channelID, "❌ "+msg)
}

func sendSuccess(client *model.Client4, channelID, msg string) {
	log.Printf("[SUCCESS] %s", msg)
	SendMessage(client, channelID, msg)
}
