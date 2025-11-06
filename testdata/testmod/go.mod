module example.com/testmod

go 1.25

require github.com/google/go-cmp v0.5.9

require (
	github.com/google/uuid v1.6.0
	github.com/oapi-codegen/runtime v1.1.2
	github.com/roryq/sesame v0.0.0
)

replace github.com/roryq/sesame => ../../
