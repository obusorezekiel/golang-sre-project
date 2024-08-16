package router

import (
 "github.com/gin-gonic/gin"
 "gin-gorm-crud/controller"
)

func StudentRouter(StudentController *controller.StudentController) *gin.Engine {
 service := gin.Default()

 router := service.Group("/api/v1/students")

 router.GET("", StudentController.FindAll)
 router.GET("/:id", StudentController.FindById)
 router.POST("",StudentController.Create)
 router.PATCH("/:id", StudentController.Update)
 router.DELETE("/:id", StudentController.Delete)

 return service
}