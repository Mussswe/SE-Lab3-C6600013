package entity

type Evaluate struct {
	ProductID int    `json:"product_id" gorm:"not null"`
	Rating    int    `json:"rating" gorm:"not null"`
	Comment   string `json:"comment" gorm:"type:text"`
}
