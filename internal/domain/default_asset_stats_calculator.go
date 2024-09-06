// default_asset_stats_calculator.go
package domain

type DefaultAssetStatsCalculator struct{}

func (d *DefaultAssetStatsCalculator) Calculate(asset *Asset) AssetStats {
	var quantity, balance, averagePrice, variation float64
	var totalTransactions int

	for _, t := range asset.Transactions {
		if t.Type == Buy {
			quantity += t.Quantity
			balance += t.Quantity * t.Price
		} else if t.Type == Sell {
			quantity -= t.Quantity
			balance -= t.Quantity * t.Price
		}
		totalTransactions++
	}

	if quantity > 0 {
		averagePrice = balance / quantity
	}

	if asset.CurrentPrice > 0 && averagePrice > 0 {
		variation = ((asset.CurrentPrice - averagePrice) / averagePrice) * 100
	}

	return AssetStats{
		Quantity:          quantity,
		AveragePrice:      averagePrice,
		Variation:         variation,
		Balance:           balance,
		TotalTransactions: totalTransactions,
	}
}

func (d *DefaultAssetStatsCalculator) Update(asset *Asset, transaction Transaction) AssetStats {
	newQuantity := asset.Stats.Quantity
	newBalance := asset.Stats.Balance

	if transaction.Type == Buy {
		newQuantity += transaction.Quantity
		newBalance += transaction.Quantity * transaction.Price
	} else if transaction.Type == Sell {
		newQuantity -= transaction.Quantity
		newBalance -= transaction.Quantity * transaction.Price
	}

	var newAveragePrice float64
	if newQuantity > 0 {
		newAveragePrice = newBalance / newQuantity
	}

	var newVariation float64
	if asset.CurrentPrice > 0 && newAveragePrice > 0 {
		newVariation = ((asset.CurrentPrice - newAveragePrice) / newAveragePrice) * 100
	}

	return AssetStats{
		Quantity:          newQuantity,
		AveragePrice:      newAveragePrice,
		Balance:           newBalance,
		Variation:         newVariation,
		TotalTransactions: asset.Stats.TotalTransactions + 1,
	}
}
