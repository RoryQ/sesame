package model

type UserModel struct {
	ID        string
	Name      string
	UpdatedAt string
	DeletedAt *string
	Address   *AddressModel
}

func (um UserModel) GetName() string {
	return um.Name
}
