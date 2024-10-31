package entity

type outputFrom struct {
	ID        uint64 `json:"ID"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type outputChat struct {
	ID        uint64 `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type result struct {
	MessageID   int64       `json:"message_id"`
	From        outputFrom  `json:"from"`
	Chat        outputChat  `json:"chat"`
	Date        int64       `json:"date"`
	Text        string      `json:"text"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type OutputBody struct {
	OK       bool   `json:"ok"`
	Result   result `json:"result"`
	Date     uint64 `json:"date"`
	Text     string `json:"text"`
	EditDate uint64 `json:"edit_date"`
}
