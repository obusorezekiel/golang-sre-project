package request

type CreateStudentRequest struct {
 Name string `validate:"required,min=1,max=200" json:"name"`
 Description string `validate:"required,min=1,max=200" json:"description"`
}

type UpdateStudentRequest struct {
 Id   int    `validate:"required"`
 Name string `validate:"required,max=200,min=1" json:"name"`
 Description string `validate:"required,max=200,min=1" json:"description"`
}