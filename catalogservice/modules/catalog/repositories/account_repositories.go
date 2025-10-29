package repositories

import (
	"fmt"
	"catalogservice/modules/entities"
	"catalogservice/modules/logs"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
	// producer    services_producer.EventProducer
}

func NewAccountRepository(db *gorm.DB) entities.AccountRepository {
	// apiData(db)
	db.Table("account").AutoMigrate(&entities.Account{})
	return accountRepository{db}
}

func (a accountRepository) CreateAccount(acc *entities.Account) (*entities.Account, error) {
	exec := a.db.Create(&acc)
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function CreatedAccount has error with user_id : %v", acc.UserID))
		return nil, exec.Error
	}
	logs.Debug(fmt.Sprintf("Created Account Successfully with ID : %v", acc.UserID))
	return acc, nil
}

func (a accountRepository) UpdateAccount(acc *entities.Account) (*entities.Account, error) {
	exec := a.db.Model(&entities.Account{}).Where("user_id=?", acc.UserID).Updates(acc)
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function UpdataAccount has error with user_id : %v", acc.UserID))
		return nil, exec.Error
	}
	logs.Debug(fmt.Sprintf("Updated Account Successfully with ID : %v", acc.UserID))
	return acc, nil

}

func (a accountRepository) DeleteAccount(userid string) (string, error) {
	// exec := a.db.Delete(&Account{}, userid)
	exec := a.db.Table("account_user").Where("user_id=?", userid).Delete(&entities.Account{})
	if exec.Error != nil {
		logs.Error(fmt.Sprintf("Function DeleteAccount has error with user_id : %v", userid))
		return "", exec.Error
	}
	logs.Debug(fmt.Sprintf("Deleted Account Successfully with ID %v", userid))
	return userid, nil
}

func (a accountRepository) CheckAccount(userid string) (string, error) {
	account := entities.Account{}

	err := a.db.Where("user_id=?", userid).First(&account).Error
	if err != nil {
		logs.Error(fmt.Sprintf("Function CheckAccount has error with user_id : %v", userid))
		return userid, err
	}
	logs.Debug(fmt.Sprintf("Check Account Data : %v", userid))
	return userid, nil
}
