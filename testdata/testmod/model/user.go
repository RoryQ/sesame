package model

type UserModel struct {
	ID        string
	Name      string
	UpdatedAt string
	DeletedAt *string
	Address   *AddressModel
}
