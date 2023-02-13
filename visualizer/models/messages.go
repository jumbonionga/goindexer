package models

type Message struct {
	Id     string `json:"_id"`
	Source source `json:"_source"`
}

type source struct {
	Date       string `json:"date"`
	From       string `jsong:"from"`
	To         string `json:"to"`
	Subject    string `json:"subject"`
	Attachment string `json:"attachment"`
	BCC        string `json:"bcc"`
	CC         string `json:"cc"`
	Body       string `json:"body"`
	MessageID  string `json:"message-id"`
	Path       string `json:"path"`
}
