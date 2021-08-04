package main

import (
	"fmt"

	"github.com/khorasany/coffee/api/backend/routers"
)

func main() {
	fmt.Println("start ICoffee rest API...")
	routers.Router()
}
