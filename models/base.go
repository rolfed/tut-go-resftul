package models

import (
	"github.com/jinzhu/gorm"
)

// Contact data model
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"` // The user that this contact belongs to
}

/*
	This struct function validate the required parameters sent through
	the heep request body returns message and true if the requirement is met
*/
func (contact *Contact) Validate() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return util.Message(false, "Contact name is undefined"), false
	}

	if contact.Phone == "" {
		return util.Message(false, "Phone name is undefined"), false
	}

	if contact.UserID <= 0 {
		return util.Message(false, "User is not recognized"), false
	}

	// All required parameters are present
	return utils.Message(true, "success"), true

}
