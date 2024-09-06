package domain

type AssetStatsCalculator interface {
	Calculate(asset *Asset) AssetStats
	Update(asset *Asset, transaction Transaction) AssetStats
}
