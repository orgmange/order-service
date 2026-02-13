package user

type CreateParam struct {
	Name  string
	Email string
}

type UpdateParam struct {
	Name  string
	Email string
}

func (req *CreateParam) ToModel() (*Model, error) {
	return NewModel(0, req.Name, req.Email)
}
