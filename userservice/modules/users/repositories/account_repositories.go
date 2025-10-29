package repositories

import (
	"fmt"
	"userservice/modules/entities"
	"userservice/modules/logs"

	"gorm.io/gorm"
)

type account_userRepository struct {
	db *gorm.DB
	// producer    services_producer.EventProducer
}

func NewAccount_userRepository(db *gorm.DB) entities.Account_userRepository {
	// apiData(db)
	db.Table("account_user").AutoMigrate(&entities.Account_user{})
	return account_userRepository{db}
}

func (a account_userRepository) CreateAccount(acc *entities.Account_user) (*entities.Account_user, error) {
	exec := a.db.Create(&acc)
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function CreatedAccount has error with user_id : %v", acc.UserID))
		return nil, exec.Error
	}
	logs.Debug(fmt.Sprintf("Created Account Successfully with ID : %v", acc.UserID))
	return acc, nil
}

func (a account_userRepository) UpdateAccount(acc *entities.Account_user) (*entities.Account_user, error) {
	exec := a.db.Model(&entities.Account_user{}).Where("user_id=?", acc.UserID).Updates(acc)
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function UpdataAccount has error with user_id : %v", acc.UserID))
		return nil, exec.Error
	}
	logs.Debug(fmt.Sprintf("Updated Account Successfully with ID : %v", acc.UserID))
	return acc, nil

}

func (a account_userRepository) DeleteAccount(userid string) (string, error) {
	// exec := a.db.Delete(&Account{}, userid)
	exec := a.db.Table("account_user").Where("user_id=?", userid).Delete(&entities.Account_user{})
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function DeleteAccount has error with user_id : %v", userid))
		return "", exec.Error
	}
	logs.Debug(fmt.Sprintf("Deleted Account Successfully with ID %v", userid))
	return userid, nil
}

func (a account_userRepository) CheckAccount(userid string) (string, error) {
	account := entities.Account_user{}

	err := a.db.Where("user_id=?", userid).First(&account).Error
	if err != nil {
		logs.Error(fmt.Sprintf("Function CheckAccount has error with user_id : %v", userid))
		return userid, err
	}
	logs.Debug(fmt.Sprintf("Check Account Data : %v", userid))
	return userid, nil
}
