package repository

import (
 "errors"
 "gin-gorm-crud/model"
 "gin-gorm-crud/data/request"
 "gorm.io/gorm"
)

type StudentRepositoryImpl struct {
 Db *gorm.DB
}

func NewStudentRepositoryImpl(Db *gorm.DB) (StudentRepository) {
 return &StudentRepositoryImpl{Db: Db}
}

func (t StudentRepositoryImpl) FindAll() (students []model.Student, err error) {
 results := t.Db.Find(&students)
 if results.Error != nil {
  return nil, results.Error
 }
 
 return students, nil
}

func (t StudentRepositoryImpl) FindById(studentId int) (student model.Student, err error) {
 result := t.Db.Find(&student, studentId)
 
 if result.Error != nil {
  return model.Student{}, result.Error
 }

 if result.RowsAffected == 0 {
  return model.Student{}, errors.New("Task is not found")
 }

 return student, nil
}

func (t *StudentRepositoryImpl) Save(student model.Student) error {
 result := t.Db.Create(&student)
 if result.Error != nil {
  return result.Error
 }

 return nil
}

func (t *StudentRepositoryImpl) Update(student model.Student) error {
 var data = request.UpdateStudentRequest{
  Id:   student.Id,
  Name: student.Name,
  Description: student.Description,
 }

 result := t.Db.Model(&student).Updates(data)
 if result.Error != nil {
  return result.Error
 }

 return nil
}

func (t *StudentRepositoryImpl) Delete(studentId int) error {
 var student model.Student

 result := t.Db.Where("id = ?", studentId).Delete(&student)
 if result.Error != nil {
  return result.Error
 }

 return nil
}