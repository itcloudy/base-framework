package middles

import (
	"crypto/rsa"
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/logs"
	"go.uber.org/zap"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/base-framework/pkg/consts"
	"github.com/itcloudy/base-framework/tools"
)

// Private key for signing and public key for verification
var (
	//verifyKey, signKey []byte
	jwtPrefix = "Bearer"
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

type JwtClaims struct {
	*jwt.StandardClaims
	Name      string
	RoleIds   []string
	RoleCodes []string
	UserId    int
	IsAdmin   bool
}

// Read the key files before starting http handlers
func InitKeys() {

	signBytes, err := ioutil.ReadFile(conf.Config.JwtPrivatePath)
	if err != nil {
		logs.Logger.Fatal("jwt private key file read failed", zap.String("db name", conf.Config.JwtPrivatePath), zap.Error(err))
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		logs.Logger.Fatal("jwt private key parse failed", zap.String("db name", conf.Config.JwtPrivatePath), zap.Error(err))
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(conf.Config.JwtPublicPath)
	if err != nil {
		logs.Logger.Fatal("jwt public key file read failed", zap.String("db name", conf.Config.JwtPrivatePath), zap.Error(err))
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		logs.Logger.Fatal("jwt public key parse failed", zap.String("db name", conf.Config.JwtPrivatePath), zap.Error(err))
		panic(err)
	}
}

// GenerateJWT generates a new JWT token
func GenerateJWT(name string, roleIds, roleCodes []string, userId int, isAdmin bool) string {

	claims := JwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * conf.Config.TokenExpire).Unix(),
		},
		name,
		roleIds,
		roleCodes,
		userId,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	if err != nil {
		logs.Logger.Error("token generator failed", zap.Error(err))
		return ""
	}

	return tools.StringsJoin(jwtPrefix, " ", ss)
}

//JwtAuthorize parse jwt info
func JwtAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		headToken := c.GetHeader("Authorization")
		if headToken != "" {
			headToken = string(headToken[len(jwtPrefix)+1:])
			var jclaim = &JwtClaims{}
			token, err := jwt.ParseWithClaims(headToken, jclaim, func(*jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})
			if err != nil {
				logs.Logger.Error("token parse failed", zap.Error(err))
			}
			if token.Valid {
				jwtClaims := token.Claims.(*JwtClaims)
				c.Set(consts.LoginUserID, jwtClaims.UserId)
				c.Set(consts.LoginUserName, jwtClaims.Name)
				c.Set(consts.LoginUserRoleIds, jwtClaims.RoleIds)
				c.Set(consts.LoginUserRoleCodes, jwtClaims.RoleCodes)
				c.Set(consts.LoginIsAdmin, jwtClaims.IsAdmin)
				c.Set(consts.TokenValid, true)
			} else {
				c.Set(consts.LoginUserID, 0)
				c.Set(consts.LoginUserName, "")
				c.Set(consts.LoginUserRoleIds, []string{})
				c.Set(consts.LoginUserRoleCodes, []string{})
				c.Set(consts.LoginIsAdmin, false)
				c.Set(consts.TokenValid, false)
			}
		}

	}
}
