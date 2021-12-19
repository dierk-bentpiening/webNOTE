package baseapifunc
import "github.com/gin-gonic/gin"
import "time"

import (
	"net/http"
)


func GetTime(c *gin.Context) {
	now := time.Now().Format(time.RFC850)
	c.String(http.StatusOK, now)
}
