package api

import (
	"encoding/hex"

	"github.com/eliothedeman/heath/db"
	"github.com/gin-gonic/gin"
)

func getBlock(c *gin.Context) {
	hash, err := hex.DecodeString(c.Params.ByName("hash"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	d := c.MustGet("db").(db.Driver)
	b, err := d.GetBlockByContentHash(hash)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	if b == nil {
		c.JSON(400, gin.H{
			"error": "not found",
		})
		return
	}
	c.JSON(200, b)
}
