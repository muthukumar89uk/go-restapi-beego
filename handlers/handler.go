package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testBeego/helpers"
	"testBeego/models"
	"testBeego/repository"

	"github.com/astaxie/beego/context"
)

type HandlerController struct {
	Repo repository.RepoOperation
}

func HandlersInstance(repo repository.RepoOperation) *HandlerController {
	return &HandlerController{
		Repo: repo,
	}
}

func (h HandlerController) CreateEmployee(ctx *context.Context) {
	var employee models.Employee

	ctx.ResponseWriter.Header().Add("content-type", "application/json")

	if err := json.NewDecoder(ctx.Request.Body).Decode(&employee); err != nil {
		ctx.ResponseWriter.WriteHeader(400)

		ctx.Output.JSON(struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"`
			Error  string `json:"Error"`
		}{400, "Invalid input", err.Error()}, true, true)
		return
	}

	response, err := h.Repo.InserEmp(employee)

	if helpers.CheckResponse(ctx, response, err) {
		ctx.Output.JSON(response, true, true)
	}
}

func (h HandlerController) GetAllEmployee(ctx *context.Context) {
	ctx.ResponseWriter.Header().Add("content-type", "application/json")

	response, err := h.Repo.GetAll()
	if err != nil {
		ctx.ResponseWriter.WriteHeader(500)

		ctx.Output.JSON(struct {
			Code    int    `json:"Code"`
			Message string `json:"Message"`
			Error   string `json:"Error"`
		}{500, "Failed to fetch all employee data", err.Error()}, true, true)

		return
	}

	ctx.Output.JSON(response, true, true)
}

func (h HandlerController) GetEmployeeById(ctx *context.Context) {
	ctx.ResponseWriter.Header().Add("content-type", "application/json")

	id, err := strconv.Atoi(ctx.Input.Param(":id"))
	if err != nil {
		ctx.ResponseWriter.WriteHeader(400)

		ctx.Output.JSON(struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"`
			Error  string `json:"Error"`
		}{400, "Failed to parse the Id", err.Error()}, true, true)
		return
	}

	ID := uint(id)

	response, err := h.Repo.GetById(ID)
	if helpers.CheckResponse(ctx, response, err) {
		ctx.Output.JSON(response, true, true)
	}
}

func (h HandlerController) UpdateEmployeeById(ctx *context.Context) {
	fmt.Println("Check")
	ctx.ResponseWriter.Header().Add("content-type", "application/json")

	id, err := strconv.Atoi(ctx.Input.Param(":id"))
	if err != nil {
		ctx.ResponseWriter.WriteHeader(400)
		ctx.Output.JSON(struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"`
			Error  string `json:"Error"`
		}{400, "Failed to parse the Id", err.Error()}, true, true)
		return
	}

	ID := uint(id)

	var employee map[string]interface{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&employee); err != nil {
		ctx.ResponseWriter.WriteHeader(400)
		ctx.Output.JSON(struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"`
			Error  string `json:"Error"`
		}{400, "Invalid input", err.Error()}, true, true)
		return
	}

	response, err := h.Repo.UpdateById(ID, employee)
	if helpers.CheckResponse(ctx, response, err) {
		ctx.Output.JSON(response, true, true)
	}
}

func (h HandlerController) DeleteEmployeeById(ctx *context.Context) {
	fmt.Println("Check")
	ctx.ResponseWriter.Header().Add("content-type", "application/json")

	id, err := strconv.Atoi(ctx.Input.Param(":id"))
	if err != nil {
		ctx.ResponseWriter.WriteHeader(400)
		ctx.Output.JSON(struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"`
			Error  string `json:"Error"`
		}{400, "Failed to parse the Id", err.Error()}, true, true)
		return
	}

	ID := uint(id)

	response, err := h.Repo.DeleteById(ID)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(400)
		ctx.Output.JSON(struct {
			Code   int    `json:"Code"`
			Status string `json:"Status"`
			Error  string `json:"Error"`
		}{400, "Failed to parse the Id", err.Error()}, true, true)
		return
	}
	ctx.Output.JSON(response, true, true)
}
