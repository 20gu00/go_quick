package middleware

// logrus记录gin日志,往往是以中間件方式
import (
	"time"

	"github.com/gin-gonic/gin"
	"go_forum/common"
)

// 处理日志的中间件,使用的是logrus(全局,全局路由,组路由,单个路由)
// router中调用
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		statusCode := c.Writer.Status()
		latencyTime := endTime.Sub(startTime) //执行时间
		clientIP := c.ClientIP()
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		common.Logger.Infof("|%d|%15v|%15s|%s|%s", statusCode, latencyTime, clientIP, reqMethod, reqUri)
	}
}
