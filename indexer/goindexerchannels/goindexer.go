package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Variable creation
	var messages []EmailMessage
	channelPaths := make(chan EmailMessage)
	messagePaths := make(chan *http.Response)
	counter := 0

	/* if len(os.Args) < 2 {
		panic("Please include the folder to index")
	} else {
		path = string(os.Args[1]) + "/maildir"
	} */
	fmt.Println("Retrieving files")
	path := "../enron_mail_20110402/maildir/allen-p"
	filePaths := DirectoryWalk(path)
	fmt.Println("Number of files:", len(filePaths))

	fmt.Println("Parsing files")
	for _, fpath := range filePaths {
		go FileParser(fpath, channelPaths)
	}

	for i := 0; i < len(filePaths); i++ {
		messages = append(messages, <-channelPaths)
	}
	close(channelPaths)
	fmt.Println("Total messages:", len(messages))

	fmt.Println("Injecting messages")
	for _, emessage := range messages {
		go ZincInject(emessage, messagePaths)
	}

	for i := 0; i < len(messages); i++ {
		resp := <-messagePaths
		if resp.StatusCode == 200 {
			counter++
		}
	}
	close(messagePaths)
	fmt.Print("Total injections:", counter)
}

func DirectoryWalk(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	} else {
		return files
	}
}

func FileParser(path string, channelPath chan EmailMessage) {
	file, error := os.Open(path)
	emailMessage := EmailMessage{}
	emailMessage.SetPath(path)

	if error != nil {
		panic("Message file couldn't be found")
	} else {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			switch strings.Split(line, ":")[0] {
			case "Message-ID":
				emailMessage.SetMessageID(strings.Trim(line[11:], " <>"))
			case "Date":
				emailMessage.SetDate(strings.Trim(line[5:], " <>"))
			case "From":
				from := entryCleaner(line[5:])
				emailMessage.SetFrom(from)
			case "To":
				{
					recipients := line[3:]
					if line[len(line)-2:] == ", " {
						endline := false
						for !endline {
							scanner.Scan()
							line = scanner.Text()
							recipients = recipients + " " + strings.Trim(line, "\t")
							if line[len(line)-2:] != ", " {
								endline = true
							}
						}
						recipients = entryCleaner(recipients)
						emailMessage.SetTo(strings.Split(recipients, " ,"))
					} else {
						recipients = entryCleaner(recipients)
						emailMessage.SetTo([]string{recipients})
					}
				}
			case "CC":
				{
					recipients := strings.TrimSpace(line[3:])
					if line[len(line)-2:] == ", " {
						endline := false
						for !endline {
							scanner.Scan()
							line = scanner.Text()
							recipients = recipients + " " + strings.Trim(line, "\t")
							if line[len(line)-2:] != ", " {
								endline = true
							}
						}
						emailMessage.SetCC(strings.Split(recipients, " ,"))
					} else {
						emailMessage.SetCC([]string{recipients})
					}
				}
			case "BCC":
				{
					recipients := strings.TrimSpace(line[4:])
					if line[len(line)-2:] == ", " {
						endline := false
						for !endline {
							scanner.Scan()
							line = scanner.Text()
							recipients = recipients + " " + strings.Trim(line, "\t")
							if line[len(line)-2:] != ", " {
								endline = true
							}
						}
						emailMessage.SetBCC(strings.Split(recipients, " ,"))
					} else {
						emailMessage.SetBCC([]string{recipients})
					}
				}
			case "Subject":
				{
					subject := line[8:]
					scanner.Scan()
					line = scanner.Text()
					for !strings.Contains(line, "Mime") {
						subject = subject + line
						scanner.Scan()
						line = scanner.Text()
					}
					subject = entryCleaner(subject)

					emailMessage.SetSubject(subject)
				}
			case "X-FileName":
				{
					emailMessage.SetAttachment(strings.TrimSpace(line[11:]))
				}
			case "X-To":
				{
					if strings.Contains(line, "undisclosed-recipients") {
						emailMessage.SetTo([]string{"undisclosed recipients"})
					}
				}
			case "X-cc":
				{
					if strings.Contains(line, "undisclosed-recipients") {
						emailMessage.SetCC([]string{"undisclosed recipients"})
					}
				}
			case "X-bcc":
				{
					if strings.Contains(line, "undisclosed-recipients") {
						emailMessage.SetBCC([]string{"undisclosed recipients"})
					}
				}
			case "":
				{
					body := line + "\n"
					for scanner.Scan() {
						line = scanner.Text()
						line = entryCleaner(line)
						body = body + line + "\n"
					}
					emailMessage.SetBody(body)
				}
			}
		}
	}
	defer file.Close()
	channelPath <- emailMessage
}

func entryCleaner(entry string) string {
	cleanEntry := entry
	cleanEntry = strings.Trim(entry, " <>")
	cleanEntry = strings.ReplaceAll(cleanEntry, "\"", "'")
	cleanEntry = strings.ReplaceAll(cleanEntry, "\\", "/")
	cleanEntry = strings.ReplaceAll(cleanEntry, "'", "")
	cleanEntry = strings.ReplaceAll(cleanEntry, " <.", ", ")
	cleanEntry = strings.ReplaceAll(cleanEntry, ".@", "@")
	cleanEntry = strings.ReplaceAll(cleanEntry, ">", "")
	cleanEntry = strings.TrimSpace(cleanEntry)

	return cleanEntry
}

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

func ZincInject(message EmailMessage, messagePath chan *http.Response) {

	data := `{
		"Path": "` + message.GetPath() + `",
		"Message-ID": "` + message.GetMessageID() + `",
		"Date": "` + message.GetDate() + `",
		"From": "` + message.GetFrom() + `",
		"To": "` + message.GetTo() + `",
		"CC": "` + message.GetCC() + `",
		"BCC": "` + message.GetBCC() + `",
		"Subject": "` + message.GetSubject() + `",
		"Attachment": "` + message.GetAttachment() + `",
		"Body": "` + message.GetBody() + `"
	}`

	req, err := http.NewRequest("POST", "http://localhost:4080/api/channelinject/_doc", strings.NewReader(data))
	if err != nil {
		log.Fatal(message.GetPath())
		log.Fatal(err)
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	messagePath <- resp
	defer resp.Body.Close()
}
