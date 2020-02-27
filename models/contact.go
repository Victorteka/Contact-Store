package models

import (
	"github.com/jinzhu/gorm"
	u "go-contacts/utils"
	"fmt"
)

//Contact struct
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint `json:"user_id"`
}

//Validate payload
func (contact *Contact) Validate()(map[string] interface{}, bool){
	if contact.Name == ""{
		return u.Message(false, "Name cannot be null"), false
	}
	if len(contact.Phone) < 10{
		return u.Message(false, "Enter a valid phone number"), false
	}
	if contact.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}
		//All the required parameters are present
		return u.Message(true, "success"), true
}

//Create contact
func (contact *Contact) Create() (map[string] interface{}){
	if resp, ok := contact.Validate(); !ok{
		return resp
	}
	GetDB().Create(contact)

	resp:= u.Message(true,"success")
	resp["contact"] = contact
	return resp
}

//GetContacts return user contacts
func GetContacts(user uint) ([]*Contact){
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error

	if err != nil{
		fmt.Println(err)
		return nil
	}
	return contacts
}