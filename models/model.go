package models

type Employee struct {
	ID        uint    `json:"-" gorm:"column:id"`
	Name      string  `json:"name" gorm:"column:name"`
	Addresses Address `json:"address" gorm:"foreignKey:EmpId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Address struct {
	ID          uint   `json:"-" gorm:"column:id"`
	EmpId       uint   `json:"-"`
	City        string `json:"city" gorm:"column:city"`
	State       string `json:"state" gorm:"column:state"`
	Zip         int    `json:"zip" gorm:"column:zip"`
	PhoneNumber string `json:"phone_No" gorm:"column:phone_No"`
}

func (Employee) TableName() string {
	return "employee"
}

func (Address) TableName() string {
	return "address"
}
