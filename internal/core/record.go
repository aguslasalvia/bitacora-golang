package core

type Record struct {
	ID            int    `sql:"id"`
	Name          string `sql:"name"`
	Lab           string `sql:"lab"`
	Equipment     string `sql:"equipment"`
	StartDateTime int64  `sql:"startDateTime"`
	EndDateTime   int64  `sql:"endDateTime"`
	Received      bool   `sql:"received"`
	Returned      bool   `sql:"returned"`
	Comments      string `sql:"comments"`
	Timestamp     int64  `sql:"timestamp"`
}

// I use int64 because i want to store the dates+hours in Unix format since 1970 in ms
// After a Select for excample I convert that number to readeable DateTime format
