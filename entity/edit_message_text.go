package entity

type EditMessageText struct {
	ChatId                uint64      `json:"chat_id"`
	MessageId             uint64      `json:"message_id"`
	InlineMessageId       *string     `json:"inline_message_id,omitempty"`
	Text                  string      `json:"text"`
	ParseMode             *string     `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool       `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup"`
}

type editMessageTextBuilder struct {
	chatId                uint64
	messageId             uint64
	inlineMessageId       *string
	text                  string
	parseMode             *string
	disableWebPagePreview *bool
	replyMarkup           ReplyMarkup
}

func NewEditMessageBuilder() *editMessageTextBuilder {
	inlineKeyboard := make([][]inlineKeyboard, 0)
	rm := ReplyMarkup{
		InlineKeyboard: &inlineKeyboard,
	}

	return &editMessageTextBuilder{
		replyMarkup: rm,
	}
}

func (emtb *editMessageTextBuilder) Text(text string) *editMessageTextBuilder {
	emtb.text = text
	return emtb
}

func (emtb *editMessageTextBuilder) ChatId(chatId uint64) *editMessageTextBuilder {
	emtb.chatId = chatId
	return emtb
}

func (emtb *editMessageTextBuilder) MessageId(messageId uint64) *editMessageTextBuilder {
	emtb.messageId = messageId
	return emtb
}

func (emtb *editMessageTextBuilder) ParseMode(parseMode string) *editMessageTextBuilder {
	*emtb.parseMode = parseMode
	return emtb
}

func (emtb *editMessageTextBuilder) DisableWebPagePreview(disableWebPagePreview bool) *editMessageTextBuilder {
	*emtb.disableWebPagePreview = disableWebPagePreview
	return emtb
}

func (emtb *editMessageTextBuilder) InlineMessageId(inlineMessageId string) *editMessageTextBuilder {
	*emtb.inlineMessageId = inlineMessageId
	return emtb
}

func (emtb *editMessageTextBuilder) AddButton(ik inlineKeyboard) *editMessageTextBuilder {
	ln := len(*(emtb.replyMarkup).InlineKeyboard) - 1

	if ln == -1 {
		*(emtb.replyMarkup).InlineKeyboard = append(*(emtb.replyMarkup).InlineKeyboard, make([]inlineKeyboard, 0))
		ln = 0
	}

	(*(emtb.replyMarkup).InlineKeyboard)[ln] = append((*(emtb.replyMarkup).InlineKeyboard)[ln], ik)

	return emtb
}

func (emtb *editMessageTextBuilder) AddButtonWithAllParams(text string, url *string, callbackData *string,
	switchInlineQuery *string, switchInlineQueryCurrentChat *string, cg *callbackGame) *editMessageTextBuilder {
	inlKey := inlineKeyboard{
		Text:                         text,
		URL:                          url,
		CallbackData:                 callbackData,
		SwitchInlineQuery:            switchInlineQuery,
		SwitchInlineQueryCurrentChat: switchInlineQueryCurrentChat,
		CallbackGame:                 cg,
	}

	return emtb.AddButton(inlKey)
}

func (emtb *editMessageTextBuilder) AddButtonUrl(text string, url string) *editMessageTextBuilder {
	inlKey := inlineKeyboard{
		Text: text,
		URL:  &url,
	}

	return emtb.AddButton(inlKey)
}

func (emtb *editMessageTextBuilder) AddButtonCallBack(text string, callbackData string) *editMessageTextBuilder {
	inlKey := inlineKeyboard{
		Text:         text,
		CallbackData: &callbackData,
	}

	return emtb.AddButton(inlKey)
}

func (emtb *editMessageTextBuilder) AddButtonLine() *editMessageTextBuilder {

	*(emtb.replyMarkup).InlineKeyboard = append(*(emtb.replyMarkup).InlineKeyboard, make([]inlineKeyboard, 0))

	return emtb
}

func (emtb *editMessageTextBuilder) Build() EditMessageText {
	var emt EditMessageText

	emt.ChatId = emtb.chatId
	emt.MessageId = emtb.messageId
	emt.Text = emtb.text
	emt.ReplyMarkup = emtb.replyMarkup

	if emtb.parseMode != nil {
		emt.ParseMode = emtb.parseMode
	}

	if emtb.parseMode != nil {
		emt.ParseMode = emtb.parseMode
	}
	if emtb.disableWebPagePreview != nil {
		emt.DisableWebPagePreview = emtb.disableWebPagePreview

	}
	if emtb.inlineMessageId != nil {
		emt.InlineMessageId = emtb.inlineMessageId
	}
	//if emtb.replyMarkup != nil {
	//	emt.ReplyMarkup = emtb.replyMarkup
	//}

	return emt

}
