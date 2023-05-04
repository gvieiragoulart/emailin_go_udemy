package campaign

import (
	internalerrors "emailin/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending = "Pending"
	Started = "Started"
	Cancel  = "Canceled"
	Delete  = "Deleted"
	Done    = "Done"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:10"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].ID = xid.New().String()
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}

func (c *Campaign) Cancel() {
	c.Status = Cancel
}

func (c *Campaign) Delete() {
	c.Status = Delete
}
