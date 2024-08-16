package service

import (
 "github.com/go-playground/validator/v10"
 "gin-gorm-crud/data/request"
 "gin-gorm-crud/data/response"
 "gin-gorm-crud/model"
 "gin-gorm-crud/repository"
 "errors"
)

type StudentServiceImpl struct {
 StudentRepository repository.StudentRepository
 Validate      *validator.Validate
}

func NewStudentServiceImpl(StudentRepository repository.StudentRepository, validate *validator.Validate) (service StudentService, err error) {
 if validate == nil {
  return nil, errors.New("validator instance cannot be nil")
 }
 return &StudentServiceImpl{
  StudentRepository: StudentRepository,
  Validate:       validate,
 }, err
}

func (t StudentServiceImpl) FindAll() (students []response.StudentResponse, err error) {
 result, err := t.StudentRepository.FindAll()
 if err != nil {
  return nil, err
 }

 
 for _, value := range result {
	student := response.StudentResponse{
	 Id:   value.Id,
	 Name: value.Name,
	 Description: value.Description,
	}
	students = append(students, student)
   }
   return students, nil
  }

func (t StudentServiceImpl) FindById(studentId int) (response.StudentResponse, error) {
 data, err := t.StudentRepository.FindById(studentId)
 if err != nil {
  return response.StudentResponse{}, err
 }

 res := response.StudentResponse{
  Id:   data.Id,
  Name: data.Name,
  Description: data.Description,
 }
 return res, nil
}

func (t *StudentServiceImpl) Create(student request.CreateStudentRequest) (err error) {
 err = t.Validate.Struct(student)

 if err != nil {
  return err
 }

 m := model.Student{
  Name: student.Name,
  Description: student.Description,
 }

 t.StudentRepository.Save(m)

 return nil
}

func (t *StudentServiceImpl) Update(student request.UpdateStudentRequest) (err error) {
 data, err := t.StudentRepository.FindById(student.Id)
 
 if err != nil {
  return err
 }

 data.Name = student.Name
 data.Description = student.Description
 t.StudentRepository.Update(data)
 return nil
}

func (t *StudentServiceImpl) Delete(studentId int) (err error) {
 err = t.StudentRepository.Delete(studentId)

 if err != nil {
  return err
 }

 return nil
}