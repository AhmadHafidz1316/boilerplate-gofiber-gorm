# BOILERPLATE-GOFIBER-GORM

**Build Fast, Scale Confidently, Innovate Fearlessly**

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-00BFFF?logo=fiber)
![GORM](https://img.shields.io/badge/GORM-latest-green)
![Swagger](https://img.shields.io/badge/Swagger-API-green?logo=swagger)
![License](https://img.shields.io/badge/License-MIT-yellow)

---

## 📑 Table of Contents
- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Testing](#testing)
- [Features](#features)
- [Project Structure](#project-structure)

---

## 🌐 Overview

**boilerplate-gofiber-gorm** is a comprehensive starter template for building scalable, database-driven web APIs using **Go**, **Fiber**, and **GORM**.  
It streamlines the development of RESTful services with built-in support for:

- CRUD operations
- PostgreSQL integration
- Automatic schema migrations
- API documentation via Swagger

---

## 🚀 Why boilerplate-gofiber-gorm?

This project simplifies creating reliable and maintainable web applications.  
Key features include:

- 🛠️ **CRUD Operations**: Prebuilt endpoints (Customer) with validation and structured responses.  
- 🌐 **API Documentation**: Integrated **Swagger UI** for easy API exploration.  
- 🔌 **Database Connectivity**: Seamless PostgreSQL connection with AutoMigrate.  
- ⚙️ **Modular Architecture**: Config, Handlers, Services, Repositories all separated.  
- 📦 **Dependency Management**: Go modules with `go.mod` and `go.sum`.  
- 🚀 **Fast Startup**: Organized routing + middleware setup.

---

## 🛠️ Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) (1.22+ recommended)
- [PostgreSQL](https://www.postgresql.org/) (or Docker)

### Installation

Clone the repository:

```bash
git clone https://github.com/AhmadHafidz1316/boilerplate-gofiber-gorm
cd boilerplate-gofiber-gorm
```

Install dependencies:

```bash
go mod tidy
```

## <b>Usage</b>

Run the project:

```bash
go run ./cmd/server
```

Server runs at:
👉 http://localhost:3000

Swagger docs at:
👉 http://localhost:3000/swagger/index.html

## <b>Testing</b>

Run unit tests:

```bash
go test ./...
```

## ✨ Features

✅ Auto CRUD boilerplate
✅ Fiber v2 HTTP framework
✅ GORM ORM
✅ PostgreSQL integration
✅ Swagger auto-generated docs
✅ Modular, maintainable folder structure

## 📂 Project Structure

```bash
.
├── cmd/
│   └── server/          # Main entrypoint
├── internal/
│   ├── config/          # App config & env
│   ├── db/              # Database connection + migrations
│   ├── domain/          # Entities + interfaces
│   ├── repository/      # DB access layer
│   ├── service/         # Business logic
│   └── transport/http/  # Handlers + routers
├── pkg/                 # Helpers (response wrapper, utils, etc)
├── docs/                # Swagger generated docs
├── .env                 # Environment variables
├── go.mod
└── go.sum
```
