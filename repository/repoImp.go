package repository

import (
	"fmt"
	"testBeego/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Controller struct {
	Db *gorm.DB
}

func (c *Controller) InserEmp(employee models.Employee) (models.Resp, error) {
	if err := c.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}), // columns to update on conflict
	}).Create(&employee).Error; err != nil {
		return models.Resp{
			StatusCode: 500,
			Message:    "Internal server error",
			Data:       nil,
		}, err
	}

	return models.Resp{
		StatusCode: 200,
		Message:    "Inserted successfully",
		Data:       employee,
	}, nil
}

func (c *Controller) GetAll() ([]models.Employee, error) {
	var emp []models.Employee
	if err := c.Db.Preload("Addresses").Find(&emp).Error; err != nil {
		return nil, err
	}
	return emp, nil
}

func (c *Controller) GetById(id uint) (models.Resp, error) {
	var emp models.Employee

	if err := c.Db.Preload("Addresses").Where("id=?", id).First(&emp).Error; err != nil {
		return models.Resp{
			StatusCode: 500,
			Message:    "Internal server error",
			Data:       nil,
		}, err
	}

	return models.Resp{
		StatusCode: 200,
		Message:    "Data Retrived Successfully",
		Data:       emp,
	}, nil
}

func (c *Controller) UpdateById(id uint, data map[string]interface{}) (models.Resp, error) {
	var emp models.Employee

	if empErr := c.Db.Debug().Preload("Addresses").First(&emp, id).Error; empErr != nil {
		fmt.Println("Error fetching employee:", empErr)
		return models.Resp{
			StatusCode: 500,
			Message:    "Internal server error",
			Data:       nil,
		}, empErr
	}

	if name, ok := data["name"]; ok {
		emp.Name = name.(string)
	}

	if addressData, ok := data["address"]; ok {
		addressMap := addressData.(map[string]interface{})

		if city, ok := addressMap["city"].(string); ok {
			emp.Addresses.City = city
		}
		if state, ok := addressMap["state"].(string); ok {
			emp.Addresses.State = state
		}
		if zip, ok := addressMap["zip"].(int); ok {
			emp.Addresses.Zip = zip
		}
		if phoneNo, ok := addressMap["phone_No"].(string); ok {
			emp.Addresses.PhoneNumber = phoneNo
		}
	}

	if updateErr := c.Db.Debug().Updates(&emp).Error; updateErr != nil {
		fmt.Println("Error updating employee:", updateErr)
		return models.Resp{
			StatusCode: 500,
			Message:    "Internal server error",
			Data:       nil,
		}, updateErr
	}

	fmt.Println("Emp", emp)

	return models.Resp{
		StatusCode: 200,
		Message:    "Data Updated Successfully",
		Data:       emp,
	}, nil
}

func (c *Controller) DeleteById(id uint) (map[string]interface{}, error) {
	var emp models.Employee
	if err := c.Db.Preload("Addresses").Delete(&emp, "id=?", id).Error; err != nil {
		return map[string]interface{}{
			"StatusCode": 500,
			"Message":    "Internal server error",
		}, err
	}

	return map[string]interface{}{
		"StatusCode": 200,
		"Message":    "Data Deleted Successfully",
	}, nil
}
