# Go Blog

![Build Status](https://github.com/wkdwilliams/go-blog/actions/workflows/delpoy.yml/badge.svg)

[Live Demo](https://lewiswilliams.info/)

---

## Introduction

Go Blog is a scalable blog platform built with the Go programming language, following a hexagonal architecture (ports and adapters). This design pattern ensures that the application is maintainable, adaptable, and ready for future scaling. The frontend is rendered using pure HTML, CSS, and JS, with no reliance on JavaScript frameworks.

Key technologies used:

1. Frontend Rendering: [Templ](https://templ.guide/) – A templating engine for Go.
2. Database ORM: [Gorm](https://gorm.io/) - An ORM library for Go, facilitating smooth database interactions.
3. Database Migrations: [Migrations](https://github.com/golang-migrate/migrate) - Handles versioned migrations and database schema changes.

---

## Getting Started

### Setting Up Environment Variables

To begin, copy the .env.example file to .env and modify the variables according to your local setup:

```bash
cp .env.example .env
```

### Deploying for local development

#### 1. Launching the MySQL Database & phpMyAdmin

The application uses Docker to manage services like MySQL and phpMyAdmin. To set up the environment for development:

```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

#### 2. Running the Application

After the services are up, follow these steps to get the application running:

Run Database Migrations:
Apply database migrations to set up the initial schema:

```bash
make migrate-up
```

#### 2. Seed the Database:

Populate the database with initial data:

```bash
make seed
```

#### 3. Start the Application:

Run the application in development mode:

```bash
make dev
```

You can now access the admin dashboard at http://localhost:8000/admin with the default login credentials:

```
username: admin
Password: pass
```

All the essential commands for development, such as migration and database seeding, can be found in the Makefile.

---

### Production Deployment

For deploying the application in a production environment, use the following command:

```bash
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

Ensure that the environment variables in the .env file are properly configured for the production setup.

### Makefile Commands

The project includes a Makefile with useful commands for development and deployment tasks:

1. **make migrate-up** – Apply database migrations.
2. **make seed** – Seed the database with default data.
3. **make dev** – Run the application in development mode.
4. **make run** - Runs the application in production mode
5. **make compile** - compiles the application to a binary
6. **make migrate-create** - Creates a new migration
7. **make migrate-up** - Executes the migrate up command
8. **make migrate-down** - Executes the migrate down command

Refer to the Makefile for more commands related to testing, building, and other development processes.

## Technologies Used

-   **Go** – Backend programming language.
-   **Templ** – Go-based templating engine for rendering frontend views.
-   **Gorm** – Database access and ORM library.
-   **Docker** – Containerization for development and production environments.
-   **Migrate** – Database migration management.
