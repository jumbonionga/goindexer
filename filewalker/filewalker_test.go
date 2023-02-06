package filewalker

import "testing"

func TestDirectoryWalk(t *testing.T) {
	table := []struct {
		path string
		size int
	}{
		{"/home/fernando/golang/src/goindexer/enron_mail_20110402/maildir/allen-p", 3034},
		{"/home/fernando/golang/src/goindexer/enron_mail_20110402/maildir", 517424},
	}

	for _, item := range table {
		results := DirectoryWalk(item.path)
		if len(results) != item.size {
			t.Errorf("Incorrect number of files found in path %s. Found: %d. Expected: %d", item.path, len(results), item.size)
		}
	}
}
