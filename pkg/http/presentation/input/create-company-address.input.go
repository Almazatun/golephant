package input

type CreateCompanyAddressInput struct {
	Title         string `json:"address" validate:"required,address"`
	IsBaseAddress bool   `json:"is_base_adress" validate:"required"`
}
