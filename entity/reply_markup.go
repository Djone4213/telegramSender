package entity

type inlineKeyboard struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *callbackGame `json:"callback_game,omitempty"`
}

type ReplyMarkup struct {
	InlineKeyboard *[][]inlineKeyboard `json:"inline_keyboard,omitempty"`
}
