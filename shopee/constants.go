package shopee

var (
	ContentTypeJson = "application/json"
	AcceptJson      = "application/json"

	DefaultTimeout = 60

	// Production environment
	ApiGateway = "https://partner.shopeemobile.com"
	// APIGateway = "https://openplatform.shopee.cn"     // for developers who deployed their services near Chinese Mainland.
	// APIGateway = "https://openplatform.shopee.com.br" // for developer who deployed their services near US.
	// APIGateway = "https://partner.shopeemobile.com"   // for developer who deployed their services near SG.

	// Sandbox environment
	// ApiGateway = "https://partner.test-stable.shopeemobile.com" // for All developers
)
