package service

import (
 "gin-gorm-crud/data/request"
 "gin-gorm-crud/data/response"
)

type StudentService interface {
 Create(student request.CreateStudentRequest) error
 Update(student request.UpdateStudentRequest) error
 Delete(studentId int) error
 FindById(studentId int) (student response.StudentResponse, err error)
 FindAll() (students []response.StudentResponse, err error)
}