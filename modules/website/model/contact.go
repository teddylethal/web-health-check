package modelwebsite

import (
	"errors"
	"github.com/teddlethal/web-health-check/appCommon"
	"strings"
)

const (
	WebsiteContactEntity = "website contact"
	ContactLimit         = 5
)

var (
	ErrAddressCannotBeEmpty       = appCommon.NewErrorResponse(errors.New("address cannot be empty"), "address cannot be empty", "address cannot be empty", "ErrAddressCannotBeEmpty")
	ErrContactMethodCannotBeEmpty = appCommon.NewErrorResponse(errors.New("contact method cannot be empty"), "contact method cannot be empty", "contact method cannot be empty", "ErrContactMethodCannotBeEmpty")
	ErrContactExisted             = appCommon.NewErrorResponse(errors.New("contact is already existed"), "contact is already existed", "contact is already existed", "ErrContactExisted")
	ErrContactExceedLimit         = appCommon.NewErrorResponse(errors.New("the number of contacts is exceeding the limit"), "the number of contacts is exceeding the limit", "the number of contacts is exceeding the limit", "ErrContactExceedLimit")
)

type WebsiteContact struct {
	Id            int    `json:"id" gorm:"column:id;"`
	WebsiteId     string `json:"website_id" gorm:"column:website_id;"`
	Address       string `json:"address" gorm:"column:address;"`
	ContactMethod string `json:"contact_method" gorm:"column:contact_method;"`
}

func (WebsiteContact) TableName() string {
	return "website_contacts"
}

type WebsiteContactCreation struct {
	WebsiteId     int    `json:"website_id" gorm:"column:website_id;"`
	Address       string `json:"address" gorm:"column:address;"`
	ContactMethod string `json:"contact_method" gorm:"column:contact_method;"`
}

func (data *WebsiteContactCreation) Validate() error {
	data.Address = strings.TrimSpace(data.Address)
	if data.Address == "" {
		return ErrAddressCannotBeEmpty
	}

	data.ContactMethod = strings.TrimSpace(data.ContactMethod)
	if data.ContactMethod == "" {
		return ErrContactMethodCannotBeEmpty
	}

	return nil
}

func (WebsiteContactCreation) TableName() string { return WebsiteContact{}.TableName() }

type WebsiteContactUpdate struct {
	Address       *string `json:"address" gorm:"column:address;"`
	ContactMethod *string `json:"contact_method" gorm:"column:contact_method;"`
}

func (WebsiteContactUpdate) TableName() string { return WebsiteContact{}.TableName() }

func (data *WebsiteContactUpdate) Validate() error {
	address := strings.TrimSpace(*data.Address)
	if address == "" {
		return ErrAddressCannotBeEmpty
	}

	contactMethod := strings.TrimSpace(*data.ContactMethod)
	if contactMethod == "" {
		return ErrContactMethodCannotBeEmpty
	}

	return nil
}

type WebsiteContactDelete struct {
	Id int `json:"id" gorm:"column:id;"`
}

func (WebsiteContactDelete) TableName() string { return WebsiteContact{}.TableName() }
