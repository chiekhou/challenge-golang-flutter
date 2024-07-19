package dao

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"example/hello/internal/models"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                   uint      `gorm:"primary_key; not null" json:"id"`
	FirstName            string    `gorm:"size:64; not null; unique" json:"first_name"`
	LastName             string    `gorm:"size:64; not null" json:"last_name"`
	Photo                string    `gorm:"size:255;" json:"photo"`
	Username             string    `gorm:"size:64; not null; unique " json:"username"`
	Password             string    `gorm:"size:255; not null" json:"password"`
	Email                string    `gorm:"size:100; not null; unique" json:"email"`
	Address              string    `gorm:"size:255; not null" json:"address"`
	RoleID               uint      `gorm:"size: 64; default:1; not null" json:"role_id"`
	ResetPasswordToken   string    `gorm:"size:255;" json:"reset_password_token"`
	ResetPasswordExpires time.Time `gorm:"" json:"reset_password_expires"`
	// GroupVoyage       []GroupeVoyage `gorm:"many2many:group_members;"json:"groupe_voyage"`
}

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) CreateUser(user *models.User) error {
	result := dao.db.Create(user)
	return result.Error
}

func (dao *UserDAO) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := dao.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func generateUniqueToken(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("token length must be greater than zero")
	}

	tokenBytes := make([]byte, length)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(tokenBytes), nil
}

func (dao *UserDAO) GeneratePasswordResetToken(userID uint) (string, error) {
	token, err := generateUniqueToken(32) // Adjust the token length as needed
	if err != nil {
		return "", err
	}

	expiration := time.Now().Add(time.Hour * 24) // Token valid for 24 hours
	err = dao.db.Model(&User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"reset_password_token":   token,
		"reset_password_expires": expiration,
	}).Error
	if err != nil {
		return "", err
	}

	return token, nil
}

func (dao *UserDAO) SetPasswordResetToken(userID uint, token string, expiration time.Time) error {
	err := dao.db.Model(&User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"reset_password_token":   token,
		"reset_password_expires": expiration,
	}).Error
	return err
}

func (dao *UserDAO) VerifyPasswordResetToken(userID uint, token string) (bool, error) {
	var user User
	err := dao.db.First(&user, userID).Error
	if err != nil {
		return false, err
	}

	if user.ResetPasswordToken != token || user.ResetPasswordExpires.Before(time.Now()) {
		return false, nil
	}

	return true, nil
}

func (dao *UserDAO) UpdatePassword(userID uint, newPassword string) error {
	err := dao.db.Model(&User{}).Where("id = ?", userID).Update("password", newPassword).Error
	return err
}
