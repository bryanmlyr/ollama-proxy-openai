package main

import (
	"encoding/json"
	"fmt"
	"ollamaproxy/config"
	"ollamaproxy/model"
	"ollamaproxy/service"

	"github.com/gin-gonic/gin"
)

func main() {
	configLoader := config.NewConfigLoader()
	proxyService := service.NewOllamaProxyService(configLoader)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Ollama is running")
	})

	r.GET("/api/tags", func(c *gin.Context) {
		tags, err := proxyService.GetTags()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, model.OllamaTagsResponse{Models: tags})
	})

	r.POST("/api/chat", func(c *gin.Context) {
		var request model.OllamaChatRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}
		ch, err := proxyService.Chat(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "application/x-ndjson")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// Write NDJSON one per line as they arrive
		w := c.Writer
		enc := json.NewEncoder(w)
		notify := c.Request.Context().Done()
		flusher, ok := w.(gin.ResponseWriter) // gin's Writer implements http.Flusher
		if !ok {
			c.JSON(500, gin.H{"error": "Streaming not supported"})
			return
		}

		for {
			select {
			case res, ok := <-ch:
				if !ok {
					return
				}
				enc.Encode(res)
				flusher.Flush()
			case <-notify:
				return
			}
		}
	})

	fmt.Println("Listening on :11434 (gin)")
	r.Run(":11434")
}
