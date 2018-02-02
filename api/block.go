package api

import (
	"bytes"
	"encoding/hex"

	"github.com/eliothedeman/heath/block"

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
	var b *block.Block
	err = db.EachBlock(d, func(x *block.Block) bool {
		if bytes.Equal(x.GetHash(), hash) {
			b = x
			return false
		}
		return true
	})
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
