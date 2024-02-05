package repository

import (
	"github.com/Adilfarooque/Footgo/db"
	"github.com/Adilfarooque/Footgo/utils/models"
)

func ShowImage(productID int) ([]models.Image, error) {
	var image []models.Image
	err := db.DB.Raw(`SELECT url FROM images WHERE images.product_id = $1`, productID).Scan(&image).Error
	if err != nil {
		return nil, err
	}
	return image, nil
}
