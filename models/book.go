package models

type (
	Book struct {
		ISBN   uint64  `gorm:"column:id" json:"isbn"`
		Title  string  `gorm:"title" json:"title"`
		Author string  `gorm:"column:author" json:"author"`
		Price  float32 `gorm: "column:price" json:"price"`
	}

	CreateBook struct {
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Price  float32 `json: "price"`
	}

	UpdateBook struct {
		ISBN   uint64  `json: "isbn"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Price  float32 `json: "price"`
	}
)
