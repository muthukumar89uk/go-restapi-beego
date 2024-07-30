package routers

import (
	"testBeego/handlers"
	"testBeego/repository"

	"github.com/astaxie/beego"
)

func Router(repo repository.RepoOperation) {
	operation := handlers.HandlersInstance(repo)

	beego.Post("/create/employee", operation.CreateEmployee)
	beego.Get("/get/employee", operation.GetAllEmployee)
	beego.Get("/getById/:id", operation.GetEmployeeById)
	beego.Put("/updateById/:id", operation.UpdateEmployeeById)
	beego.Delete("/deleteById/:id", operation.DeleteEmployeeById)

	beego.Run(":9000")
}
