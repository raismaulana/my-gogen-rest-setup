package controller

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	Router   gin.IRouter
	Enforcer *casbin.Enforcer
}

func (r *BaseController) Authorize() {

}
