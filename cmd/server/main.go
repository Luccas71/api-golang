package main

import (
	"fmt"

	"github.com/Luccas1/api-golang/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBHost)
}
