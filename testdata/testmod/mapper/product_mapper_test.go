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

	const hardcodedModelValue string = "HARDCODED_MODEL_GETTER_VALUE"
	const hardcodedDomainValue string = "HARDCODED_DOMAIN_GETTER_VALUE"

	t.Run("model to domain", func(t *testing.T) {
		sourceModel := &model.ProductModel{
			Title: "ignore-me",
		}

		var destDomain domain.Product
		if err = productMapper.ProductModelToProduct(ctx, sourceModel, &destDomain); err != nil {
			t.Fatal(err)
		}

		if destDomain.Title != hardcodedModelValue {
			t.Errorf("Expected %s, got %s", hardcodedModelValue, destDomain.Title)
		}

		if destDomain.Title == sourceModel.Title {
			t.Errorf("Expected domain Title to not equal source model Title")
		}
	})

	t.Run("domain to model", func(t *testing.T) {
		sourceModel := &domain.Product{
			Title: "ignore-me",
		}

		var destModel model.ProductModel
		if err = productMapper.ProductToProductModel(ctx, sourceModel, &destModel); err != nil {
			t.Fatal(err)
		}

		if destModel.Title != hardcodedDomainValue {
			t.Errorf("Expected %s, got %s", hardcodedDomainValue, destModel.Title)
		}
		if destModel.Title == sourceModel.Title {
			t.Errorf("Expected model Title to not equal source domain Title")
		}
	})

}
