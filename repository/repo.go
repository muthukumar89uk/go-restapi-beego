package repository

import (
	"testBeego/models"

	"gorm.io/gorm"
)

type RepoOperation interface {
	InserEmp(models.Employee) (models.Resp, error)
	GetAll() ([]models.Employee, error)
	GetById(id uint) (models.Resp, error)
	UpdateById(uint, map[string]interface{}) (models.Resp, error)
	DeleteById(id uint) (map[string]interface{}, error)
}

func RepoController(db *gorm.DB) RepoOperation {
	return &Controller{Db: db}
}
