package core

type CreateCategoryReq struct {
	Name string  `json:"name"`
	Desc *string `json:"desc"`
}

func (u CreateCategoryReq) Validate() error {
	return validateCategoryFields(u.Name, u.Desc)
}
