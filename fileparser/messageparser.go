package fileparser

import (
	"bufio"
	"models"
	"os"
	"strings"
)

func FileParser(path string) models.EmailMessage {
	file, error := os.Open(path)
	emailMessage := models.EmailMessage{}
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
				from := entryCleaner(line[5:], "email")
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
						recipients = entryCleaner(recipients, "email")
						emailMessage.SetTo(strings.Split(recipients, " ,"))
					} else {
						recipients = entryCleaner(recipients, "email")
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
						recipients = entryCleaner(recipients, "email")
						emailMessage.SetCC(strings.Split(recipients, " ,"))
					} else {
						recipients = entryCleaner(recipients, "email")
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
						recipients = entryCleaner(recipients, "email")
						emailMessage.SetBCC(strings.Split(recipients, " ,"))
					} else {
						recipients = entryCleaner(recipients, "email")
						emailMessage.SetBCC([]string{recipients})
					}
				}
			case "Subject":
				{
					subject := line[8:]
					scanner.Scan()
					line = scanner.Text()
					for !strings.Contains(line, "Mime") {
						if len(line) != 0 && line[0:1] != " " && line[0:1] != "\t" {
							line = " " + line
						}
						subject = subject + line
						scanner.Scan()
						line = scanner.Text()
					}
					subject = entryCleaner(subject, "subject")

					emailMessage.SetSubject(subject)
				}
			case "X-FileName":
				{
					emailMessage.SetAttachment(strings.TrimSpace(line[11:]))
				}
			/* case "X-To":
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
				} */
			case "":
				{
					body := line + "\n"
					for scanner.Scan() {
						line = scanner.Text()
						line = entryCleaner(line, "body")
						body = body + line + "\n"
					}
					emailMessage.SetBody(body)
				}
			}
		}
	}
	defer file.Close()
	return emailMessage
}

func entryCleaner(entry string, entryType string) string {
	cleanEntry := entry
	switch entryType {
	case "email":
		{
			cleanEntry = strings.Trim(entry, " <>")
			cleanEntry = strings.ReplaceAll(cleanEntry, "\"", "")
			cleanEntry = strings.ReplaceAll(cleanEntry, ":@", "@")
			cleanEntry = strings.ReplaceAll(cleanEntry, "'.'", ".")
			cleanEntry = strings.ReplaceAll(cleanEntry, " .", "")
			cleanEntry = strings.ReplaceAll(cleanEntry, ".@", "@")
			cleanEntry = strings.ReplaceAll(cleanEntry, "..", ".")
			cleanEntry = strings.ReplaceAll(cleanEntry, "\\", "")
			cleanEntry = strings.ReplaceAll(cleanEntry, ")@", "@")
			//return cleanEntry
		}
	case "subject":
		{
			cleanEntry = strings.Trim(entry, " <>")
			cleanEntry = strings.ReplaceAll(cleanEntry, "\"", "''")
			cleanEntry = strings.ReplaceAll(cleanEntry, "\t ", " ")
			cleanEntry = strings.ReplaceAll(cleanEntry, "\\", "/")
			//return cleanEntry
		}
	case "body":
		{
			cleanEntry = strings.ReplaceAll(cleanEntry, "\"", "''")
			cleanEntry = strings.ReplaceAll(cleanEntry, "\\", "/")
		}
	}

	return cleanEntry
	/*
		cleanEntry = strings.Trim(entry, " <>")
		cleanEntry = strings.TrimSpace(cleanEntry)
		cleanEntry = strings.ReplaceAll(cleanEntry, "\"", "'")
		cleanEntry = strings.ReplaceAll(cleanEntry, "\\", "/")
		cleanEntry = strings.ReplaceAll(cleanEntry, "'", "")
		cleanEntry = strings.ReplaceAll(cleanEntry, " <.", ", ")
		cleanEntry = strings.ReplaceAll(cleanEntry, ".@", "@")
		cleanEntry = strings.ReplaceAll(cleanEntry, ">", "")
		cleanEntry = strings.TrimSpace(cleanEntry)

		return cleanEntry
	*/
}
