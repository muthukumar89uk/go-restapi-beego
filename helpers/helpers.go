package helpers

import (
	"testBeego/models"

	"github.com/astaxie/beego/context"
)

var (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = "password"
	Dbname   = "beegoDB"
)

func CheckResponse(ctx *context.Context, res models.Resp, err error) bool {
	if err != nil {
		ctx.ResponseWriter.WriteHeader(res.StatusCode)
		
		ctx.Output.JSON(struct {
			Code    int    `json:"Code"`
			Message string `json:"Message"`
			Error   string `json:"Error"`
		}{res.StatusCode, res.Message, err.Error()}, true, true)
		return false
	}

	ctx.ResponseWriter.WriteHeader(200)

	return true
}
