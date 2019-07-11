package learngo

import "github.com/shopspring/decimal"

func GetDecimalValue() *decimal.Decimal {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		return nil
	}
	return &price
}
