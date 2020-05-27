package util

import (
	"self-site/setting"
)

func Setup() {
	jwtSecret = []byte(setting.Config.APP.JwtSecret)
}
