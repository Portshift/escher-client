module github.com/Portshift/escher-client

go 1.16

require (
	github.com/EscherAuth/escher v0.0.0-20200415232717-f57476610940
	github.com/Portshift/escher-client v0.0.0-00010101000000-000000000000
	github.com/go-openapi/errors v0.20.1
	github.com/go-openapi/runtime v0.21.0
	github.com/go-openapi/strfmt v0.21.1
	github.com/go-openapi/swag v0.19.15
	github.com/go-openapi/validate v0.20.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.9.0
)

replace (
	github.com/Portshift/escher-client => ../escher-client
	github.com/go-openapi/errors => github.com/go-openapi/errors v0.19.4
	github.com/go-openapi/runtime => github.com/go-openapi/runtime v0.19.15
	github.com/go-openapi/strfmt => github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/validate => github.com/go-openapi/validate v0.19.5
)
