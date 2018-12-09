package middles

import (
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/consts"
	"net/http"
)

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

//CasbinJwtAuthorize returns the authorizer, uses a Casbin enforcer as input
func CasbinJwtAuthorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &BasicAuthorizer{enforcer: e}
		if !a.CheckPermission(c) {
			a.RequirePermission(c)
			//c.Abort is must
			return
		}
	}
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(c *gin.Context) bool {
	if !c.GetBool(consts.TokenValid) {
		return false
	}
	userId := c.GetInt(consts.LoginUserID)
	roles := c.GetStringSlice(consts.LoginUserRoleIds)
	isAdmin := c.GetBool(consts.LoginIsAdmin)
	if isAdmin {
		return true
	}
	roles = append(roles, fmt.Sprintf("user_%d", userId))
	authOk := false
	path := c.Request.URL.Path
	method := c.Request.Method
	for _, key := range roles {
		if a.enforcer.Enforce(key, path, method) {
			authOk = true
			break
		}
	}
	return authOk
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	tokenValid := c.GetBool(consts.TokenValid)
	var res conf.ResponseJson
	if tokenValid {
		c.Writer.WriteHeader(http.StatusForbidden)
		//common.GenResponse(c, consts.PermissionErr, nil, "you have no right for this action1")
		res = conf.ResponseJson{
			Code:    consts.PermissionErr,
			Data:    nil,
			Message: "you have no right for this action",
		}
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		//common.GenResponse(c, consts.TokenValidErr, nil, "token error or expire or invalid1")
		res = conf.ResponseJson{
			Code:    consts.TokenValidErr,
			Data:    nil,
			Message: "token error or expire",
		}
	}

	m, _ := json.Marshal(res)
	c.Writer.Write([]byte(m))
	c.Abort()
}
