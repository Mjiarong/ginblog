package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct{
	Username string `json:"username"`
	//Password string `json:"password"`
	jwt.StandardClaims
}

//生成token
func SetToken(username string)(string,int){
	expireTime:=time.Now().Add(10*time.Hour)
	SetClaims := MyClaims{
		Username: username,
		//Password: password,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "ginblog",//签发人
		},
	}

	reqClaim:=jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token,err:=reqClaim.SignedString(JwtKey)
	if err !=nil{
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCESS
}

//验证token
func CheckToken(token string) (*MyClaims,int) {
	setToken,_:=jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})

	if key,_:= setToken.Claims.(*MyClaims);setToken.Valid{
		return  key,errmsg.SUCCESS
	}else {
		return nil,errmsg.ERROR
	}

}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader:=c.Request.Header.Get("Authorization")
		code:=errmsg.SUCCESS
		if tokenHeader==""{
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ",2)
		if len(checkToken)!=2&&checkToken[0]!="Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key,tCode:=CheckToken(checkToken[1])
		if tCode == errmsg.ERROR{
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
		}
		if time.Now().Unix()>key.ExpiresAt{
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
		}

		c.Set("username",key.Username)
		c.Next()
	}
}