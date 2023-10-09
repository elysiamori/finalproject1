package main

import (
	"github.com/elysiamori/finalproject1/kelompok6/routers"
)

func main() {
	r := routers.StartRouter()
	r.Run(":3000")
}
