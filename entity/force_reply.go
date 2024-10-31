package entity

type ForceReply struct {
	ForceReply bool  `json:"force_reply"`
	Selective  *bool `json:"selective"`
}
