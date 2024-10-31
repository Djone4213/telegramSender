package entity

type callbackGame struct {
}

type MessageText struct {
	ChatId                uint64      `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             *string     `json:"parse_mode,omitempty"`
	DisableWebPagePreview *bool       `json:"disable_web_page_preview,omitempty"`
	DisableNotification   *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageId      *uint64     `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type messageTextBuilder struct {
	chatId                uint64
	text                  string
	parseMode             *string
	disableWebPagePreview *bool
	disableNotification   *bool
	replyToMessageId      *uint64
	replyMarkup           interface{}
}

func NewMessageBuilder() *messageTextBuilder {
	return &messageTextBuilder{}
}

func (mtb *messageTextBuilder) InitInlineKeyboard() *messageTextBuilder {
	inlineKeyboard := make([][]inlineKeyboard, 0)
	rm := ReplyMarkup{
		InlineKeyboard: &inlineKeyboard,
	}

	mtb.replyMarkup = rm

	return mtb
}

func (mtb *messageTextBuilder) InitReplyKeyboardMarkup(resizeKeyboard bool) *messageTextBuilder {
	keyboardButton := make([][]keyboardButton, 0)
	rm := ReplyKeyboardMarkup{
		Keyboard:       &keyboardButton,
		ResizeKeyboard: &resizeKeyboard,
	}

	mtb.replyMarkup = rm

	return mtb
}

func (mtb *messageTextBuilder) InitReplyKeyboardHide() *messageTextBuilder {
	rm := ReplyKeyboardHide{
		HideKeyboard: true,
	}

	mtb.replyMarkup = rm

	return mtb
}

func (mtb *messageTextBuilder) Text(text string) *messageTextBuilder {
	mtb.text = text
	return mtb
}

func (mtb *messageTextBuilder) ChatId(chatId uint64) *messageTextBuilder {
	mtb.chatId = chatId
	return mtb
}

func (mtb *messageTextBuilder) ParseMode(parseMode string) *messageTextBuilder {
	*mtb.parseMode = parseMode
	return mtb
}

func (mtb *messageTextBuilder) DisableWebPagePreview(disableWebPagePreview bool) *messageTextBuilder {
	*mtb.disableWebPagePreview = disableWebPagePreview
	return mtb
}

func (mtb *messageTextBuilder) DisableNotification(disableNotification bool) *messageTextBuilder {
	*mtb.disableNotification = disableNotification
	return mtb
}

func (mtb *messageTextBuilder) ReplyToMessageId(replyToMessageId uint64) *messageTextBuilder {
	*mtb.replyToMessageId = replyToMessageId
	return mtb
}

func (mtb *messageTextBuilder) AddInlineButton(ik inlineKeyboard) *messageTextBuilder {
	ln := len(*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard) - 1

	if ln == -1 {
		*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard = append(*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard, make([]inlineKeyboard, 0))
		ln = 0
	}

	(*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard)[ln] = append((*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard)[ln], ik)

	return mtb
}

func (mtb *messageTextBuilder) AddInlineButtonWithAllParams(text string, url *string, callbackData *string,
	switchInlineQuery *string, switchInlineQueryCurrentChat *string, cg *callbackGame) *messageTextBuilder {
	inlKey := inlineKeyboard{
		Text:                         text,
		URL:                          url,
		CallbackData:                 callbackData,
		SwitchInlineQuery:            switchInlineQuery,
		SwitchInlineQueryCurrentChat: switchInlineQueryCurrentChat,
		CallbackGame:                 cg,
	}

	return mtb.AddInlineButton(inlKey)
}

func (mtb *messageTextBuilder) AddInlineButtonUrl(text string, url string) *messageTextBuilder {
	inlKey := inlineKeyboard{
		Text: text,
		URL:  &url,
	}

	return mtb.AddInlineButton(inlKey)
}

func (mtb *messageTextBuilder) AddInlineButtonCallBack(text string, callbackData string) *messageTextBuilder {
	inlKey := inlineKeyboard{
		Text:         text,
		CallbackData: &callbackData,
	}

	return mtb.AddInlineButton(inlKey)
}

func (mtb *messageTextBuilder) AddInlineButtonLine() *messageTextBuilder {
	*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard = append(*(mtb.replyMarkup.(ReplyMarkup)).InlineKeyboard, make([]inlineKeyboard, 0))

	return mtb
}

// Add KeyBordButton

func (mtb *messageTextBuilder) AddKeyboardButton(ik keyboardButton) *messageTextBuilder {
	ln := len(*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard) - 1

	if ln == -1 {
		*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard = append(*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard, make([]keyboardButton, 0))
		ln = 0
	}

	(*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard)[ln] = append((*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard)[ln], ik)

	return mtb
}

func (mtb *messageTextBuilder) AddKeyboardButtonText(text string) *messageTextBuilder {
	inlKey := keyboardButton{
		Text: &text,
	}

	return mtb.AddKeyboardButton(inlKey)
}

func (mtb *messageTextBuilder) AddKeyboardButtonLocation(text string) *messageTextBuilder {
	requestLocation := true

	inlKey := keyboardButton{
		Text:            &text,
		RequestLocation: &requestLocation,
	}

	return mtb.AddKeyboardButton(inlKey)
}

func (mtb *messageTextBuilder) AddKeyboardButtonContact(text string) *messageTextBuilder {
	requestContact := true

	inlKey := keyboardButton{
		Text:           &text,
		RequestContact: &requestContact,
	}

	return mtb.AddKeyboardButton(inlKey)
}

func (mtb *messageTextBuilder) AddKeyboardButtonLine() *messageTextBuilder {
	*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard = append(*(mtb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard, make([]keyboardButton, 0))

	return mtb
}

func (mtb *messageTextBuilder) Build() MessageText {
	var mt MessageText

	mt.ChatId = mtb.chatId
	mt.Text = mtb.text

	if mtb.parseMode != nil {
		mt.ParseMode = mtb.parseMode
	}

	if mtb.parseMode != nil {
		mt.ParseMode = mtb.parseMode
	}
	if mtb.disableWebPagePreview != nil {
		mt.DisableWebPagePreview = mtb.disableWebPagePreview

	}
	if mtb.disableNotification != nil {
		mt.DisableNotification = mtb.disableNotification
	}
	if mtb.replyToMessageId != nil {
		mt.ReplyToMessageId = mtb.replyToMessageId
	}
	if mtb.replyMarkup != nil {
		mt.ReplyMarkup = mtb.replyMarkup
	}

	return mt

}
