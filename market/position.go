package market

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/vanclief/ez"
	"github.com/vanclief/state/interfaces"
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
	ID               string       `json:"id"`
	Type             PositionType `json:"type"`
	Pair             *Pair        `json:"pair"`
	Open             bool         `json:"open"`
	OpenPrice        float64      `json:"open_price"`
	ClosePrice       float64      `json:"close_price"`
	Quantity         float64      `json:"quantity"`
	Profit           float64      `json:"profit"`
	ProfitPercentage float64      `json:"profit_percentage"`
	OpenDate         time.Time    `json:"open_date"`
	CloseDate        time.Time    `json:"close_date"`
	Trades           []string     `json:"trades"`
	Tag              string       `json:"tag"`
	Data             string       `json:"data"`
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
		ID:        xid.New().String(),
		Type:      positionType,
		Pair:      trade.Pair,
		OpenPrice: trade.Price,
		Quantity:  trade.Quantity,
		Open:      true,
		OpenDate:  trade.ExecutionDate,
		Trades:    []string{trade.ID},
	}

	return position
}

// GetSchema returns the database schema for the Position model
func (p *Position) GetSchema() *interfaces.Schema {
	return &interfaces.Schema{Name: "positions", PKey: "id"}
}

// GetID returns the ID from the Position model
func (p *Position) GetID() string {
	return p.ID
}

// Update sets the value of the Position instance
func (p *Position) Update(i interface{}) error {
	const op = "Position.Update"

	position, ok := i.(*Position)
	if !ok {
		return ez.New(op, ez.EINVALID, "Provided interface is not of type Position", nil)
	}

	*p = *position

	return nil
}

func (p *Position) String() string {
	return fmt.Sprintf("ID: %s Type: %s Pair: %s Open: %t OpenPrice: %.2f ClosePrice: %.2f Quantity: %.4f Profit: %.2f OpenDate: %s CloseDate: %s # Trades: %d", p.ID[0:8], p.Type, p.Pair.String(), p.Open, p.OpenPrice, p.ClosePrice, p.Quantity, p.Profit, p.OpenDate, p.CloseDate, len(p.Trades))
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
	p.Trades = append(p.Trades, trade.ID)

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

	p.Trades = append(p.Trades, trade.ID)

	return nil
}
