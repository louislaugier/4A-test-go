package content

import (
	"time"
)

// Content export
type Content interface {
	importFromCSV(filePath string) ([]*Item, error)
	itemsGET(items []*Item, limit int)
}

// Item export
type Item struct {
	ID            int
	Type          string
	Article       *Article
	Video         *Video
	Mooc          *Mooc
	Certification *Certification
	Source        string
	URL           string
}

// User export
type User struct {
	Username  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	LastSeen  time.Time
}

// Err export
type Err struct {
	StatusCode int
	Message    error
}

// Error export
// func (e *Err) Error() string {
// 	txt := e.Message + "(" + strconv.Itoa(e.StatusCode) + ")"
// 	if e.StatusCode != 0 {
// 		txt += "(" + strconv.Itoa(e.StatusCode) + ")"
// 	}
// 	return txt
// }
