package domain

type Product struct {
	Title string
}

// GetTitle returns a hardcoded value to prove getter is called
// This will ALWAYS return this hardcoded string, ignoring the Title field value
func (p Product) GetTitle() string {
	return "HARDCODED_DOMAIN_GETTER_VALUE"
}
