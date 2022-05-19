package server

import (
	"fmt"
	"os"
)

func Init() {
	r := SetupRouter()
	fmt.Printf(" server port i:%v\n", os.Getenv("SERVER_PORT"))
	err := r.Run(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		return
	}
}
