```markdown
# Expense Tracker CLI

A simple command-line application built with Go to track expenses. The application allows you to add expenses and view them in a tabular format.

## Features

- Add new expenses with a product name, cost, and the current date.
- View all recorded expenses in a formatted table.
- Keeps track of the total amount remaining after expenses.

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Achanandhi-M/Expense-Tracker-using-GoLang.git
   cd Expense-Tracker-using-GoLang
   ```

2. **Install dependencies:**
   Ensure you have Go installed on your machine. The required dependencies are managed using `go.mod` and `go.sum` files. Simply run:

   ```bash
   go mod tidy
   ```

3. **Build the project:**
   Compile the project into an executable:

   ```bash
   go build -o expense-tracker
   ```

## Usage

1. **Run the application:**
   Execute the program using the following command:

   ```bash
   ./expense-tracker
   ```

2. **Adding an expense:**
   When prompted, enter the product name and the amount spent.

3. **Viewing expenses:**
   Choose the option to view expenses to see a table of all recorded expenses.

## Example

```bash
What do you want to do?
1. Add new expense
2. View expenses
2
+---------+-------+---------------------+
| PRODUCT | COST  |        DATE         |
+---------+-------+---------------------+
| Busfare | 30.00 | 2024-09-03 16:02:50 |
| Milk    | 24.00 | 2024-09-03 16:03:16 |
+---------+-------+---------------------+
Total amount remaining: 9946.00
```

## Project Structure

- **`main.go`**: The main entry point of the application.
- **`expenses.json`**: The JSON file where all the expenses and remaining amount are stored.
- **`go.mod`**: Dependency management file.
- **`go.sum`**: Checksums for dependencies to ensure integrity.

## Dependencies

- [Viper](https://github.com/spf13/viper) for managing configuration.
- [Tablewriter](https://github.com/olekukonko/tablewriter) for displaying expenses in a table format.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

