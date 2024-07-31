package routers

import (
	"testBeego/handlers"
	"testBeego/repository"

	"github.com/astaxie/beego"
)

func Router(repo repository.RepoOperation) {
	operation := handlers.HandlersInstance(repo)

	beego.Post("/v1/api/create/employee", operation.CreateEmployee)
	beego.Get("/v1/api/get/employees", operation.GetAllEmployee)
	beego.Get("/v1/api/getById/:id", operation.GetEmployeeById)
	beego.Put("/v1/api/updateById/:id", operation.UpdateEmployeeById)
	beego.Delete("/v1/api/deleteById/:id", operation.DeleteEmployeeById)

	beego.Run(":9000")
}
