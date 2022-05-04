package server

import "github.com/andree37/rlld/config"

func Init() {
	config := config.Getconfig()
	r := SetupRouter()
	r.Run(config.GetString("server.port"))
}
