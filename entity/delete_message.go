package entity

type DeleteMessage struct {
	ChatId    uint64 `json:"chat_id"`
	MessageId uint64 `json:"message_id"`
}

type deleteMessageBuilder struct {
	chatId    uint64 `json:"chat_id"`
	messageID uint64 `json:"message_id"`
}

func NewDeleteMessageBuilder() *deleteMessageBuilder {
	return &deleteMessageBuilder{}
}

func (dmb *deleteMessageBuilder) ChatId(chatId uint64) *deleteMessageBuilder {
	dmb.chatId = chatId
	return dmb
}

func (dmb *deleteMessageBuilder) MessageId(messageID uint64) *deleteMessageBuilder {
	dmb.messageID = messageID
	return dmb
}

func (dmb *deleteMessageBuilder) Build() DeleteMessage {
	var dm DeleteMessage
	dm.ChatId = dmb.chatId
	dm.MessageId = dmb.messageID
	return dm
}
