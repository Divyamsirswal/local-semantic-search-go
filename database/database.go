package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ollama/ollama/api"
	"github.com/pgvector/pgvector-go"
)

func toFloat32(in []float64) []float32 {
	out := make([]float32, len(in))
	for i, v := range in {
		out[i] = float32(v)
	}
	return out
}

type SearchResult struct {
	Content    string  `json:"content"`
	Similarity float64 `json:"similarity"`
}

func NewConnection(dbUrl string) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), dbUrl)
}

func SearchSentences(dbpool *pgxpool.Pool, embedding []float64) ([]SearchResult, error) {
	rows, err := dbpool.Query(context.Background(),
		`SELECT content, 1 - (embedding <=> $1) AS similarity
		 FROM sentences
		 ORDER BY similarity DESC
		 LIMIT 5`,
		pgvector.NewVector(toFloat32(embedding)))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var r SearchResult
		if err := rows.Scan(&r.Content, &r.Similarity); err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

func SeedDatabase(dbpool *pgxpool.Pool, ollamaClient *api.Client) {
	var count int
	err := dbpool.QueryRow(context.Background(), "SELECT COUNT(*) FROM sentences").Scan(&count)
	if err != nil {
		log.Printf("Could not check sentence count, skipping seed: %v", err)
		return
	}
	if count > 0 {
		fmt.Println("Database already seeded. Skipping.")
		return
	}

	fmt.Println("Seeding database with sample sentences...")
	sentences := []string{
		"The sun rises in the east.",
		"A cat is a small, furry mammal.",
		"Go is an open-source programming language.",
	}

	for _, s := range sentences {
		req := &api.EmbeddingRequest{
			Model:  "mxbai-embed-large",
			Prompt: s,
		}
		resp, err := ollamaClient.Embeddings(context.Background(), req)
		if err != nil {
			log.Printf("Could not get embedding for '%s': %v", s, err)
			continue
		}

		_, err = dbpool.Exec(context.Background(),
			"INSERT INTO sentences (content, embedding) VALUES ($1, $2)",
			s, pgvector.NewVector(toFloat32(resp.Embedding)))
		if err != nil {
			log.Printf("Could not insert sentence '%s': %v", s, err)
		}
	}
	fmt.Println("Database seeding complete!")
}
