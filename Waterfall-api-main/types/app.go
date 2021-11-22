package types

import ()

type AppConfig struct {
	Port                            string   `yaml:"port"`
	PlatformAddress                 string   `yaml:"plaform_address"`
	CampaignContinuousCyclesAddress string   `yaml:"campaignContinuousCycleAddress"`
	RPCs                            []string `yaml:"rpc"`
}

type AppResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type TokenInfo struct {
	Symbol   string `bson:"symbol" json:"symbol" yaml:"symbol"`
	Address  string `bson:"address" json:"address" yaml:"address"`
	Decimals uint8  `bson:"decimal" json:"decimal" yaml:"decimal"`
	Icon     string `bson:"icon,omitempty" json:"icon,omitempty" yaml:"icon,omitempty"`
}
