package main

import (
	"testBeego/repository"
	"testBeego/routers"
)

func main() {
	db := repository.DbConnection()
	repository.Migration(db)

	operation := repository.RepoController(db)
	routers.Router(operation)
}
