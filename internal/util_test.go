package internal

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestTypeAliasHandling(t *testing.T) {
	src := `package test

type User struct {
	Name string
}

type UserAlias = User

type Container struct {
	AliasedUser UserAlias
	DirectUser  User
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("test", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Get the Container struct
	containerObj := pkg.Scope().Lookup("Container")
	if containerObj == nil {
		t.Fatal("Container not found")
	}

	containerNamed := containerObj.Type().(*types.Named)
	containerStruct := containerNamed.Underlying().(*types.Struct)

	// Get AliasedUser field
	var aliasedUserField *types.Var
	var directUserField *types.Var
	for i := 0; i < containerStruct.NumFields(); i++ {
		f := containerStruct.Field(i)
		if f.Name() == "AliasedUser" {
			aliasedUserField = f
		}
		if f.Name() == "DirectUser" {
			directUserField = f
		}
	}

	if aliasedUserField == nil || directUserField == nil {
		t.Fatal("Fields not found")
	}

	t.Logf("AliasedUser type: %T %v", aliasedUserField.Type(), aliasedUserField.Type())
	t.Logf("DirectUser type: %T %v", directUserField.Type(), directUserField.Type())

	// Test GetStructType on both
	aliasedStruct, aliasedOK := GetStructType(aliasedUserField.Type())
	directStruct, directOK := GetStructType(directUserField.Type())

	t.Logf("GetStructType(AliasedUser): ok=%v, struct=%v", aliasedOK, aliasedStruct)
	t.Logf("GetStructType(DirectUser): ok=%v, struct=%v", directOK, directStruct)

	if !aliasedOK {
		t.Error("GetStructType failed for AliasedUser field")
	}

	if !directOK {
		t.Error("GetStructType failed for DirectUser field")
	}

	// Test GetField with nested path
	nameField1, ok1 := GetField(containerStruct, "AliasedUser.Name", false)
	nameField2, ok2 := GetField(containerStruct, "DirectUser.Name", false)

	t.Logf("GetField(AliasedUser.Name): ok=%v, field=%v", ok1, nameField1)
	t.Logf("GetField(DirectUser.Name): ok=%v, field=%v", ok2, nameField2)

	if !ok1 {
		t.Error("GetField failed for AliasedUser.Name")
	}

	if !ok2 {
		t.Error("GetField failed for DirectUser.Name")
	}
}
