package middlewares

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Request() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		str, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Next()
			return
		}
		var postData = make(map[string]interface{})
		err = json.Unmarshal(str, &postData)
		if err != nil || len(postData) == 0 {
			ctx.Next()
			return
		}
		ctx.Set("userParam", postData)
		ctx.Next()
	}
}
