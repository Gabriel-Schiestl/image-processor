package entities

type ImageModel struct {
	ID int `gorm:"primaryKey"`
	Prediction string `gorm:"type:varchar(255)"`
	Image      []byte `gorm:"type:bytea"`
}