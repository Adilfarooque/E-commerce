package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo/repository"
	"github.com/Adilfarooque/Footgo/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {
	productDetails, err := repository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	for i := range productDetails {
		p := &productDetails[i]
		if p.Stock <= 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	//
	for j := range productDetails {
		discount_percentage, err := repository.FindDiscountPercentageForProduct(int(productDetails[j].ID))
		if err != nil {
			return []models.ProductBrief{}, errors.New("there was some error in finding the discount prices")
		}
		var discount float64
		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}
		productDetails[j].DiscountedPrice = productDetails[j].Price - discount

		discount_percentageCategory, err := repository.FindDiscountPercentageForProduct(int(productDetails[j].CategoryID))

		if err != nil {
			return []models.ProductBrief{}, errors.New("there was some error in finding the dicount prices")
		}

		var categorydiscount float64
		if discount_percentageCategory > 0 {
			categorydiscount = (productDetails[j].Price * float64(discount_percentageCategory)) / 100
		}
		productDetails[j].DiscountedPrice = productDetails[j].DiscountedPrice - categorydiscount
	}
	var updateProductDetails []models.ProductBrief
	for _, p := range productDetails {
		img, err := repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		p.Image = img
		updateProductDetails = append(updateProductDetails, p)
	}
	return updateProductDetails, nil
}

func FilerCategory(data map[string]int) ([]models.ProductBrief, error) {
	err := repository.CheckValidateCategory(data)
	if err != nil {
		return []models.ProductBrief{}, err
	}
	var ProductFromCategory []models.ProductBrief
	for _, id := range data {
		product, err := repository.GetProductFromCategory(id)
		if err != nil {
			return []models.ProductBrief{}, err
		}
		for _, products := range product {
			stock, err := repository.GetQuantityFromProductID(int(products.ID))
			if err != nil {
				return []models.ProductBrief{}, err
			}
			if stock <= 0 {
				products.ProductStatus = "out of stock"
			} else {
				products.ProductStatus = "in stock"
			}
			if products.ID != 0 {
				ProductFromCategory = append(ProductFromCategory, products)
			}
		}
	}
	for j := range ProductFromCategory {
		discount_percentage, err := repository.FindDiscountPercentageForProduct(ProductFromCategory[j].ID)
		if err != nil {
			return []models.ProductBrief{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64
		if discount_percentage > 0 {
			discount = (ProductFromCategory[j].Price * float64(discount_percentage)) / 100
		}
		ProductFromCategory[j].DiscountedPrice = ProductFromCategory[j].Price - discount

		discount_percentageCategory, err := repository.FindDiscountPercentageForProduct(int(ProductFromCategory[j].CategoryID))
		if err != nil {
			return []models.ProductBrief{}, errors.New("there was some error in finding the dicounted prices")
		}
		var categorydiscount float64
		if discount_percentageCategory > 0 {
			categorydiscount = (ProductFromCategory[j].Price * float64(discount_percentageCategory)) / 100
		}
		ProductFromCategory[j].DiscountedPrice = ProductFromCategory[j].DiscountedPrice - categorydiscount
	}
	updateProductDetails := make([]models.ProductBrief, 0)
	for _, p := range ProductFromCategory {
		img, err := repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		p.Image = img
		updateProductDetails = append(updateProductDetails, p)
	}
	return updateProductDetails, err
}

