package main

import (
	"fangao.com/glocalconfigtest/config"
	"fangao.com/glocalconfigtest/user"
)

func main() {
	config.SetConfig(config.Config{
		UserName: "admin",
		PassWord: "123456",
		Age:      18,
		Name:     "admin",
	})

	user.Login()
}
