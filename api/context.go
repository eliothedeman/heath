package api

import (
	"crypto/ecdsa"

	"github.com/eliothedeman/heath/db"
	"github.com/eliothedeman/heath/keystore"
	"github.com/gin-gonic/gin"
)

// RegisterAPI sets up all the api methods on this router
func RegisterAPI(r gin.IRouter) {
	r.GET("block/:hash", getBlock)
}

// NewEngine register all context objects
func NewEngine(k *ecdsa.PrivateKey, d db.Driver) *gin.Engine {
	e := gin.Default()
	a := e.Group("api", WithDB(d), WithKeystore(keystore.NewMemoryStore(k)))
	RegisterAPI(a)

	return e
}

// WithDB returns a handler that sets the database
func WithDB(d db.Driver) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", d)
	}
}

// WithKeystore sets a keystore on the context
func WithKeystore(k keystore.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("keystore", k)
	}
}
