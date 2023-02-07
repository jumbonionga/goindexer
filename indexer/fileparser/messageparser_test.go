package fileparser

import (
	"models"
	"testing"
)

func TestFileParser(t *testing.T) {
	messagePath := "/home/fernando/golang/src/goindexer/enron_mail_20110402/maildir/baughman-d/deleted_items/303."
	testMessage := models.EmailMessage{}

	messageResult := FileParser(messagePath)

	if messageResult.GetMessageID() != "9348242.1075843012076.JavaMail.evans@thyme" {
		t.Errorf("Message ID not parsed correctly. Parsed: %s. Expected: %s", messageResult.GetMessageID(), testMessage.GetMessageID())
	}
}

func TestEntryCleaner(t *testing.T) {
	table := []struct {
		teststring string
		testresult string
		testtype   string
	}{
		{"bass'.'k.@enron.com", "bass.k@enron.com", "email"},
		{" Congestion in  the   transmission system for the Eastern Interconnect\"", "Congestion in  the   transmission system for the Eastern Interconnect''", "subject"},
		{"<\"mona.petrochko\":@enron.com>", "mona.petrochko@enron.com", "email"},
		{"Alan's Made Some Big Changes", "Alan's Made Some Big Changes", "subject"},
		{"<munn\".\"mary@enron.com>", "munn.mary@enron.com", "email"},
		{"NYISO - Message to MPs regarding email Subject \"NYPA study-Winter\t Locational ICAP requirements\"", "NYISO - Message to MPs regarding email Subject ''NYPA study-Winter Locational ICAP requirements''", "subject"},
		{"\n\n1)  In Nancy's absence, I'd thought I'd keep the\ninteresting-quotes-sharing-practice alive.  This one's from John Waters in\nthe NYT:\n\"Show me a kid who's not sneaking into R-rated movies and I'll show you\na failure.  All the future CEO's of this country are sneaking into\nR-rated movies.\"\n\n\n2) Cameron, you are a football pool GENIUS!",
			"\n\n1)  In Nancy's absence, I'd thought I'd keep the\ninteresting-quotes-sharing-practice alive.  This one's from John Waters in\nthe NYT:\n''Show me a kid who's not sneaking into R-rated movies and I'll show you\na failure.  All the future CEO's of this country are sneaking into\nR-rated movies.''\n\n\n2) Cameron, you are a football pool GENIUS!", "body"},
		{"<\"matthew.fleming\"@enron.com>", "matthew.fleming@enron.com", "email"},
		{"<\"will.shutt\"@host2.webtwister.com@enron.com>", "will.shutt@host2.webtwister.com@enron.com", "email"},
		{"<\"price_reservations\"@nyiso.com@enron.com>", "price_reservations@nyiso.com@enron.com", "email"},
		{"<\"rob.fax\\)\"@mail.cwo.com@enron.com>", "rob.fax@mail.cwo.com@enron.com", "email"},
	}

	for _, item := range table {
		result := entryCleaner(item.teststring, item.testtype)
		if result != item.testresult {
			t.Errorf("Entry not cleaned correctly.\nResult:%s.\nExpected:%s", result, item.testresult)
		}
	}
}
