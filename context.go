package sesame

import (
	"fmt"
	"go/types"
	"path"
	"strings"
)

// MappingContext is an interface that contains contextual data for
// the generation.
type MappingContext struct {
	absPkgPath       string
	imports          map[string]string
	varCount         int
	mapperFuncFields []*MapperFuncField
	mapperFuncCount  int
	aliases          map[string]int
}

// MapperFuncField is a mapper function field.
type MapperFuncField struct {
	// FieldName is a name of the field.
	FieldName string

	// MapperFuncName is a name of the mapper function.
	MapperFuncName string

	// Source is a source type of the function.
	Source types.Type

	// Dest is a source type of the function.
	Dest types.Type
}

// Signature returns a function signature.
func (m *MapperFuncField) Signature(mctx *MappingContext) string {
	return fmt.Sprintf("func(%s) (%s, error)",
		GetPreferableTypeSource(m.Source, mctx),
		GetPreferableTypeSource(m.Dest, mctx))
}

// NewMappingContext returns new [MappingContext] .
func NewMappingContext(absPkgPath string) *MappingContext {
	return &MappingContext{
		absPkgPath:       absPkgPath,
		imports:          map[string]string{},
		aliases:          map[string]int{},
		mapperFuncFields: []*MapperFuncField{},
		mapperFuncCount:  0,
	}
}

// AbsolutePackagePath returns na absolute package path of a file will be
// generated this mapping.
func (c *MappingContext) AbsolutePackagePath() string {
	return c.absPkgPath
}

// AddImport adds import path and generate new alias name for it.
func (c *MappingContext) AddImport(importpath string) {
	if importpath == c.AbsolutePackagePath() {
		return
	}
	if _, ok := c.imports[importpath]; !ok {
		_, last := path.Split(importpath)
		alias := strings.ReplaceAll(last, "-", "_")
		if i := c.aliases[alias]; i > 0 {
			alias = fmt.Sprintf("%s%d", alias, i+1)
		}
		c.imports[importpath] = alias
		c.aliases[alias]++
	}
}

// GetImportAlias returns an alias for the given import path.
func (c *MappingContext) GetImportAlias(path string) string {
	c.AddImport(path)
	v, ok := c.imports[path]
	if !ok {
		return ""
	}
	return v
}

// GetImportPath returns a fully qualified path for the given import alias.
// If alias is not found, GetImportPath returns given alias.
func (c *MappingContext) GetImportPath(alias string) string {
	if alias == "" {
		return c.AbsolutePackagePath()
	}
	for key, value := range c.imports {
		if value == alias {
			return key
		}
	}
	return alias
}

// Imports returns a map of the all imports.
// Result map key is an import path and value is an alias.
func (c *MappingContext) Imports() map[string]string {
	return c.imports
}

// NextVarCount returns a var count and increments it.
func (c *MappingContext) NextVarCount() int {
	v := c.varCount
	c.varCount++
	return v
}

// AddMapperFuncField adds a mapper function and generates a field name for it.
func (c *MappingContext) AddMapperFuncField(sourceType types.Type, destType types.Type) {
	sname := GetQualifiedTypeName(sourceType)
	dname := GetQualifiedTypeName(destType)
	if sname == dname {
		return
	}
	mapperFuncName := mappersName(sourceType, destType)
	for _, m := range c.mapperFuncFields {
		if m.MapperFuncName == mapperFuncName {
			return
		}
	}
	fieldName := fmt.Sprintf("mapper%05d", c.mapperFuncCount)
	c.mapperFuncCount++
	c.mapperFuncFields = append(c.mapperFuncFields, &MapperFuncField{
		FieldName:      fieldName,
		MapperFuncName: mapperFuncName,
		Source:         sourceType,
		Dest:           destType,
	})
}

// GetMapperFuncFieldName returns a mapper function field name.
func (c *MappingContext) GetMapperFuncFieldName(sourceType types.Type, destType types.Type) *MapperFuncField {
	c.AddMapperFuncField(sourceType, destType)
	mapperFuncName := mappersName(sourceType, destType)
	for _, m := range c.mapperFuncFields {
		if m.MapperFuncName == mapperFuncName {
			return m
		}
	}
	return nil
}

// MapperFuncFields returns a list of [MapperFuncField] .
func (c *MappingContext) MapperFuncFields() []*MapperFuncField {
	return c.mapperFuncFields
}
