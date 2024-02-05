package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo/db"
	"github.com/Adilfarooque/Footgo/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productBrief []models.ProductBrief
	err := db.DB.Raw(`SELECT * FROM products limit ? offset ?`, count, offset).Scan(&productBrief).Error
	if err != nil {
		return nil, err
	}
	return productBrief, nil
}

func ShowAllProductsFromAdmin(page int, count int) ([]models.ProductBrief, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productBrief []models.ProductBrief
	err := db.DB.Raw(`SELECT * FROM products limit ? offset ?`, count, offset).Scan(&productBrief).Error
	if err != nil {
		return nil, err
	}
	return productBrief, nil
}

func CheckValidateCategory(data map[string]int) error {
	for _, id := range data {
		var count int
		err := db.DB.Raw("SELECT COUNT(*) FROM categories WHERE id =?", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count < 1 {
			return errors.New("doesn't exist")
		}
	}
	return nil
}

func GetProductFromCategory(id int) ([]models.ProductBrief, error) {
	var product []models.ProductBrief
	err := db.DB.Raw(`SELECT * FROM products JOIN categories ON products.category_id = categories.id WHERE categories.id = ?`, id).Scan(&product).Error
	if err != nil {
		return []models.ProductBrief{}, err
	}
	return product, nil
}

func GetQuantityFromProductID(id int) (int, error) {
	var quantity int
	err := db.DB.Raw("SELECT stock FROM products WHERE id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}
	return quantity, nil
}

func GetImage(productID int) ([]string, error) {
	var url []string
	if err := db.DB.Raw(`SELECT * FROM products WHERE product_id = $1`, productID).Scan(&url).Error; err != nil {
		return url, err
	}
	return url, nil
}
