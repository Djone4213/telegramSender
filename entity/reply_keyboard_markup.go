package entity

type keyboardButton struct {
	Text            *string `json:"text"`
	RequestContact  *bool   `json:"request_contact,omitempty"`
	RequestLocation *bool   `json:"request_location,omitempty"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        *[][]keyboardButton `json:"keyboard"`
	ResizeKeyboard  *bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard *bool               `json:"one_time_keyboard,omitempty"`
	Selective       *bool               `json:"selective,omitempty"`
}
