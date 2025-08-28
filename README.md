
# Go Semantic Search API 🚀

[![Go version](https://img.shields.io/badge/go-1.18%2B-blue.svg)](https://go.dev)
[![Docker](https://img.shields.io/badge/docker-compose-blue.svg)](https://www.docker.com)
[![PostgreSQL](https://img.shields.io/badge/postgres-pgvector-blue.svg)](https://www.postgresql.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple and powerful REST API built with Go for semantic search. It uses a local **Ollama** instance for text embeddings and **PostgreSQL** with the `pgvector` extension for efficient vector similarity search. The entire backend is containerized with **Docker** for a one-command setup.

-----

## Table of Contents

  - [About The Project](#about-the-project)
  - [Tech Stack](#tech-stack)
  - [Getting Started](#getting-started)
      - [Prerequisites](#prerequisites)
      - [Installation](#installation)
  - [API Usage](#api-usage)
      - [Search Endpoint](#search-endpoint)
      - [Example Request](#example-request)
  - [Project Structure](#project-structure)
  - [License](#license)

-----

## About The Project

This project serves as a practical, hands-on guide for building a semantic search system from scratch. Unlike keyword-based search, semantic search understands the contextual meaning of a query, allowing it to find more relevant results.

The key features are:

  - **Local-First:** Uses the local Ollama service for embeddings, requiring no external API keys or internet dependency for its core functionality.
  - **Containerized:** All backend services (Postgres, Ollama) are managed by Docker Compose, making the setup process simple and reproducible.
  - **Modern Go Backend:** Built with Go and the high-performance Echo framework.

-----

## Tech Stack

  - **Backend:** Go (Golang)
  - **Web Framework:** Echo
  - **Database:** PostgreSQL with the `pgvector` extension
  - **Embeddings:** Ollama (serving the `mxbai-embed-large` model)
  - **Containerization:** Docker & Docker Compose

-----

## Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

  - [Go](https://go.dev/dl/) (version 1.18 or higher is required for generics)
  - [Docker](https://www.docker.com/products/docker-desktop/) and Docker Compose

### Installation

1.  **Clone the repository:**

    ```bash
    git clone [https://github.com/your-username/your-repository-name.git](https://github.com/your-username/your-repository-name.git)
    cd your-repository-name
    ```

2.  **Create an environment file:**
    This project requires a `.env` file for configuration. You can create one from the example file. No changes are needed for the default local setup.

    ```bash
    cp .env.example .env
    ```

3.  **Start the backend services:**
    This command will build and start the PostgreSQL and Ollama containers in the background.

    ```bash
    docker-compose up -d
    ```

4.  **Download the embedding model:**
    Tell the running Ollama container to download the required model. This only needs to be done once, as the model will be saved in a persistent Docker volume.

    ```bash
    docker exec -it ollama_service ollama pull mxbai-embed-large
    ```

5.  **Install Go dependencies:**
    This command will download the necessary Go libraries listed in `go.mod`.

    ```bash
    go mod tidy
    ```

6.  **Run the Go application:**

    ```bash
    go run main.go
    ```

    The server will start, automatically seed the database with initial data, and be ready to accept requests on `http://localhost:8081`.

-----

## API Usage

### Search Endpoint

  - **URL:** `/search`
  - **Method:** `GET`
  - **Query Parameter:** `q=[string]` (The text you want to search for)

### Example Request

You can test the API using `curl` from your terminal:

```bash
curl "http://localhost:8081/search?q=what is a small animal"
````

#### Example Response

The API will return a JSON array of the most similar sentences, ordered by their cosine similarity score.

```json
[
    {
        "content": "A cat is a small, furry mammal.",
        "similarity": 0.8954
    },
    {
        "content": "The sun rises in the east.",
        "similarity": 0.7612
    },
    {
        "content": "Go is an open-source programming language.",
        "similarity": 0.6831
    }
]
```

-----

## Project Structure

```
.
├── database/         # Database connection, queries, and seeding
├── handlers/         # Echo web handlers for API endpoints
├── .env              # Local environment variables (ignored by Git)
├── .env.example      # Example environment file
├── .gitignore        # Files to be ignored by Git
├── docker-compose.yml# Docker setup for Postgres & Ollama
├── go.mod            # Go module definition and dependencies
├── init.sql          # Database table initialization script
└── main.go           # Main application entry point
```

-----

## License

Distributed under the MIT License. See `LICENSE` for more information.
