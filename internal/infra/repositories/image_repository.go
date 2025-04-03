package repositories

import (
	"github.com/Gabriel-Schiestl/image-processor/internal/domain/interfaces"
	"github.com/Gabriel-Schiestl/image-processor/internal/domain/models"
	"github.com/Gabriel-Schiestl/image-processor/internal/infra/entities"
	"github.com/Gabriel-Schiestl/image-processor/internal/infra/mappers"
	"gorm.io/gorm"
)

type imageRepository struct {
	DB *gorm.DB
	Mapper mappers.ImageMapper
}

func NewImageRepository(db *gorm.DB) interfaces.ImageRepository {
	db.AutoMigrate(&entities.ImageModel{})
	return &imageRepository{DB: db, Mapper: mappers.ImageMapper{}}
}

func (r *imageRepository) Save(image models.Image) error {
	model := r.Mapper.ToDomainModel(image)
	if err := r.DB.Create(model).Error; err != nil {
		return err
	}
	
	return nil
}