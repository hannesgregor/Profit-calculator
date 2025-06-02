package main

import (
	"fmt"
	"math/rand"
	"time"
)

// define what a transaction should look like
type transaction struct {
	transaction_type string
	quantity         int
	price            float64
	fee              float64
	timestamp        string
}

// define a basic portfolio that holds transactions and balance
type portfolio struct {
	transactions []transaction
	balance      float64
}

// function to add transactions to the portfolio and update the balance
func (p *portfolio) addTransaction(t transaction) {
	if t.transaction_type == "buy" {
		cost := t.price*float64(t.quantity) + t.fee
		if p.balance < cost {
			// Not enough balance. In this case, we just skip the transaction.
			// In a real-world scenario, we might want to handle this differently, e.g., throw an error or log it.
			fmt.Println("Insufficient balance for transaction:", t)
			return
		}
		p.balance -= cost
	} else if t.transaction_type == "sell" {
		p.balance += (t.price*float64(t.quantity) - t.fee)
	}
	p.transactions = append(p.transactions, t)
}

// function to generate a random transaction
func randomTransaction() transaction {
	types := []string{"buy", "sell"}
	t := types[rand.Intn(len(types))]
	qty := rand.Intn(100) + 1
	price := rand.Float64()*100 + 1
	fee := rand.Float64()*5 + 0.1
	ts := time.Now().Add(time.Duration(-rand.Intn(1000000)) * time.Second).Format(time.RFC3339)
	return transaction{
		transaction_type: t,
		quantity:         qty,
		price:            price,
		fee:              fee,
		timestamp:        ts,
	}
}

// generate a list of random transactions (n amount)
func generateTransactions(n int) []transaction {
	transactions := make([]transaction, 0, n)
	for i := 0; i < n; i++ {
		transactions = append(transactions, randomTransaction())
	}
	return transactions
}

// this function calculates the profitability of the portfolio based on the transactions
// note that in this case we start with 0 balance but in a real-world scenario, we might want to start with an initial balance.
func calculateProfitability(transactions []transaction) float64 {
	var totalBuys, totalSells float64
	for _, t := range transactions {
		if t.transaction_type == "buy" {
			totalBuys += t.price*float64(t.quantity) + t.fee
		} else if t.transaction_type == "sell" {
			totalSells += t.price*float64(t.quantity) - t.fee
		}
	}
	return totalSells - totalBuys
}

func main() {
	p := portfolio{}
	transactions := generateTransactions(1000)
	for _, t := range transactions {
		p.addTransaction(t)
	}
	fmt.Printf("Generated %d transactions\n", len(transactions))
	fmt.Printf("Final balance: %.2f\n", p.balance)
	fmt.Printf("Profitability: %.2f\n", calculateProfitability(transactions))

	// too many (1000) transactions to print, but you can uncomment the next line to see them
	// fmt.Println(transactions)

}
