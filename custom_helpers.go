package tgbotapi

import (
	"time"
)

// DeleteAfter deletes message in background after duration
func (m *Message) DeleteAfter(bot *BotAPI, duration time.Duration) {
	go func(msg *Message) {
		time.Sleep(duration)
		bot.Request(DeleteMessageConfig{MessageID: msg.MessageID, ChatID: msg.Chat.ID})
	}(m)
}

// DeleteMessage deletes message
func (bot *BotAPI) DeleteMessage(chatID int64, messageID int) (*APIResponse, error) {
	return bot.Request(DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	})
}
