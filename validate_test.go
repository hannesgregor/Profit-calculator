package main

import "testing"

// Unit test for zero transactions. Make sure that the balance and profit is zero.
func TestProfitability_ZeroTransactions(t *testing.T) {
	var transactions []transaction
	got := calculateProfitability(transactions)
	if got != 0 {
		t.Errorf("Expected 0, got %v", got)
	}
}

// unit test for an attempt to buy with insufficient balance
func TestAddTransaction_InsufficientBalance(t *testing.T) {
	p := portfolio{}
	t1 := transaction{
		transaction_type: "buy",
		quantity:         10,
		price:            100.0,
		fee:              5.0,
		timestamp:        "2023-10-01T00:00:00Z",
	}

	// Attempt to add a buy transaction with insufficient balance
	p.addTransaction(t1)

	if p.balance != 0 {
		t.Errorf("Expected balance to be 0, got %v", p.balance)
	}
}

// unit test for transactions with negative quanitity, price or fee
func TestAddTransaction_NegativeValues(t *testing.T) {
	p := portfolio{}
	t1 := transaction{
		transaction_type: "buy",
		quantity:         -10,
		price:            -100.0,
		fee:              -5.0,
		timestamp:        "2023-10-01T00:00:00Z",
	}

	// Attempt to add a transaction with negative values
	p.addTransaction(t1)

	if p.balance != 0 {
		t.Errorf("Expected balance to be 0, got %v", p.balance)
	}
}

// Test only sell transactions (should increase balance and profit)
func TestOnlySellTransactions(t *testing.T) {
	p := portfolio{}
	transactions := []transaction{
		{transaction_type: "sell", quantity: 5, price: 100.0, fee: 2.0, timestamp: "2023-10-01T00:00:00Z"},
		{transaction_type: "sell", quantity: 3, price: 50.0, fee: 1.0, timestamp: "2023-10-02T00:00:00Z"},
	}
	for _, tr := range transactions {
		p.addTransaction(tr)
	}
	profit := calculateProfitability(transactions)
	expectedProfit := (5*100.0 - 2.0) + (3*50.0 - 1.0)
	if p.balance != expectedProfit {
		t.Errorf("Expected balance %v, got %v", expectedProfit, p.balance)
	}
	if profit != expectedProfit {
		t.Errorf("Expected profit %v, got %v", expectedProfit, profit)
	}
}

// Test only buy transactions (should decrease balance and profit)
func TestOnlyBuyTransactions(t *testing.T) {
	p := portfolio{balance: 10000} // Give enough balance to allow buys
	transactions := []transaction{
		{transaction_type: "buy", quantity: 5, price: 100.0, fee: 2.0, timestamp: "2023-10-01T00:00:00Z"},
		{transaction_type: "buy", quantity: 3, price: 50.0, fee: 1.0, timestamp: "2023-10-02T00:00:00Z"},
	}
	for _, tr := range transactions {
		p.addTransaction(tr)
	}
	profit := calculateProfitability(transactions)
	expectedLoss := (5*100.0 + 2.0) + (3*50.0 + 1.0)
	if p.balance != 10000-expectedLoss {
		t.Errorf("Expected balance %v, got %v", 10000-expectedLoss, p.balance)
	}
	if profit != -expectedLoss {
		t.Errorf("Expected profit %v, got %v", -expectedLoss, profit)
	}
}
