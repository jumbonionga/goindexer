package zincinject

import (
	"io"
	"log"
	"models"
	"net/http"
	"os"
	"strings"
)

func ZincInject(message models.EmailMessage, file *os.File) {

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

	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_test/_doc", strings.NewReader(data))
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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	} else {
		if strings.Contains(string(body), "error") {
			_, error := file.WriteString(message.GetPath() + ": " + string(body) + "\n")
			if error != nil {
				panic(error)
			}
		}
	}
}
