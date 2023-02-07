package tgbotapi

func NewForwardThread(chatID int64, fromChatID int64, messageID int, threadID int) ForwardConfig {
	return ForwardConfig{
		BaseChat:   BaseChat{ChatID: chatID, MessageThreadID: threadID},
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

func NewCopyMessageThread(chatID int64, fromChatID int64, messageID int, threadID int) CopyMessageConfig {
	return CopyMessageConfig{
		BaseChat:   BaseChat{ChatID: chatID, MessageThreadID: threadID},
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

func NewMessageThread(chatID int64, text string, threadID int) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
			MessageThreadID:  threadID,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}
}
