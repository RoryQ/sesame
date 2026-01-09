package mapper

import (
	"context"
	"testing"

	"example.com/testmod/domain"
	"example.com/testmod/model"
	"github.com/roryq/sesame"
)

func TestProductMapperUsesGetterNotField(t *testing.T) {
	mappers := NewMappers()
	ctx := context.TODO()

	productMapper, err := sesame.Get[ProductMapper](mappers, "ProductMapper")
	if err != nil {
		t.Fatal(err)
	}

	// Test ProductModel -> Product
	// Set the Title field to a different value than what GetTitle() returns
	sourceModel := &model.ProductModel{
		Title: "field-value-that-should-be-ignored",
	}

	// Verify the getter returns the hardcoded value, not the field
	if sourceModel.GetTitle() != "HARDCODED_MODEL_GETTER_VALUE" {
		t.Fatalf("Test setup error: GetTitle() should return hardcoded value, got %s", sourceModel.GetTitle())
	}

	var destDomain domain.Product
	err = productMapper.ProductModelToProduct(ctx, sourceModel, &destDomain)
	if err != nil {
		t.Fatal(err)
	}

	// The mapper should have used GetTitle(), which returns the hardcoded value
	// NOT the field value
	if destDomain.Title != "HARDCODED_MODEL_GETTER_VALUE" {
		t.Errorf("Expected Title to be 'HARDCODED_MODEL_GETTER_VALUE' (from GetTitle()), got '%s'", destDomain.Title)
	}

	if destDomain.Title == "field-value-that-should-be-ignored" {
		t.Error("Mapper used direct field access instead of getter method!")
	}

	// Test Product -> ProductModel (reverse direction)
	sourceDomain := &domain.Product{
		Title: "another-field-value-that-should-be-ignored",
	}

	// Verify the getter returns the hardcoded value
	if sourceDomain.GetTitle() != "HARDCODED_DOMAIN_GETTER_VALUE" {
		t.Fatalf("Test setup error: GetTitle() should return hardcoded value, got %s", sourceDomain.GetTitle())
	}

	var destModel model.ProductModel
	err = productMapper.ProductToProductModel(ctx, sourceDomain, &destModel)
	if err != nil {
		t.Fatal(err)
	}

	// The mapper should have used GetTitle(), which returns the hardcoded value
	if destModel.Title != "HARDCODED_DOMAIN_GETTER_VALUE" {
		t.Errorf("Expected Title to be 'HARDCODED_DOMAIN_GETTER_VALUE' (from GetTitle()), got '%s'", destModel.Title)
	}

	if destModel.Title == "another-field-value-that-should-be-ignored" {
		t.Error("Mapper used direct field access instead of getter method!")
	}
}
