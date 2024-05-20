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

func NewCampaign(name string, content string, emails []string) *Campaign {
	contacts := make([]Contact, len(emails)) // creates the array with defined range
	for index, email := range emails {       // populate
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        1,
		Name:      name,
		Content:   content,
		CreatedAt: time.Now(),
		Contacts:  contacts,
	}
}
