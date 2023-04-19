package models

// We also specify the tags on each field using backtick annotation.
// This allows us to map each field into a different name when we send them as a
// response since JSON and Go have different naming conventions.

type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
}
