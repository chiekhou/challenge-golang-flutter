package dao

import (
	"example/hello/internal/models"

	"gorm.io/gorm"
)

type GroupeVoyageDAO struct {
	db *gorm.DB
}

func NewGroupeVoyageDAO(db *gorm.DB) *GroupeVoyageDAO {
	return &GroupeVoyageDAO{db: db}
}

func (dao *GroupeVoyageDAO) CreateGroupeVoyage(groupeVoyage *models.GroupeVoyage) error {
	return dao.db.Create(groupeVoyage).Error
}

func (dao *GroupeVoyageDAO) GetGroupeVoyages() ([]models.GroupeVoyage, error) {
	var groupes []models.GroupeVoyage
	err := dao.db.Find(&groupes).Error
	return groupes, err
}

func (dao *GroupeVoyageDAO) GetGroupeVoyageByID(id uint) (*models.GroupeVoyage, error) {
	var groupe models.GroupeVoyage
	err := dao.db.First(&groupe, id).Error
	return &groupe, err
}

func (dao *GroupeVoyageDAO) UpdateGroupeVoyage(groupeVoyage *models.GroupeVoyage) error {
	return dao.db.Save(groupeVoyage).Error
}

func (dao *GroupeVoyageDAO) DeleteGroupeVoyage(id uint) error {
	return dao.db.Delete(&models.GroupeVoyage{}, id).Error
}
