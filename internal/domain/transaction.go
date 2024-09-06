package domain

type TransactionType string

const (
	Buy  TransactionType = "BUY"
	Sell TransactionType = "SELL"
)

type Transaction struct {
	ID       int
	Quantity float64
	Price    float64
	Date     string
	Type     TransactionType
}

func NewTransaction(id int, quantity, price float64, date string, tType TransactionType) Transaction {
	return Transaction{
		ID:       id,
		Quantity: quantity,
		Price:    price,
		Date:     date,
		Type:     tType,
	}
}
