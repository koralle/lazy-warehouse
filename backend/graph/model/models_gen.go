// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type ModelBase interface {
	IsModelBase()
	GetID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Group struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     *string   `json:"description"`
	ProfileImageURL *string   `json:"profile_image_url"`
	CreatedAt       time.Time `json:"_created_at"`
	UpdatedAt       time.Time `json:"_updated_at"`
}

func (Group) IsModelBase()                 {}
func (this Group) GetID() string           { return this.ID }
func (this Group) GetCreatedAt() time.Time { return this.CreatedAt }
func (this Group) GetUpdatedAt() time.Time { return this.UpdatedAt }

type Role struct {
	ID        string       `json:"id"`
	Name      RoleCategory `json:"name"`
	CreatedAt time.Time    `json:"_created_at"`
	UpdatedAt time.Time    `json:"_updated_at"`
}

func (Role) IsModelBase()                 {}
func (this Role) GetID() string           { return this.ID }
func (this Role) GetCreatedAt() time.Time { return this.CreatedAt }
func (this Role) GetUpdatedAt() time.Time { return this.UpdatedAt }

type User struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	ProfileImageURL *string   `json:"profile_image_url"`
	CreatedAt       time.Time `json:"_created_at"`
	UpdatedAt       time.Time `json:"_updated_at"`
}

func (User) IsModelBase()                 {}
func (this User) GetID() string           { return this.ID }
func (this User) GetCreatedAt() time.Time { return this.CreatedAt }
func (this User) GetUpdatedAt() time.Time { return this.UpdatedAt }

type SetRoleInput struct {
	UserID  string       `json:"user_id"`
	GroupID string       `json:"group_id"`
	Role    RoleCategory `json:"role"`
}

type SetRoleResult struct {
	User  *User        `json:"user"`
	Group *Group       `json:"group"`
	Role  RoleCategory `json:"role"`
}

type RoleCategory string

const (
	RoleCategoryOwner  RoleCategory = "OWNER"
	RoleCategoryEditor RoleCategory = "EDITOR"
	RoleCategoryReader RoleCategory = "READER"
	RoleCategoryNone   RoleCategory = "NONE"
)

var AllRoleCategory = []RoleCategory{
	RoleCategoryOwner,
	RoleCategoryEditor,
	RoleCategoryReader,
	RoleCategoryNone,
}

func (e RoleCategory) IsValid() bool {
	switch e {
	case RoleCategoryOwner, RoleCategoryEditor, RoleCategoryReader, RoleCategoryNone:
		return true
	}
	return false
}

func (e RoleCategory) String() string {
	return string(e)
}

func (e *RoleCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RoleCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RoleCategory", str)
	}
	return nil
}

func (e RoleCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
