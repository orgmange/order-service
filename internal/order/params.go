package order

type CreateParam struct {
	CreatorID uint
}

type UpdateParam struct {
	CreatorID uint
}

func (req *CreateParam) ToModel() *Model {
	return &Model{
		CreatorID: req.CreatorID,
	}
}
