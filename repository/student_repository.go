package repository

import (
 "gin-gorm-crud/model"
)

type StudentRepository interface {
 FindAll() ([]model.Student, error)
 FindById(studentId int) (student model.Student, err error)
 Save(student model.Student) error
 Update(student model.Student) error
 Delete(stdentId int) error
}