package router

import (
	"github.com/spf13/viper"
	"github.com/xzeu/tz-layout-advanced/internal/handler"
	"github.com/xzeu/tz-layout-advanced/pkg/jwt"
	"github.com/xzeu/tz-layout-advanced/pkg/log"
)

type RouterDeps struct {
	Logger      *log.Logger
	Config      *viper.Viper
	JWT         *jwt.JWT
	UserHandler *handler.UserHandler
}
