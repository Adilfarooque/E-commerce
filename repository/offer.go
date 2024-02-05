package repository

import "github.com/Adilfarooque/Footgo/db"

func FindDiscountPercentageForProduct(id int) (int, error) {
	var percetage int
	err := db.DB.Raw("SELECT discount_percentage FROM product_offers WHERE product_id = $1 ", id).Scan(&percetage).Error
	if err != nil {
		return 0, err
	}
	return percetage, nil
}
