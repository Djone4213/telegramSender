package entity

type ReplyKeyboardHide struct {
	HideKeyboard bool  `json:"hide_keyboard"`
	Selective    *bool `json:"selective,omitempty"`
}
