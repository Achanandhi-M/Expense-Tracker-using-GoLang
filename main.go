package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"os"
	"time"
)

const configFile = "expenses.json"

type ExpenseData struct {
	Product string  `json:"product"`
	Cost    float64 `json:"cost"`
	Date    string  `json:"date"`
}

type OutputData struct {
	Expenses        []ExpenseData `json:"expenses"`
	AmountRemaining float64       `json:"amount_remaining"`
}

// Read existing configuration

func readConfig() (OutputData, error) {
	var config OutputData
	file, err := os.Open(configFile)
	if err != nil {
		return config, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err := dec.Decode(&config); err != nil {
		return config, fmt.Errorf("error decoding config file: %v", err)
	}
	return config, nil
}

func calculateExpense() {
	// Initialize Viper
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")

	// Load configuration
	if err := viper.ReadInConfig(); err != nil && os.IsNotExist(err) {
		// Create default config if not exists
		defaultConfig := OutputData{AmountRemaining: 0.0}
		file, err := os.Create(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating config file: %v\n", err)
			return
		}
		defer file.Close()

		enc := json.NewEncoder(file)
		if err := enc.Encode(defaultConfig); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing default config: %v\n", err)
			return
		}
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
		return
	}

	// Read the current configuration
	config, err := readConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	var product string
	var purchaseAmount float64
	config.AmountRemaining = 10000 // Example starting balance

	fmt.Print("For What you Spent?:")
	_, err = fmt.Scanf("%s", &product)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Amount you spent:")
	_, err = fmt.Scanf("%f", &purchaseAmount)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Capture the current date and time in a formatted string
	currentDate := time.Now().Format("2006-01-02 15:04:05")

	// Add new expense and update total
	newExpense := ExpenseData{
		Product: product,
		Cost:    purchaseAmount,
		Date:    currentDate,
	}
	config.Expenses = append(config.Expenses, newExpense)
	config.AmountRemaining -= purchaseAmount

	// Write updated configuration
	file, err := os.Create(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating config file: %v\n", err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	if err := enc.Encode(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing config file: %v\n", err)
		return
	}

	fmt.Print("Expense added successfully")
}

func viewExpenses() {
	// Read the current configuration
	config, err := readConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print expenses in a table format
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Spent", "Cost", "Date"})

	for _, expense := range config.Expenses {
		table.Append([]string{
			expense.Product,
			fmt.Sprintf("%.2f", expense.Cost),
			expense.Date,
		})
		table.Append([]string{"", "", ""})
	}
	table.SetFooter([]string{"", "Total", fmt.Sprintf("%.2f", config.AmountRemaining)})
	table.Render()
}
func main() {
	var choice string

	fmt.Println("What do you want to do?")
	fmt.Println("1. Add new expense")
	fmt.Println("2. View expenses")
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		calculateExpense()
	case "2":
		viewExpenses()
	default:
		fmt.Println("Invalid choice")
	}
}
