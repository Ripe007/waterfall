package abis

import (
	_ "embed"
)

// //go:embed *.json
// var ABI embed.FS

//go:embed CampaignContinuousCycles.json
var CampaignContinuousCycles []byte

//go:embed IAlpacaLike.json
var IAlpacaLike []byte

//go:embed ICompoundLike.json
var ICompoundLike []byte

//go:embed IERC20MintBurn.json
var IERC20MintBurn []byte

//go:embed Token.json
var TOken []byte
