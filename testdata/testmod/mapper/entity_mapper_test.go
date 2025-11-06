package mapper_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"example.com/testmod/domain"
	"example.com/testmod/mapper"
	"example.com/testmod/model"
	"github.com/roryq/sesame"
)

// TestEntityMapper_UUIDConversion tests that UUID (named type with underlying [16]byte)
// is properly converted between string and uuid.UUID types.
// This test verifies the fix in util.go where named types are checked before Kind().
func TestEntityMapper_UUIDConversion(t *testing.T) {
	mappers := mapper.NewMappers()
	mapper.AddUUIDConverter(mappers)

	entityMapper, err := sesame.Get[mapper.EntityMapper](mappers, "EntityMapper")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	expectedID := uuid.MustParse("c606e9a0-10c6-49de-bfc3-62761fcd8f0f")

	// Test Model -> Entity
	source := &model.EntityModel{
		ID:   expectedID.String(),
		Name: "Test Entity",
	}
	var entity domain.Entity
	err = entityMapper.EntityModelToEntity(ctx, source, &entity)
	if err != nil {
		t.Fatalf("EntityModelToEntity failed: %v", err)
	}

	if entity.ID != expectedID {
		t.Errorf("Expected ID %v, got %v", expectedID, entity.ID)
	}
	if entity.Name != "Test Entity" {
		t.Errorf("Expected Name 'Test Entity', got '%s'", entity.Name)
	}

	// Test Entity -> Model
	entitySrc := &domain.Entity{
		ID:   expectedID,
		Name: "Test Entity 2",
	}
	var entityModel model.EntityModel
	err = entityMapper.EntityToEntityModel(ctx, entitySrc, &entityModel)
	if err != nil {
		t.Fatalf("EntityToEntityModel failed: %v", err)
	}

	if entityModel.ID != expectedID.String() {
		t.Errorf("Expected ID %s, got %s", expectedID.String(), entityModel.ID)
	}
	if entityModel.Name != "Test Entity 2" {
		t.Errorf("Expected Name 'Test Entity 2', got '%s'", entityModel.Name)
	}
}

// TestAliasedEntityMapper_UUIDConversion tests that type aliases (types.UUID = uuid.UUID)
// are properly converted. This test verifies the fix in internal/util.go for getQualifiedTypeName.
func TestAliasedEntityMapper_UUIDConversion(t *testing.T) {
	mappers := mapper.NewMappers()
	mapper.AddUUIDConverter(mappers)

	aliasedMapper, err := sesame.Get[mapper.AliasedEntityMapper](mappers, "AliasedEntityMapper")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	expectedID := uuid.MustParse("f1f491e0-82b9-49c8-addc-96edad768d8e")

	// Test Model -> Entity (with type alias)
	source := &model.AliasedEntityModel{
		ID:   expectedID.String(),
		Name: "Test Aliased Entity",
	}
	var entity domain.AliasedEntity
	err = aliasedMapper.AliasedEntityModelToAliasedEntity(ctx, source, &entity)
	if err != nil {
		t.Fatalf("AliasedEntityModelToAliasedEntity failed: %v", err)
	}

	// types.UUID is an alias for uuid.UUID, so they should be equal
	if uuid.UUID(entity.ID) != expectedID {
		t.Errorf("Expected ID %v, got %v", expectedID, entity.ID)
	}
	if entity.Name != "Test Aliased Entity" {
		t.Errorf("Expected Name 'Test Aliased Entity', got '%s'", entity.Name)
	}

	// Test Entity -> Model (with type alias)
	entitySrc := &domain.AliasedEntity{
		ID:   expectedID, // types.UUID can be assigned from uuid.UUID
		Name: "Test Aliased Entity 2",
	}
	var entityModel model.AliasedEntityModel
	err = aliasedMapper.AliasedEntityToAliasedEntityModel(ctx, entitySrc, &entityModel)
	if err != nil {
		t.Fatalf("AliasedEntityToAliasedEntityModel failed: %v", err)
	}

	if entityModel.ID != expectedID.String() {
		t.Errorf("Expected ID %s, got %s", expectedID.String(), entityModel.ID)
	}
	if entityModel.Name != "Test Aliased Entity 2" {
		t.Errorf("Expected Name 'Test Aliased Entity 2', got '%s'", entityModel.Name)
	}
}
