package main

import (
	"github.com/elysiamori/finalproject1/kelompok6/config"
	"github.com/elysiamori/finalproject1/kelompok6/routers"
)

func main() {
	_, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	r := routers.StartRouter()
	r.Run(":3000")
}
