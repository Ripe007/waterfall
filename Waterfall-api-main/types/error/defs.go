package error

import ()

const (
	OK = 200

	BadParams = 400
	Contract  = 401
	Internal  = 402
	ABIPack   = 403

	BadAmount               = 1000
	InsufficientBalance     = 1001
	InvalidTemplateID       = 1002
	FailToCreateTransaction = 1003
	InvalidAuctionID        = 1004
	UnknownToken            = 1005
	UnknownSSID             = 1006
	ExceedQuotaLimits       = 1007
	FailToGetAccountBalance = 1008
	InvalidSSID             = 1009
	NoTokensFound           = 1010
	InvalidToken            = 1011
	NotAllowed              = 1012
	InvalidSessionData      = 1013
)

const (
	BadAmountMsg               = "bad amount"
	BadParamsMsg               = "bad parameters"
	InsufficientBalanceMsg     = "insufficient balance"
	InvalidTemplateIDMsg       = "invalid template id"
	FailToCreateTransactionMsg = "fail to create transaction"
	InvalidAuctionIDMsg        = "invalid auction id"
	UnknownTokenMsg            = "unknown token"
	UnknownSSIDMsg             = "unknown SSID"
	ExceedQuotaLimitsMsg       = "exceed quota limits"
	InvalidSSIDMsg             = "invalid session id"
	NoTokensFoundMsg           = "no tokens found"
	InvalidTokenMsg            = "invalid token"
	NotAllowedMsg              = "not allowed"
)
