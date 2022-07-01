package models

type (
	Book struct {
		ISBN   uint64 `gorm:"column:id" json:"isbn"`
		Title  string `gorm:"title" json:"title"`
		Author string `gorm:"column:author" json:"author"`
	}

	CreateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	UpdateBook struct {
		ISBN  uint64
		Title string `json:"title"`
	}
)
