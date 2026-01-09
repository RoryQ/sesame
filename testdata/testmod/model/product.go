package model

type ProductModel struct {
	Title string
}

// GetTitle returns a hardcoded value to prove getter is called
// This will ALWAYS return this hardcoded string, ignoring the Title field value
func (p ProductModel) GetTitle() string {
	return "HARDCODED_MODEL_GETTER_VALUE"
}
