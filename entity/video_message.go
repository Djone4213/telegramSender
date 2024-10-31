package entity

type VideoMessage struct {
	ChatId      uint64      `json:"chat_id"`
	Video       string      `json:"video"`
	Caption     *string     `json:"caption,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type videoMessageBuilder struct {
	chatId      uint64
	video       string
	caption     *string
	replyMarkup interface{}
}

func NewVideoMessageBuilder() *videoMessageBuilder {
	return &videoMessageBuilder{}
}

func (vmb *videoMessageBuilder) ChatId(chatId uint64) *videoMessageBuilder {
	vmb.chatId = chatId
	return vmb
}

func (vmb *videoMessageBuilder) VideoFileId(fileId string) *videoMessageBuilder {
	vmb.video = fileId
	return vmb
}

func (vmb *videoMessageBuilder) Caption(caption string) *videoMessageBuilder {
	vmb.caption = &caption
	return vmb
}

func (vmb *videoMessageBuilder) InitInlineKeyboard() *videoMessageBuilder {
	inlineKeyboard := make([][]inlineKeyboard, 0)
	rm := ReplyMarkup{
		InlineKeyboard: &inlineKeyboard,
	}

	vmb.replyMarkup = rm

	return vmb
}

func (vmb *videoMessageBuilder) InitReplyKeyboardMarkup(resizeKeyboard bool) *videoMessageBuilder {
	keyboardButton := make([][]keyboardButton, 0)
	rm := ReplyKeyboardMarkup{
		Keyboard:       &keyboardButton,
		ResizeKeyboard: &resizeKeyboard,
	}

	vmb.replyMarkup = rm

	return vmb
}

func (vmb *videoMessageBuilder) InitReplyKeyboardHide() *videoMessageBuilder {
	rm := ReplyKeyboardHide{
		HideKeyboard: true,
	}

	vmb.replyMarkup = rm

	return vmb
}

func (vmb *videoMessageBuilder) AddInlineButton(ik inlineKeyboard) *videoMessageBuilder {
	ln := len(*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard) - 1

	if ln == -1 {
		*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard = append(*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard, make([]inlineKeyboard, 0))
		ln = 0
	}

	(*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard)[ln] = append((*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard)[ln], ik)

	return vmb
}

func (vmb *videoMessageBuilder) AddInlineButtonWithAllParams(text string, url *string, callbackData *string,
	switchInlineQuery *string, switchInlineQueryCurrentChat *string, cg *callbackGame) *videoMessageBuilder {
	inlKey := inlineKeyboard{
		Text:                         text,
		URL:                          url,
		CallbackData:                 callbackData,
		SwitchInlineQuery:            switchInlineQuery,
		SwitchInlineQueryCurrentChat: switchInlineQueryCurrentChat,
		CallbackGame:                 cg,
	}

	return vmb.AddInlineButton(inlKey)
}

func (vmb *videoMessageBuilder) AddInlineButtonUrl(text string, url string) *videoMessageBuilder {
	inlKey := inlineKeyboard{
		Text: text,
		URL:  &url,
	}

	return vmb.AddInlineButton(inlKey)
}

func (vmb *videoMessageBuilder) AddInlineButtonCallBack(text string, callbackData string) *videoMessageBuilder {
	inlKey := inlineKeyboard{
		Text:         text,
		CallbackData: &callbackData,
	}

	return vmb.AddInlineButton(inlKey)
}

func (vmb *videoMessageBuilder) AddInlineButtonLine() *videoMessageBuilder {
	*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard = append(*(vmb.replyMarkup.(ReplyMarkup)).InlineKeyboard, make([]inlineKeyboard, 0))

	return vmb
}

// Add KeyBordButton

func (vmb *videoMessageBuilder) AddKeyboardButton(ik keyboardButton) *videoMessageBuilder {
	ln := len(*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard) - 1

	if ln == -1 {
		*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard = append(*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard, make([]keyboardButton, 0))
		ln = 0
	}

	(*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard)[ln] = append((*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard)[ln], ik)

	return vmb
}

func (vmb *videoMessageBuilder) AddKeyboardButtonText(text string) *videoMessageBuilder {
	inlKey := keyboardButton{
		Text: &text,
	}

	return vmb.AddKeyboardButton(inlKey)
}

func (vmb *videoMessageBuilder) AddKeyboardButtonLocation(text string) *videoMessageBuilder {
	requestLocation := true

	inlKey := keyboardButton{
		Text:            &text,
		RequestLocation: &requestLocation,
	}

	return vmb.AddKeyboardButton(inlKey)
}

func (vmb *videoMessageBuilder) AddKeyboardButtonContact(text string) *videoMessageBuilder {
	requestContact := true

	inlKey := keyboardButton{
		Text:           &text,
		RequestContact: &requestContact,
	}

	return vmb.AddKeyboardButton(inlKey)
}

func (vmb *videoMessageBuilder) AddKeyboardButtonLine() *videoMessageBuilder {
	*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard = append(*(vmb.replyMarkup.(ReplyKeyboardMarkup)).Keyboard, make([]keyboardButton, 0))

	return vmb
}

func (vmb *videoMessageBuilder) Build() VideoMessage {
	var vm VideoMessage
	vm.ChatId = vmb.chatId
	vm.Video = vmb.video

	if vmb.caption != nil {
		vm.Caption = vmb.caption
	}

	if vmb.replyMarkup != nil {
		vm.ReplyMarkup = vmb.replyMarkup
	}
	return vm
}
