package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Adilfarooque/Footgo/db"
	"github.com/Adilfarooque/Footgo/domain"
	"github.com/Adilfarooque/Footgo/helper"
	"github.com/Adilfarooque/Footgo/utils/models"
)

func LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
	var details domain.Admin
	if err := db.DB.Raw("SELECT * FROM users WHERE email = ? AND isadmin = true", adminDetails).Scan(&details).Error; err != nil {
		return domain.Admin{}, err
	}
	return details, nil
}

func DashBoardUserDetails() (models.DashBoardUser, error) {
	var userDetails models.DashBoardUser
	err := db.DB.Raw("SELECT COUNT(*) FROM users WHERE isadmin = 'false'").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM users WHERE blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	return models.DashBoardUser{}, nil
}

func DashBoardProductDetails() (models.DashBoardProduct, error) {
	var productDetails models.DashBoardProduct
	err := db.DB.Raw("SELECT COUNT(*) FROM products").Scan(&productDetails.ToatalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM products WHERE sotck <= 0").Scan(&productDetails.OutofStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	return productDetails, nil
}

func DashBoardOrder() (models.DashboardOrder, error) {
	var orderDetails models.DashboardOrder
	err := db.DB.Raw("SELECT COUNT(*) FROM orders WHERE payment_status = 'paid' AND approval = true").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM orders WHERE shipment_status = 'pending' OR shipment_status = 'processing'").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return models.DashboardOrder{}, err
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM orders WHERE shipment_status = 'cancelled'").Scan(&orderDetails.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM orders").Scan(&orderDetails.ToatlOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = db.DB.Raw("SELECT COALESCE(SUM(quantity),0) FROM carts").Scan(&orderDetails.TotalOrderItem).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}
	return orderDetails, nil
}

func TotalRevenue() (models.DashboardRevenue, error) {
	var revenueDetails models.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := db.DB.Raw("SELECT COALESECE(SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true AND created_at >= ? AND created+at <= ?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = helper.GetTimeFromPeriod("month")
	err = db.DB.Raw("SELECT COALESCE (SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true AND created_at >= ? AND created_at <= ? ", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	startTime, endTime = helper.GetTimeFromPeriod("year")
	err = db.DB.Raw("SELECT COALESCE (SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true AND created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}
	return revenueDetails, nil
}

func AmountDetails() (models.DashboardAmount, error) {
	var amountDetails models.DashboardAmount
	err := db.DB.Raw("SELECT COALESCE (SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	err = db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status ='not paid' AND shipment_status ='processing' OR shipment_status = 'pending' OR shipment_status = 'order placed'").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}
	return amountDetails, nil
}

func ShowAllUsersIn(page, count int) ([]models.UserDetailsAtAdmin, error) {
	var user []models.UserDetailsAtAdmin
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * count
	err := db.DB.Raw("SELECT id,firstname,lastname,email,phone,blocked FROM users WHERE isadmin = 'false' limit = ?", count, offset).Scan(&user).Error
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}
	return user, nil
}

func GetUserByID(id string) (domain.User, error) {
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}

	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM users WHERE id = ?", user_id).Scan(&count).Error; err != nil {
		return domain.User{}, err
	}

	if count < 1 {
		return domain.User{}, errors.New("user for the given id doesn't exists")
	}

	var userDetails domain.User
	if err := db.DB.Raw("SELECT * FROM users WHERE id = ?", user_id).Scan(&userDetails).Error; err != nil {
		return domain.User{}, err
	}
	return userDetails, nil
}

func UpdateBlockUserByID(user domain.User) error {
	err := db.DB.Exec("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}
	return nil
}
