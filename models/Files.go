package models

// Type: Document[txt, pdf], Image

type Files struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Type  string `json:"type" gorm:"unique"`
	Owner string `json:"owner" gorm:"not null"`
}
