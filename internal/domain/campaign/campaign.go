package campaign

import "time"

type Contact struct {
	Email string
}

type Campaign struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
}
