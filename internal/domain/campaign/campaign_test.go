// unit test
// go test (inside of dir)

package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign 1"
	content := "Body"
	contacts := []string{"email1@bol.com", "email2@bol.com"} // creates a slice of strings

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != 1 {
		t.Errorf("expected 1")
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts")
	}
}
