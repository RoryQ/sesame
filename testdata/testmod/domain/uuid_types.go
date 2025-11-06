package domain

import (
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

// Entity demonstrates handling of named array types (uuid.UUID is [16]byte)
type Entity struct {
	ID   uuid.UUID // Named type with underlying [16]byte
	Name string
}

// AliasedEntity demonstrates handling of type aliases (types.UUID = uuid.UUID)
type AliasedEntity struct {
	ID   types.UUID // Type alias for uuid.UUID
	Name string
}
