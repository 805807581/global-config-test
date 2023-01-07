package user

import (
	"fmt"

	"fangao.com/glocalconfigtest/config"
)

func Login() {
	cfg := config.GetConfig()
	fmt.Println("登录成功: ", cfg.UserName, cfg.PassWord, cfg.Age, cfg.Name)

}
