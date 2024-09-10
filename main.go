package main

import (
	"github.com/lenovoz460/GoMenu/infra/api"
	_ "github.com/lenovoz460/GoMenu/infra/repositories/postgres"
)

func main() {
	api.RunServer()
}
