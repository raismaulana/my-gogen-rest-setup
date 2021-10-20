package controller

import "github.com/casbin/casbin/v2"

type BaseController struct {
	Enforcer *casbin.Enforcer
}
