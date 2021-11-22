package types

type BalanceInput struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

type FarmInput struct {
	Address string `json:"address"`
	Index   uint8  `json:"index"`
	User    string `json:"user"`
}

type OperatorInf struct {
	Address string `json:"address" `
}
