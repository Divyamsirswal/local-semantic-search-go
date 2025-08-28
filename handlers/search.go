package handlers

import (
	"net/http"
	"semantic-search-api/database"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/ollama/ollama/api"
)

func MakeSearchHandler(dbpool *pgxpool.Pool, ollamaClient *api.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParam("q")
		if query == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'q' is required"})
		}

		req := &api.EmbeddingRequest{
			Model:  "mxbai-embed-large",
			Prompt: query,
		}
		ctx := c.Request().Context()
		resp, err := ollamaClient.Embeddings(ctx, req)
		if err != nil {
			c.Logger().Errorf("Failed to get query embedding: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get query embedding"})
		}

		// This now correctly passes the []float64 slice to the updated function.
		results, err := database.SearchSentences(dbpool, resp.Embedding)
		if err != nil {
			c.Logger().Errorf("Database query failed: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database query failed"})
		}

		return c.JSON(http.StatusOK, results)
	}
}
