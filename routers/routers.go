package routers

import (
	"github.com/gin-gonic/gin"
	v1 "myblog/routers/api/v1"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	apiV1 := r.Group("/v1")
	{
		apiV1.GET("/category", v1.GetCategoryList)
		apiV1.GET("/category/:id", v1.GetCategory)
		apiV1.POST("/category", v1.CreateCategory)
		apiV1.PUT("/category/:id", v1.UpdateCategory)
		apiV1.DELETE("/category/:id", v1.DeleteCategory)

		apiV1.GET("/tag", v1.GetTagList)
		apiV1.POST("/tag", v1.CreateTag)
		apiV1.PUT("/tag/:id", v1.UpdateTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)
	}

	return r
}
