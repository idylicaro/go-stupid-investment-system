package domain

import (
	"testing"
)

type MockAssetStatsCalculator struct{}

func (m *MockAssetStatsCalculator) Update(a *Asset, t Transaction) AssetStats {
	return AssetStats{
		Quantity:          10,
		AveragePrice:      100,
		Balance:           1000,
		Variation:         5,
		TotalTransactions: 1,
	}
}

func (m *MockAssetStatsCalculator) Calculate(a *Asset) AssetStats {
	return AssetStats{
		Quantity:          10,
		AveragePrice:      100,
		Balance:           1000,
		Variation:         5,
		TotalTransactions: 1,
	}
}

func TestNewAsset(t *testing.T) {
	asset := NewAsset(1, "Test Asset", "TST", "Category", "Description", 100.0)

	if asset.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", asset.ID)
	}
	if asset.Name != "Test Asset" {
		t.Errorf("Expected Name to be 'Test Asset', got %s", asset.Name)
	}
	if asset.Code != "TST" {
		t.Errorf("Expected Code to be 'TST', got %s", asset.Code)
	}
	if asset.Category != "Category" {
		t.Errorf("Expected Category to be 'Category', got %s", asset.Category)
	}
	if asset.Description != "Description" {
		t.Errorf("Expected Description to be 'Description', got %s", asset.Description)
	}
	if asset.CurrentPrice != 100.0 {
		t.Errorf("Expected CurrentPrice to be 100.0, got %f", asset.CurrentPrice)
	}
	if len(asset.Transactions) != 0 {
		t.Errorf("Expected Transactions to be empty, got %d", len(asset.Transactions))
	}
	if asset.StatsCalculator == nil {
		t.Errorf("Expected StatsCalculator to be non-nil")
	}
}

func TestUpdateStats(t *testing.T) {
	asset := NewAsset(1, "Test Asset", "TST", "Category", "Description", 100.0)
	mockCalculator := &MockAssetStatsCalculator{}
	asset.StatsCalculator = mockCalculator

	transaction := Transaction{} // Assuming Transaction struct is defined elsewhere
	asset.UpdateStats(transaction)

	expectedStats := mockCalculator.Update(&asset, transaction)
	if asset.Stats != expectedStats {
		t.Errorf("Expected Stats to be %+v, got %+v", expectedStats, asset.Stats)
	}
}
