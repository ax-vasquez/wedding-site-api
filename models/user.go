package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

// User table
type User struct {
	BaseModel
	// We override Gorm's CreatedAt field so we can set the gorm:"<-:create" directive,
	// which prevents this field from being altered once the record is created
	CreatedAt               time.Time     `gorm:"<-:create"`
	IsAdmin                 bool          `json:"is_admin"`
	IsGoing                 bool          `json:"is_going"`
	CanInviteOthers         bool          `json:"can_invite_others"`
	FirstName               string        `json:"first_name" binding:"required"`
	LastName                string        `json:"last_name" binding:"required"`
	Email                   string        `json:"email" gorm:"uniqueIndex" binding:"required"`
	HorsDoeuvresSelectionId *uuid.UUID    `json:"hors_doeuvres_selection_id"`
	HorsDoeuvresSelection   *HorsDoeuvres `gorm:"foreignKey:HorsDoeuvresSelectionId"`
	EntreeSelectionId       *uuid.UUID    `json:"entree_selection_id"`
	EntreeSelection         *Entree       `gorm:"foreignKey:EntreeSelectionId"`
}

// Maybe create users with given data (if no errors) and returns the number of inserted records
func CreateUsers(users *[]User) (*[]User, error) {
	result := db.Create(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Maybe update a user (if no errors) and returns the number of inserted records
func UpdateUser(u *User) (*[]User, error) {
	result := db.Updates(&u).Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	// USER IN METHOD:  &{{0001-01-01 00:00:00 +0000 UTC 2024-07-27 16:22:00.478746 -0500 CDT {0001-01-01 00:00:00 +0000 UTC false} 22d20919-c685-423e-a2f5-2352ce90d970} 0001-01-01 00:00:00 +0000 UTC false false false  Circlepants  <nil> <nil> <nil> <nil>}
	fmt.Println("USER IN METHOD: ", u)
	res := []User{*u}
	return &res, nil
}

// Maybe delete a user (if no errors) and returns the number of deleted records
func DeleteUser(id uuid.UUID) (*int64, error) {
	// Since our models have DeletedAt set, this makes Gorm "soft delete" records on normal delete operations.
	// We can add .Unscoped() prior to the .Delete() call if we want to permanently-delete them.
	result := db.Delete(&User{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &result.RowsAffected, nil
}

// Find Users by the given ids; returns a User slice
func FindUsers(ids []uuid.UUID) (*[]User, error) {
	var users []User
	result := db.Find(&users, ids)
	if result.Error != nil {
		log.Println("ERROR: ", result.Error.Error())
		return nil, result.Error
	}
	return &users, nil
}
