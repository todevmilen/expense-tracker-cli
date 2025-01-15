# Expense Tracker CLI

Manage your expenses with ease using this simple and efficient Command Line Interface (CLI) tool!

## Features

- Add new expenses effortlessly.
- Delete expenses by their ID.
- View a list of all your expenses.
- Get a summary of total expense amounts.
- Autocompletion support for smooth command usage.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/todevmilen/expense-tracker-cli.git
   cd expense-tracker-cli
   ```

2. Build the project:

   ```bash
   go build -o expense-tracker
   ```

3. Add the executable to your PATH for global usage (optional).

## Usage

Run the CLI tool by executing:

```bash
./expense-tracker [command]
```

### Available Commands

- **add**: Adds a new expense.

  ```bash
  ./expense-tracker add --amount 50 --category food --description "Lunch at cafe"
  ```

- **delete**: Delete an expense by its ID.

  ```bash
  ./expense-tracker delete --id 1
  ```

- **list**: Lists all expenses.

  ```bash
  ./expense-tracker list
  ```

- **summary**: Displays a summary of expenses.

  ```bash
  ./expense-tracker summary
  ```

- **completion**: Generate the autocompletion script for your shell (bash, zsh, etc.).

  ```bash
  ./expense-tracker completion
  ```

- **hello**: Prints "hello world" (example command).

  ```bash
  ./expense-tracker hello
  ```

- **help**: Displays help information for commands.
  ```bash
  ./expense-tracker help
  ```

## Example

Add an expense:

```bash
./expense-tracker add --amount 100 --category travel --description "Taxi fare"
```

List all expenses:

```bash
./expense-tracker list
```

Delete an expense by ID:

```bash
./expense-tracker delete --id 2
```

Get a summary of all expenses:

```bash
./expense-tracker summary
```

## This project is part of [RoadMapSH Projects](https://roadmap.sh/projects/expense-tracker)
