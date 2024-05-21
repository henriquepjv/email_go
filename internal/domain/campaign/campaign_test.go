// unit test
// go test (inside of dir)

package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ( // const only accept primitive kinds, var creates a slice that should not be modified
	Name     = "Campaign 1"
	Content  = "Body"
	Contacts = []string{"email1@bol.com", "email2@bol.com"} // creates a slice of strings
)

func Test_NewCampaign_Creation(t *testing.T) {
	assert := assert.New(t) // external lib testify
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(Name, Content, Contacts)

	assert.NotNil(campaign.ID)
	assert.NotNil(campaign.CreatedAt)
	assert.Greater(campaign.CreatedAt, now)
	assert.Equal(Name, campaign.Name)
	assert.Equal(Content, campaign.Content)
	assert.Equal(len(Contacts), len(campaign.Contacts))
}

func Test_NewCampaign_Validate_Name(t *testing.T) {
	assert := assert.New(t) // external lib testify

	_, err := NewCampaign("", Content, Contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_Validate_Content(t *testing.T) {
	assert := assert.New(t) // external lib testify

	_, err := NewCampaign(Name, "", Contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_Validate_Contacts(t *testing.T) {
	assert := assert.New(t) // external lib testify

	_, err := NewCampaign(Name, Content, []string{})

	assert.Equal("at least one email is required", err.Error())
}
