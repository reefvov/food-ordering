package main

import (
	"food-ordering-api/routers/routers"
)

func main() {
	r := setupRouter()
	r.Run()
}
