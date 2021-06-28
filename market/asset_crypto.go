package market

import "github.com/vanclief/ez"

func NewCryptoAsset(symbol string) (*Asset, error) {
	const op = "market.NewCryptoAsset"

	switch symbol {
	case "AAVE":
		return NewAsset(symbol, "Aave")
	case "ALGO":
		return NewAsset(symbol, "Algorand")
	case "BAT":
		return NewAsset(symbol, "Basic Attention Token")
	case "BTC":
		return NewAsset(symbol, "Bitcoin")
	case "BCH":
		return NewAsset(symbol, "Bitcoin Cash")
	case "ADA":
		return NewAsset(symbol, "Cardano")
	case "LINK":
		return NewAsset(symbol, "Chainlink")
	case "COMP":
		return NewAsset(symbol, "Compound")
	case "ATOM":
		return NewAsset(symbol, "Cosmos")
	case "DAI":
		return NewAsset(symbol, "DAI")
	case "DASH":
		return NewAsset(symbol, "Dash")
	case "DOGE":
		return NewAsset(symbol, "Dogecoin")
	case "EOS":
		return NewAsset(symbol, "EOS.IO")
	case "ETH":
		return NewAsset(symbol, "Ethereum")
	case "ETC":
		return NewAsset(symbol, "Ethereum Classic")
	case "EWT":
		return NewAsset(symbol, "Energy Web Token")
	case "EWTB":
		return NewAsset(symbol, "Energy Web Token Bridged")
	case "FILE":
		return NewAsset(symbol, "Filecoin")
	case "FLOW":
		return NewAsset(symbol, "Flow")
	case "LTC":
		return NewAsset(symbol, "Litecoin")
	case "XLM":
		return NewAsset(symbol, "Lumen")
	case "XMR":
		return NewAsset(symbol, "Monero")
	case "NANO":
		return NewAsset(symbol, "Nano")
	case "MATIC":
		return NewAsset(symbol, "Polygon")
	case "QTUM":
		return NewAsset(symbol, "Qtum")
	case "OCEAN":
		return NewAsset(symbol, "Ocean Token")
	case "DOT":
		return NewAsset(symbol, "Polkadot")
	case "USDT":
		return NewAsset(symbol, "Tether USD")
	case "GRT":
		return NewAsset(symbol, "The Graph")
	case "UNI":
		return NewAsset(symbol, "Uniswap")
	case "USDC":
		return NewAsset(symbol, "USD Coin")
	case "XRP":
		return NewAsset(symbol, "Ripple")
	case "ZED":
		return NewAsset(symbol, "ZCash")
	default:
		return nil, ez.New(op, ez.ENOTFOUND, "No translatable asset found", nil)
	}
}

func NewForexAsset(symbol string) (*Asset, error) {
	const op = "market.NewForexAsset"

	switch symbol {
	case "USD":
		return NewAsset(symbol, "United States Dollar")
	case "EUR":
		return NewAsset(symbol, "Euro")

	default:
		return nil, ez.New(op, ez.ENOTFOUND, "No translatable asset found", nil)
	}
}

func NewStockAsset(symbol string) (*Asset, error) {
	const op = "market.NewStockAsset"

	switch symbol {
	case "AAPL":
		return NewAsset(symbol, "Apple Inc")

	default:
		return nil, ez.New(op, ez.ENOTFOUND, "No translatable asset found", nil)
	}
}
