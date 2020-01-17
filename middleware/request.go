package middleware

import (
	"fmt"
	"go-gin-web/pkg/common"
	"go-gin-web/pkg/config"
	"go-gin-web/pkg/errMsg"
	"go-gin-web/pkg/rateLimter"
	"go-gin-web/pkg/util"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

// TraceMiddleware 跟踪ID中间件
func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 优先从请求头中获取请求ID，如果没有则使用UUID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = util.GenUuid("trace")
		}
		c.Set("X-Request-Id", traceID)
		c.Next()
	}
}

// CORSMiddleware 跨域请求中间件
// func Cors() gin.HandlerFunc {
// 	var (
// 		corsCfg = config.CorsCfg
// 	)

// 	if !corsCfg.Enable {
// 		return func(c *gin.Context) {
// 			c.Next()
// 		}
// 	}

// 	// return cors.Default()

// 	return cors.New(cors.Config{
// 		AllowOrigins:     corsCfg.AllowOrigins,
// 		AllowMethods:     corsCfg.AllowMethods,
// 		AllowHeaders:     corsCfg.AllowHeaders,
// 		AllowCredentials: corsCfg.AllowCredentials,
// 		MaxAge:           time.Second * time.Duration(corsCfg.MaxAge),
// 	})
// }

// 跨域设置
func CorsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,xauthtoken")
			c.Header("Access-Control-Expose-Headers", "X-Trace-Id")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}

// 限流
func RateLimiter(rateLimiter *rateLimter.RateLimiter) gin.HandlerFunc {

	return func(c *gin.Context) {
		var (
			res     = common.Res{C: c}
			rateCfg = config.RateCfg
		)

		if rateCfg.Enable {
			if !rateLimiter.IsLimit() {
				res.SendJSON(http.StatusTooManyRequests, errMsg.MANYREQUESTS, rateLimiter.ReqCount)
			}
		}

		rateLimiter.Incr()

		c.Next()
	}
}

// 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			res = common.Res{C: c}
		)
		defer func() {
			if err := recover(); err != nil {
				res.SendJSON(http.StatusInternalServerError, errMsg.ERROR, err)
			}
		}()
		c.Next()
	}
}

// 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			res = common.Res{C: c}
		)
		res.SendJSON(http.StatusNotFound, errMsg.PageNotFound, nil)
	}
}

// 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			res = common.Res{C: c}
		)
		res.SendJSON(http.StatusNotFound, errMsg.PageNotFound, nil)
	}
}

// func APIStats() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		t := time.Now()
// 		c.Next()

// 		// if config.StatsDConfig.URL == "" {
// 		// 	return
// 		// }

// 		duration := time.Since(t)
// 		durationMS := int64(duration / 1e6) // 转成毫秒

// 		reqPath := getReqPath(c)
// 		if err := (*model.StatsdClient).Timing(reqPath, durationMS, 1); err != nil {
// 			fmt.Println(err)
// 		}

// 		status := c.Writer.Status()
// 		if status != http.StatusGatewayTimeout && durationMS > 5000 {
// 			timeoutReqPath := strings.Join([]string{"timeout", reqPath}, ":")
// 			if err := (*model.StatsdClient).Inc(timeoutReqPath, 1, 1); err != nil {
// 				fmt.Println(err)
// 			}
// 		}
// 	}
// }
