package mongoinit

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// InjectAsMiddleware injects the Mongo client into the Gin context
func InjectAsMiddleware(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	}
}
