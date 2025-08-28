package core

type UpdateCategoryReq struct {
	Name string  `json:"name"`
	Desc *string `json:"desc"`
}

func (u UpdateCategoryReq) Validate() error {
	return validateCategoryFields(u.Name, u.Desc)
}
