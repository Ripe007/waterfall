module waterfall-contract-api

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.4 // indirect
	github.com/labstack/echo/v4 v4.3.0 // indirect
	github.com/waterfall/contracts v0.0.0-00010101000000-000000000000 // indirect
	github.com/waterfall/types v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect

)

replace github.com/waterfall/types => ./types

replace github.com/waterfall/contracts => ./contracts
