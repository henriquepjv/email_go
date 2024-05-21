// unit test
// go test (inside of dir)

package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t) // external lib testify
	name := "Campaign 1"
	content := "Body"
	contacts := []string{"email1@bol.com", "email2@bol.com"} // creates a slice of strings

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(1, campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
}
