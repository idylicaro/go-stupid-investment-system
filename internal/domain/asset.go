package domain

// Asset represents an financial asset that can be part of an investment portfolio

type Asset struct {
	ID           int
	Name         string
	Code         string
	Category     string
	Description  string
	CurrentPrice float64

	Stats        AssetStats
	Transactions []Transaction

	StatsCalculator AssetStatsCalculator
}

type AssetStats struct {
	Quantity          float64
	AveragePrice      float64
	Balance           float64
	Variation         float64
	TotalTransactions int
}

func NewAsset(id int, name, code, category, description string, currentPrice float64) Asset {
	return Asset{
		ID:              id,
		Name:            name,
		Code:            code,
		Category:        category,
		Description:     description,
		CurrentPrice:    currentPrice,
		Stats:           AssetStats{},
		Transactions:    []Transaction{},
		StatsCalculator: &DefaultAssetStatsCalculator{},
	}
}

func (a *Asset) AddTransaction(transaction Transaction) {
	a.Transactions = append(a.Transactions, transaction)
	a.UpdateStats(transaction)
}

func (a *Asset) RemoveTransaction(transaction Transaction) {
	for i, t := range a.Transactions {
		if t.ID == transaction.ID {
			// remove transaction from slice
			a.Transactions = append(a.Transactions[:i], a.Transactions[i+1:]...)
			break
		}
	}
	a.CalculateStats()

}

func (a *Asset) UpdateStats(transaction Transaction) {
	if a.StatsCalculator != nil {
		a.Stats = a.StatsCalculator.Update(a, transaction)
	}
}

func (a *Asset) CalculateStats() {
	if a.StatsCalculator != nil {
		a.Stats = a.StatsCalculator.Calculate(a)
	}
}
