//Routes/Routes.go
package Routes

import (
	"go_medium/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetData)
		grp1.POST("user", Controllers.CreateData)
		grp1.GET("user/:id", Controllers.GetDataByID)
		grp1.PUT("user/:id", Controllers.UpdateData)
		grp1.DELETE("user/:id", Controllers.DeleteData)
	}
	return r
}
