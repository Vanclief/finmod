package market

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/vanclief/ez"
)

// PositionType determines if its a long or short position
type PositionType string

const (
	// LongPosition buy the asset
	LongPosition PositionType = "long"
	// ShortPosition sell the asset
	ShortPosition PositionType = "short"
)

// Position represents a market position
type Position struct {
	ID               xid.ID
	Type             PositionType
	Pair             *Pair
	Open             bool
	OpenPrice        float64
	ClosePrice       float64
	Quantity         float64
	CurrentValue     float64
	Profit           float64
	ProfitPercentage float64
	OpenDate         time.Time
	CloseDate        time.Time
	Trades           []Trade
}

// NewPosition creates a new Position
func NewPosition(trade *Trade) *Position {

	var positionType PositionType
	if trade.Action == BuyAction {
		positionType = LongPosition
	} else if trade.Action == SellAction {
		positionType = ShortPosition
	}

	position := &Position{
		ID:           xid.New(),
		Type:         positionType,
		Pair:         trade.Pair,
		OpenPrice:    trade.Price,
		Quantity:     trade.Quantity,
		CurrentValue: trade.Price * trade.Quantity,
		Open:         true,
		OpenDate:     trade.ExecutionDate,
		Trades:       []Trade{*trade},
	}

	return position
}

func (p *Position) String() string {
	return fmt.Sprintf("ID: %s Type: %s Pair: %s Open: %t OpenPrice: %.2f ClosePrice: %.2f Quantity: %.4f Profit: %.2f OpenDate: %s CloseDate: %s # Trades: %d", p.ID.String()[0:8], p.Type, p.Pair.String(), p.Open, p.OpenPrice, p.ClosePrice, p.Quantity, p.Profit, p.OpenDate, p.CloseDate, len(p.Trades))
}

// Update updates the P\L of a position
func (p *Position) Update(candle *Candle) error {

	if p.Type == LongPosition {
		p.Profit = (candle.Close - p.OpenPrice) * p.Quantity
	} else {
		p.Profit = (p.OpenPrice - candle.Close) * p.Quantity
	}

	p.CurrentValue = candle.Close * p.Quantity

	return nil
}

// Modify receives a new trade that updates the position
func (p *Position) Modify(trade *Trade) error {
	const op = "position.Modify"

	if p.Pair != trade.Pair {
		return ez.New(op, ez.EINVALID, "A trade must have same pair as the position", nil)
	}

	if p.Type == LongPosition && trade.Action == BuyAction {
		p.add(trade)
	} else if p.Type == LongPosition && trade.Action == SellAction {
		p.substract(trade)
	} else if p.Type == ShortPosition && trade.Action == BuyAction {
		p.substract(trade)
	} else if p.Type == ShortPosition && trade.Action == SellAction {
		p.add(trade)
	}

	return nil
}

// add increments the position size and averages the price
func (p *Position) add(trade *Trade) error {

	totalQuantity := p.Quantity + trade.Quantity

	p.OpenPrice = (p.OpenPrice * p.Quantity / totalQuantity) + (trade.Price * trade.Quantity / totalQuantity)
	p.Quantity = totalQuantity
	p.Trades = append(p.Trades, *trade)

	return nil
}

// substract decrements the position size and averages the price
func (p *Position) substract(trade *Trade) error {
	const op = "Position.substract"

	totalQuantity := p.Quantity - float64(trade.Quantity)

	if totalQuantity < 0 && totalQuantity < -0.0001 {
		return ez.New(op, ez.EINVALID, "Trade volume is larger than the position size, this would not only close the position but open a new opossite one", nil)
	} else if totalQuantity <= 0 {
		// Close the trade
		p.ClosePrice = trade.Price
		if p.Type == LongPosition {
			p.Profit = (p.ClosePrice - p.OpenPrice) * p.Quantity
		} else {
			p.Profit = (p.OpenPrice - p.ClosePrice) * p.Quantity
		}

		p.Open = false
		p.CloseDate = trade.ExecutionDate

	} else {
		// Update the trade
		p.OpenPrice = (p.OpenPrice - (trade.Price * (trade.Quantity / p.Quantity))) * p.Quantity / totalQuantity
		p.Quantity = totalQuantity
	}

	p.Trades = append(p.Trades, *trade)

	return nil
}
