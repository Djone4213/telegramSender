package entity

type EditMessageReplyMarkup struct {
	BusinessConnectionId *string     `json:"business_connection_id,omitempty"`
	ChatId               *uint64     `json:"chat_id,omitempty"`
	MessageId            *uint64     `json:"message_id,omitempty"`
	InlineMessageId      *string     `json:"inline_message_id,omitempty"`
	ReplyMarkup          ReplyMarkup `json:"reply_markup"`
}

type editMessageReplyMarkupBuilder struct {
	businessConnectionId *string
	chatId               *uint64
	messageId            *uint64
	inlineMessageId      *string
	replyMarkup          ReplyMarkup
}

func NewEditMessageReplyMarkupBuilder() *editMessageReplyMarkupBuilder {
	inlineKeyboard := make([][]inlineKeyboard, 0)
	rm := ReplyMarkup{
		InlineKeyboard: &inlineKeyboard,
	}

	return &editMessageReplyMarkupBuilder{
		replyMarkup: rm,
	}
}

func (emrmb *editMessageReplyMarkupBuilder) BusinessConnectionId(businessConnectionId string) *editMessageReplyMarkupBuilder {
	emrmb.businessConnectionId = &businessConnectionId
	return emrmb
}

func (emrmb *editMessageReplyMarkupBuilder) ChatId(chatId uint64) *editMessageReplyMarkupBuilder {
	emrmb.chatId = &chatId
	return emrmb
}

func (emrmb *editMessageReplyMarkupBuilder) MessageId(messageId uint64) *editMessageReplyMarkupBuilder {
	emrmb.messageId = &messageId
	return emrmb
}

func (emrmb *editMessageReplyMarkupBuilder) InlineMessageId(inlineMessageId string) *editMessageReplyMarkupBuilder {
	emrmb.inlineMessageId = &inlineMessageId
	return emrmb
}

func (emrmb *editMessageReplyMarkupBuilder) AddButton(ik inlineKeyboard) *editMessageReplyMarkupBuilder {
	ln := len(*(emrmb.replyMarkup).InlineKeyboard) - 1

	if ln == -1 {
		*(emrmb.replyMarkup).InlineKeyboard = append(*(emrmb.replyMarkup).InlineKeyboard, make([]inlineKeyboard, 0))
		ln = 0
	}

	(*(emrmb.replyMarkup).InlineKeyboard)[ln] = append((*(emrmb.replyMarkup).InlineKeyboard)[ln], ik)

	return emrmb
}

func (emrmb *editMessageReplyMarkupBuilder) AddButtonWithAllParams(text string, url *string, callbackData *string,
	switchInlineQuery *string, switchInlineQueryCurrentChat *string, cg *callbackGame) *editMessageReplyMarkupBuilder {
	inlKey := inlineKeyboard{
		Text:                         text,
		URL:                          url,
		CallbackData:                 callbackData,
		SwitchInlineQuery:            switchInlineQuery,
		SwitchInlineQueryCurrentChat: switchInlineQueryCurrentChat,
		CallbackGame:                 cg,
	}

	return emrmb.AddButton(inlKey)
}

func (emrmb *editMessageReplyMarkupBuilder) AddButtonUrl(text string, url string) *editMessageReplyMarkupBuilder {
	inlKey := inlineKeyboard{
		Text: text,
		URL:  &url,
	}

	return emrmb.AddButton(inlKey)
}

func (emrmb *editMessageReplyMarkupBuilder) AddButtonCallBack(text string, callbackData string) *editMessageReplyMarkupBuilder {
	inlKey := inlineKeyboard{
		Text:         text,
		CallbackData: &callbackData,
	}

	return emrmb.AddButton(inlKey)
}

func (emtb *editMessageReplyMarkupBuilder) AddButtonLine() *editMessageReplyMarkupBuilder {

	*(emtb.replyMarkup).InlineKeyboard = append(*(emtb.replyMarkup).InlineKeyboard, make([]inlineKeyboard, 0))

	return emtb
}

func (emrmb *editMessageReplyMarkupBuilder) Build() EditMessageReplyMarkup {
	emrm := EditMessageReplyMarkup{}

	if emrmb.businessConnectionId != nil {
		emrm.BusinessConnectionId = emrmb.businessConnectionId
	}

	if emrmb.chatId != nil {
		emrm.ChatId = emrmb.chatId
	}

	if emrmb.messageId != nil {
		emrm.MessageId = emrmb.messageId
	}

	if emrmb.inlineMessageId != nil {
		emrm.InlineMessageId = emrmb.inlineMessageId
	}

	emrm.ReplyMarkup = emrmb.replyMarkup

	return emrm
}
