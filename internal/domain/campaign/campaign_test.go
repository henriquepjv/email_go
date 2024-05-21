// unit test
// go test (inside of dir)

package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_Creation(t *testing.T) {
	assert := assert.New(t) // external lib testify
	name := "Campaign 1"
	content := "Body"
	contacts := []string{"email1@bol.com", "email2@bol.com"} // creates a slice of strings
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.NotNil(campaign.CreatedAt)
	assert.Greater(campaign.CreatedAt, now)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
}
