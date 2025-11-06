package model

// EntityModel uses string for UUID fields
type EntityModel struct {
	ID   string
	Name string
}

// AliasedEntityModel uses string for UUID fields
type AliasedEntityModel struct {
	ID   string
	Name string
}
