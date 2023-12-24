# Credit App Go

a simple Go application..


## Prerequisites

Before running the application, make sure you have Docker and Go installed on your machine.

## Getting Started

1. Start the PostgreSQL container by running the following Docker command:

    ```bash
    docker run --name postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
    ```

   This command will pull the PostgreSQL image (if not available locally) and start a container named `postgres` with the specified password.

2. Run the following command to setup and migrate the SQL tables for the project:

    ```bash
    make migrate-setup
    ```

   This command initializes the database and applies the initial migrations using the `go migrate` tool.

3. Create the necessary SQL tables by running:

    ```bash
    make migrate-up
    ```

   This command specifically applies the database migrations and creates the required tables.

5. Build the Go project by running:

    ```bash
    make build
    ```

   This command compiles the Go code and creates an executable binary.

6. Run the application using the following command:

    ```bash
    make run
    ```

   This command executes the compiled binary and starts the App.

## Admin User

The admin user is automatically added to the database during the migration process.

- Email: admin@admin.com
- Password: admin

The admin user has full access to view all user informations and transactions.

## Cleanup

To stop and remove the PostgreSQL container, run the following command:

```bash
docker stop postgres && docker rm postgres

```

---

## Features

- **Login and Register:**
  - Users can create accounts and log in to the system securely.

- **Money Transfer:**
  - Users can transfer money to each other securely.

- **Transaction History:**
  - Users can query their transaction history to see all past transactions.

- **Custom Transaction Queries:**
  - Users can perform custom queries to retrieve transaction history for a specific date and time.

- **Role-Based Access Control:**
  - User role: Users can log in, transfer money, and query their own transaction history.
  - Admin Role: Admin users have full access to view all user informations and transactions.
