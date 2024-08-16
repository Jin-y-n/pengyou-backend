package entity

type Tag struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:63" json:"name"`
	Description string `gorm:"size:255" json:"description"`
}

func (table *Tag) TableName() string {
	return "tag"
}
