package loader

import (
	"fmt"

	"github.com/joho/godotenv"
)

func EnvLoader() {
	err := godotenv.Load("./.env")
	if err == nil {
		fmt.Println("Env loaded successfully ...")
	} else {
		panic("Env loaded Failed ...")
	}
}
