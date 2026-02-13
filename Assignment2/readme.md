# Assignment 2 — Banking System API

A REST API that replicates the workings of a banking system, built with **Go**, **Gin**, **GORM**, and **PostgreSQL**.

---

## Problem Statement

Design a system that replicates a working of a banking system. It will contain multiple banks and branches, a customer can open a savings account or can take a loan at an interest rate of 12%, can view his account, transactions and loan details (such as loan pending, interest to be paid this year etc), a customer can perform various actions such as deposit and withdraw cash, repay loan, take a loan, etc.

---

## Features

- **Banks & Branches**: Create and manage multiple banks, each with multiple branches
- **Customer Management**: Register customers with unique email, full CRUD
- **Savings Accounts**: Open accounts at any branch, track balance
- **Joint Accounts**: Multiple customers can co-hold a single account
- **Deposits & Withdrawals**: With balance validation (insufficient balance check)
- **Loans**: Take loans at a fixed 12% annual interest rate
- **Loan Repayment**: Partial or full repayment, capped at remaining amount
- **Loan Details**: View pending amount, yearly interest, and total payable
- **Transaction History**: Auto-recorded on every deposit, withdrawal, and loan repayment

---

## Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go 1.25 |
| Framework | Gin (HTTP router) |
| ORM | GORM |
| Database | PostgreSQL |
| Config | godotenv (.env file) |

---

## Project Structure

```
Assignment2/
├── main.go                  # Entry point, starts Gin server
├── .env                     # Environment variables (DB_URL, PORT)
├── go.mod / go.sum          # Go module dependencies
├── initializers/
│   ├── connectToDB.go       # PostgreSQL connection via GORM
│   └── loadEnvVars.go       # Load .env using godotenv
├── models/
│   ├── bank.go              # Bank model
│   ├── branch.go            # Branch model (belongs to Bank)
│   ├── customer.go          # Customer model (unique email)
│   ├── account.go           # Account model (savings, balance)
│   ├── loan.go              # Loan model (12% interest)
│   ├── transaction.go       # Transaction model (deposit/withdrawal/loan_repayment)
│   └── jointAccountHolder.go # Join table for joint accounts
├── services/
│   ├── bankService.go       # Bank CRUD
│   ├── branchService.go     # Branch CRUD + by bank
│   ├── customerService.go   # Customer CRUD
│   ├── accountService.go    # Account CRUD, deposit, withdraw, joint holders
│   ├── loanService.go       # Loan CRUD, repayment, interest calculation
│   └── transactionService.go # Transaction listing
├── controllers/
│   ├── bankController.go
│   ├── branchController.go
│   ├── customerController.go
│   ├── accountController.go  # Includes joint account handlers
│   ├── loanController.go
│   └── transactionController.go
├── routes/
│   └── routes.go             # All API route definitions
└── migrate/
    └── migrate.go            # Auto-migrate all models to PostgreSQL
```

---

## Database ER Diagram

```
┌──────────┐       ┌───────────┐       ┌──────────────┐
│   BANK   │1─────*│  BRANCH   │1─────*│   ACCOUNT    │
│──────────│       │───────────│       │──────────────│
│ id (PK)  │       │ id (PK)   │       │ id (PK)      │
│ name     │       │ bank_id   │       │ customer_id  │
│          │       │ name      │       │ branch_id    │
│          │       │ address   │       │ account_type │
│          │       │           │       │ balance      │
└──────────┘       └───────────┘       └──────┬───────┘
                                              │1
                          ┌───────────────────┤
                          │                   │
                         *│                  *│
                   ┌──────┴───────┐   ┌───────┴──────────┐
                   │ TRANSACTION  │   │ JOINT_ACCOUNT_   │
                   │──────────────│   │ HOLDER           │
                   │ id (PK)      │   │──────────────────│
                   │ account_id   │   │ id (PK)          │
                   │ type         │   │ account_id       │
                   │ amount       │   │ customer_id      │
                   │ description  │   │ is_primary       │
                   └──────────────┘   └──────────────────┘

┌────────────┐
│  CUSTOMER  │1─────*┌─────────────────────┐
│────────────│       │   LOAN              │
│ id (PK)    │       │─────────────────────│
│ name       │       │ id (PK)             │
│ email (UK) │       │ customer_id         │
│ phone      │       │ branch_id           │
│            │       │ amount              │
│            │       │ interest_rate (12%) │    
│            │       │ remaining_amount    │
└────────────┘       └─────────────────────┘

Relationships:
  Bank     1 ──* Branch
  Branch   1 ──* Account
  Branch   1 ──* Loan
  Customer 1 ──* Account (primary holder)
  Customer 1 ──* Loan
  Account  1 ──* Transaction
  Account  * ──* Customer (via JointAccountHolder)
```

---

## Setup & Installation

### Prerequisites
- Go 1.25+
- PostgreSQL running locally

### Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/busychirag/assignments.git
   cd assignments/Assignment2
   ```

2. **Configure environment** — create/edit `.env`:
   ```
   PORT = "3000"
   DB_URL = "host=localhost user=postgres password=YOUR_PASSWORD dbname=postgres port=5432"
   ```

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run database migration** (creates all tables):
   ```bash
   go run migrate/migrate.go
   ```

5. **Start the server**:
   ```bash
   go run main.go
   ```
   Server runs at `http://localhost:3000`

---

## API Documentation

### Banks

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|-------------|
| GET | `/api/banks` | List all banks | — |
| POST | `/api/banks` | Create a bank | `{"name": "HDFC Bank"}` |
| GET | `/api/bank/:id` | Get bank by ID (with branches) | — |
| PUT | `/api/bank/:id` | Update bank | `{"name": "Updated Name"}` |
| DELETE | `/api/bank/:id` | Delete bank | — |
| GET | `/api/bank/:id/branches` | List branches of a bank | — |

---

### Branches

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|-------------|
| GET | `/api/branches` | List all branches | — |
| POST | `/api/branches` | Create a branch | `{"name": "Main Branch", "address": "Mumbai", "bank_id": 1}` |
| GET | `/api/branch/:id` | Get branch by ID | — |
| PUT | `/api/branch/:id` | Update branch | `{"name": "Updated", "address": "Delhi"}` |
| DELETE | `/api/branch/:id` | Delete branch | — |

---

### Customers

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|-------------|
| GET | `/api/customers` | List all customers | — |
| POST | `/api/customers` | Create a customer | `{"name": "John Doe", "email": "john@test.com", "phone": "9999999999"}` |
| GET | `/api/customer/:id` | Get customer (with accounts & loans) | — |
| PUT | `/api/customer/:id` | Update customer | `{"name": "John Updated"}` |
| DELETE | `/api/customer/:id` | Delete customer | — |
| GET | `/api/customer/:id/accounts` | List customer's accounts (incl. joint) | — |
| GET | `/api/customer/:id/loans` | List customer's loans | — |

---

### Accounts

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|-------------|
| GET | `/api/accounts` | List all accounts | — |
| POST | `/api/accounts` | Open a savings account | `{"customer_id": 1, "branch_id": 1}` |
| GET | `/api/account/:id` | Get account details (with holders) | — |
| POST | `/api/account/:id/deposit` | Deposit cash | `{"amount": 10000}` |
| POST | `/api/account/:id/withdraw` | Withdraw cash | `{"amount": 2000}` |
| GET | `/api/account/:id/transactions` | View transaction history | — |

**Business Rules:**
- New accounts are type `savings` with balance `0`
- Deposit amount must be > 0
- Withdrawal amount must be > 0 and ≤ current balance
- Every deposit/withdrawal creates a transaction record automatically

---

### Joint Accounts

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|-------------|
| GET | `/api/account/:id/holders` | List all holders of an account | — |
| POST | `/api/account/:id/holders` | Add a joint holder | `{"customer_id": 2}` |
| DELETE | `/api/account/:id/holders` | Remove a joint holder | `{"customer_id": 2}` |

**Business Rules:**
- The account creator is auto-registered as the **primary holder**
- Joint holders (non-primary) can be added or removed
- **Cannot remove the primary holder**
- **Cannot add duplicate holders**
- When querying a customer's accounts, joint-held accounts are included

---

### Loans

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|-------------|
| GET | `/api/loans` | List all loans | — |
| POST | `/api/loans` | Take a loan (12% interest) | `{"customer_id": 1, "branch_id": 1, "amount": 50000}` |
| GET | `/api/loan/:id` | Get loan details | — |
| POST | `/api/loan/:id/repay` | Repay loan (partial/full) | `{"amount": 5000}` |

**Business Rules:**
- Interest rate is fixed at **12% per annum**
- Loan amount must be > 0
- Repayment is capped at `remaining_amount` (cannot overpay)
- Fully repaid loans cannot be repaid again
- Repayments are recorded as `loan_repayment` transactions

**Loan Details Response** (`GET /api/loan/:id`):
```json
{
  "data": {
    "amount": 50000,
    "interest_rate": 12,
    "remaining_amount": 45000,
    "loan_pending": 45000,
    "interest_this_year": 5400,
    "total_payable": 50400,
    "customer": { "name": "John Doe", ... },
    "branch": { "name": "Main Branch", ... }
  }
}
```

---

### Transactions

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/account/:id/transactions` | List all transactions for an account |

**Transaction Types:**
- `deposit` — cash deposited
- `withdrawal` — cash withdrawn
- `loan_repayment` — loan repayment recorded against account

Transactions are sorted by most recent first.

---

## Example Workflow (Using Postman)

> **Base URL**: `http://localhost:3000`
> For all POST/PUT/DELETE requests, set the **Header**: `Content-Type: application/json`

### Step 1 — Create a Bank

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/banks` |
| **Body (raw JSON)** | 👇 |

```json
{
  "name": "HDFC Bank"
}
```

---

### Step 2 — Create a Branch

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/branches` |
| **Body (raw JSON)** | 👇 |

```json
{
  "name": "Main Branch",
  "address": "Mumbai",
  "bank_id": 1
}
```

---

### Step 3 — Create Customers

**Customer 1 (John):**

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/customers` |
| **Body (raw JSON)** | 👇 |

```json
{
  "name": "John Doe",
  "email": "john@test.com",
  "phone": "9999999999"
}
```

**Customer 2 (Jane):**

Same URL and method, change body to:
```json
{
  "name": "Jane Doe",
  "email": "jane@test.com",
  "phone": "8888888888"
}
```

---

### Step 4 — Open a Savings Account

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/accounts` |
| **Body (raw JSON)** | 👇 |

```json
{
  "customer_id": 1,
  "branch_id": 1
}
```

> Account is created with type `savings` and balance `0` automatically.

---

### Step 5 — Add Jane as Joint Holder on Account 1

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/account/1/holders` |
| **Body (raw JSON)** | 👇 |

```json
{
  "customer_id": 2
}
```

---

### Step 6 — Deposit ₹10,000

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/account/1/deposit` |
| **Body (raw JSON)** | 👇 |

```json
{
  "amount": 10000
}
```

> **Expected**: Balance becomes `10000`

---

### Step 7 — Withdraw ₹2,000

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/account/1/withdraw` |
| **Body (raw JSON)** | 👇 |

```json
{
  "amount": 2000
}
```

> **Expected**: Balance becomes `8000`

---

### Step 8 — Take a Loan of ₹50,000

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/loans` |
| **Body (raw JSON)** | 👇 |

```json
{
  "customer_id": 1,
  "branch_id": 1,
  "amount": 50000
}
```

> **Expected**: Loan created with `interest_rate: 12` and `remaining_amount: 50000`

---

### Step 9 — View Loan Details

| Field | Value |
|-------|-------|
| **Method** | `GET` |
| **URL** | `http://localhost:3000/api/loan/1` |
| **Body** | None |

> **Expected Response** includes:
> - `loan_pending`: 50000
> - `interest_this_year`: 6000 (12% of 50000)
> - `total_payable`: 56000

---

### Step 10 — Repay ₹5,000 on Loan

| Field | Value |
|-------|-------|
| **Method** | `POST` |
| **URL** | `http://localhost:3000/api/loan/1/repay` |
| **Body (raw JSON)** | 👇 |

```json
{
  "amount": 5000
}
```

> **Expected**: `remaining_amount` drops to `45000`

---

### Step 11 — View Transaction History

| Field | Value |
|-------|-------|
| **Method** | `GET` |
| **URL** | `http://localhost:3000/api/account/1/transactions` |
| **Body** | None |

> **Expected**: 3 transactions — `deposit`, `withdrawal`, `loan_repayment` (most recent first)

---

### Step 12 — View Joint Account Holders

| Field | Value |
|-------|-------|
| **Method** | `GET` |
| **URL** | `http://localhost:3000/api/account/1/holders` |
| **Body** | None |

> **Expected**: Shows John (primary) and Jane (joint holder)