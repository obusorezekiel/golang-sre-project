package controller

import (
 "github.com/gin-gonic/gin"
 "net/http"
 "gin-gorm-crud/data/request"
 "gin-gorm-crud/data/response"
 "gin-gorm-crud/service"
 "strconv"
)

type StudentController struct {
 StudentService service.StudentService
}

func NewStudentController(service service.StudentService) *StudentController {
 return &StudentController{StudentService: service}
}

func (controller *StudentController) FindAll(ctx *gin.Context) {
 data, err := controller.StudentService.FindAll()

 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   data,
 }
 
 ctx.JSON(http.StatusOK, res)
}

func (controller *StudentController) FindById(ctx *gin.Context) {
 studentId := ctx.Param("id")
 id, err := strconv.Atoi(studentId)

 data, err := controller.StudentService.FindById(id)
 if err != nil {
  ctx.JSON(http.StatusNotFound, response.ErrorResponse{
   Code: 404,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   data,
 }
 ctx.JSON(http.StatusOK, res)
}

func (controller *StudentController) Create(ctx *gin.Context) {
 req := request.CreateStudentRequest{}
 ctx.ShouldBindJSON(&req)
 
 err := controller.StudentService.Create(req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   nil,
 }

 ctx.JSON(http.StatusOK, res)
}

func (controller *StudentController) Update(ctx *gin.Context) {
 req := request.UpdateStudentRequest{}
 err := ctx.ShouldBindJSON(&req)

 studentId := ctx.Param("id")
 id, err := strconv.Atoi(studentId)

 _, err = controller.StudentService.FindById(id)
 if err != nil {
  ctx.JSON(http.StatusNotFound, response.ErrorResponse{
   Code: 404,
   Message: err.Error(),
  })
  return
 }

 req.Id = id

 err = controller.StudentService.Update(req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   nil,
 }

 ctx.JSON(http.StatusOK, res)
}

func (controller *StudentController) Delete(ctx *gin.Context) {
 studentId := ctx.Param("id")
 id, err := strconv.Atoi(studentId)

 _, err = controller.StudentService.FindById(id)
 if err != nil {
  ctx.JSON(http.StatusNotFound, response.ErrorResponse{
   Code: 404,
   Message: err.Error(),
  })
  return
 }
 
 err = controller.StudentService.Delete(id)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   nil,
 }

 ctx.JSON(http.StatusOK, res)
}