package utils

import (
	"errors"
	"golang_graphql_postgres/configs"
	"golang_graphql_postgres/internal/models"
	"log"

	"gorm.io/gorm"
)

func findUsersByField(field string, value interface{}) (bool, error) {
	var user []models.User

	result := configs.DB.Where(field+" = ?", value).Limit(1).Find(&user)
	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func SearchByEmail(email string) (bool, error) {
	var user1 []models.User
	result1 := configs.DB.Where("email = ?", email).Limit(1).Find(&user1)
	if result1.Error != nil {
		return false, result1.Error
	}
	log.Println(result1.RowsAffected)
	if result1.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

func SearchByUserId(id string) (bool, error) {
	return findUsersByField("id", id)
}

func SearchByUsername(username string) (bool, error) {
	var user2 []models.User
	result2 := configs.DB.Where("username = ?", username).Limit(1).Find(&user2)

	if result2.Error != nil {
		return false, result2.Error
	}
	if result2.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}

func GetRolName(id int) (*models.Role, error) {
	var role models.Role

	result := configs.DB.Where("id = ?", id).First(&role)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &role, nil
}

// func GetUserInfo(userName string) (*models.User, error) {
// 	var user models.User

// 	result := configs.DB.Where("username = ?", userName).First(&user)

//		if result.Error != nil {
//			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//				return nil, nil
//			}
//			return nil, result.Error
//		}
//		return &user, nil
//	}
func GetUserInfo(userName string) (*models.User, error) {
	var user models.User
	result := configs.DB.Where("username = ?", userName).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("username does not exist")
		}
		return nil, result.Error
	}
	return &user, nil
}
