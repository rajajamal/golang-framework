package models

type (
	User struct {
		ID       int    `gorm:"column:id" json:"id"`
		Username string `gorm:"column:username" json:"username"`
		Password string `gorm:"column:password" json:"-"`
	}

	CreateUser struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	UpdateUser struct {
		ID       int
		Username string `json:"username" validate:"required"`
	}

	Login struct {
		CreateUser
	}

	UpdatePassword struct {
		ID          int
		OldPassword string `json:"old_password" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}
)
