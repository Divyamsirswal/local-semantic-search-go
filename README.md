# Go Semantic Search API with Ollama & PGVector

This project is a simple yet powerful REST API built with Go that performs semantic search over a collection of text. It uses Ollama for local text embedding generation and PostgreSQL with the `pgvector` extension for efficient vector similarity search. The entire backend is containerized with Docker for easy setup and deployment.

---
## Tech Stack

- **Backend:** Go (Golang)
- **Web Framework:** Echo
- **Database:** PostgreSQL + `pgvector` extension
- **Embeddings:** Ollama (serving the `mxbai-embed-large` model)
- **Containerization:** Docker & Docker Compose

---
## Features

- **Semantic Search:** Search for text based on meaning and context, not just keywords.
- **REST API:** A simple `GET /search` endpoint to perform searches.
- **Dockerized Environment:** All backend services (Postgres, Ollama) are managed by Docker Compose for a one-command setup.
- **Local First:** Uses the local Ollama service for embeddings, requiring no external API keys or internet dependency for core functionality.

---
## Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.18 or higher)
- [Docker](https://www.docker.com/products/docker-desktop/) and Docker Compose

### Installation & Setup

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/your-repository-name.git](https://github.com/your-username/your-repository-name.git)
    cd your-repository-name
    ```

2.  **Create an environment file:**
    Copy the example `.env` file. No changes are needed if you are running locally.
    ```bash
    cp .env.example .env
    ```

3.  **Start the backend services:**
    This command will start the PostgreSQL and Ollama containers in the background.
    ```bash
    docker-compose up -d
    ```

4.  **Download the embedding model:**
    Tell the running Ollama container to download the required model. This only needs to be done once.
    ```bash
    docker exec -it ollama_service ollama pull mxbai-embed-large
    ```

5.  **Install Go dependencies:**
    ```bash
    go mod tidy
    ```

6.  **Run the Go application:**
    ```bash
    go run main.go
    ```
    The server will start, seed the database with initial data, and be ready to accept requests on `http://localhost:8081`.

---
## API Usage

### Search for Similar Sentences

- **Endpoint:** `GET /search`
- **Query Parameter:** `q` (your search query)
- **Example:**

  ```bash
  curl "http://localhost:8081/search?q=what is a small animal"

Example Response:

[
    {
        "content": "A cat is a small, furry mammal.",
        "similarity": 0.891234567
    },
    {
        "content": "The sun rises in the east.",
        "similarity": 0.751234567
    },
    {
        "content": "Go is an open-source programming language.",
        "similarity": 0.681234567
    }
]

Project Structure
/
├── database/               # Database connection, seeding, and query logic
├── handlers/               # Echo web handlers for API endpoints
├── .env                    # Local environment variables (ignored by Git)
├── .env.example            # Example environment file
├── .gitignore              # Files and folders to ignore in Git
├── docker-compose.yml      # Docker setup for Postgres & Ollama
├── go.mod                  # Go module dependencies
├── init.sql                # Database table initialization script
└── main.go                 # Main application entry point
