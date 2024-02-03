package urls

import "gorm.io/gorm"

type Urls struct {
	gorm.Model

	Url  string
	Hash string `gorm:"unqiue"`
}
