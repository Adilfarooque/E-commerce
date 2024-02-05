package usecase

import (
	"github.com/Adilfarooque/Footgo/repository"
	"github.com/Adilfarooque/Footgo/utils/models"
)

func ShowImages(productID int) ([]models.Image, error) {
	image, err := repository.ShowImages(productID)
	if err != nil {
		return nil, err
	}
	return image, nil
}
