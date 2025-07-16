package model

import "database/sql"

type Execs struct {
	Id                  int              `json:"id"`
	FirstName           string           `json:"firstName"`
	LastName            string           `json:"lastName"`
	Email               string           `json:"email"`
	Role                string           `json:"role"`
	UserName            string           `json:"username"`
	Password            string           `json:"password"`
	PasswordChangeAt    sql.NullString   `json:"passwordChangeAt"`
	UserCreatedAt       sql.NullString   `json:"userCreatedAt"`
	PasswordResetToken   sql.NullString   `json:"passwordResetToken"`
	PasswordCodeExpires sql.NullString   `json:"passwordCodeExpires"`
	InactiveStatus      bool             `json:"inactiveStatus"`
}