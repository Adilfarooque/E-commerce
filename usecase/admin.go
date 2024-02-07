package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo/domain"
	"github.com/Adilfarooque/Footgo/helper"
	"github.com/Adilfarooque/Footgo/repository"
	"github.com/Adilfarooque/Footgo/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {
	adminCompareDetails, err := repository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	//compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	var adminDetailsResponse models.AdminDetailsResponse
	//copy all details except password and send it back to the front end
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil
}

func DashBoard() (models.CompleteAdminDashboard, error) {
	userDetails, err := repository.DashBoardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	productDetails, err := repository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := repository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	totalRevenue, err := repository.TotalRevenue()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	amountDetails, err := repository.AmountDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashboardProduct: productDetails,
		DashboardOrder:   orderDetails,
		DashboardRevenue: totalRevenue,
		DashboardAmount:  amountDetails,
	}, nil
}

func ShowAllUsers(page, count int) ([]models.UserDetailsAtAdmin, error) {
	users, err := repository.ShowAllUsersIn(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}
	return users, nil
}

func BlockedUser(id string) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}
	err = repository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func UnBlockedUser(id string) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}
	err = repository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}
/*
func FilteredSalesReport(timePeriod string) (models.SalesReport, error) {
	startTime, endTime := helper.GetTimeFromPeriod(timePeriod)
	saleReport, err := repository.FilteredSalesReport(startTime,endTime)
	if err != nil {
		return models.SalesReport{}, err
	}
	return saleReport, nil
}
*/