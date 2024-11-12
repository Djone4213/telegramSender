package entity

type outputFrom struct {
	ID        uint64 `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type outputChat struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type outputThumb struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     uint64 `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type outputVideo struct {
	Duration     int         `json:"duration"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	FileName     string      `json:"file_name"`
	MimeType     string      `json:"mime_type"`
	Thumbnail    outputThumb `json:"thumbnail"`
	Thumb        outputThumb `json:"thumb"`
	FileId       string      `json:"file_id"`
	FileUniqueId string      `json:"file_unique_id"`
	FileSize     int         `json:"file_size"`
}

type result struct {
	MessageID   int64        `json:"message_id"`
	From        outputFrom   `json:"from"`
	Chat        outputChat   `json:"chat"`
	Date        int64        `json:"date"`
	Text        *string      `json:"text"`
	EditDate    *uint64      `json:"edit_date"`
	Caption     *string      `json:"caption"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup"`
	Video       *outputVideo `json:"video"`
}

type OutputBodyMessage struct {
	OK     bool   `json:"ok"`
	Result result `json:"result"`
}

type OutputBodyDeleteMessage struct {
	OK     bool `json:"ok"`
	Result bool `json:"result"`
}
