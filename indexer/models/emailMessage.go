package models

import "strings"

type EmailMessage struct {
	path       string
	messageID  string
	date       string
	from       string
	to         []string
	cc         []string
	bcc        []string
	subject    string
	body       string
	attachment string
}

func (message *EmailMessage) GetMessageID() string {
	return message.messageID
}

func (message *EmailMessage) SetMessageID(messageID string) {
	message.messageID = messageID
}

func (message *EmailMessage) GetDate() string {
	return message.date
}

func (message *EmailMessage) SetDate(date string) {
	message.date = date
}

func (message *EmailMessage) GetFrom() string {
	return message.from
}

func (message *EmailMessage) SetFrom(from string) {
	message.from = from
}

func (message *EmailMessage) GetTo() string {
	var toSimplify = strings.Join(message.to, ", ")
	return toSimplify
}

func (message *EmailMessage) SetTo(to []string) {
	message.to = to
}

func (message *EmailMessage) GetSubject() string {
	return message.subject
}

func (message *EmailMessage) SetSubject(subject string) {
	message.subject = subject
}

func (message *EmailMessage) GetBody() string {
	return message.body
}

func (message *EmailMessage) SetBody(body string) {
	message.body = body
}

func (message *EmailMessage) GetAttachment() string {
	return message.attachment
}

func (message *EmailMessage) SetAttachment(attachment string) {
	message.attachment = attachment
}

func (message *EmailMessage) GetCC() string {
	var ccSimplify = strings.Join(message.cc, ", ")
	return ccSimplify
}

func (message *EmailMessage) SetCC(cc []string) {
	message.cc = cc
}

func (message *EmailMessage) GetBCC() string {
	var bccSimplify = strings.Join(message.cc, ", ")
	return bccSimplify
}

func (message *EmailMessage) SetBCC(bcc []string) {
	message.bcc = bcc
}

func (message *EmailMessage) GetPath() string {
	return message.path
}

func (message *EmailMessage) SetPath(path string) {
	message.path = path
}
