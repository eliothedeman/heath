package api_test

import (
	"crypto/ecdsa"
	"io/ioutil"
	"os"
	"testing"

	"github.com/eliothedeman/heath/keystore"
	"github.com/eliothedeman/heath/util"

	"github.com/eliothedeman/heath/api"
	"github.com/eliothedeman/heath/db"
	"github.com/gin-gonic/gin"
)

func testRouter(t *testing.T, key *ecdsa.PrivateKey) *gin.Engine {
	t.Helper()
	e := gin.Default()
	v1 := e.Group("v1", api.WithDB(newTestDB(t)), api.WithKeystore(keystore.NewMemoryStore(key)))
	api.RegisterAPI(v1)
	return e
}

func newKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()
	k, err := util.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	return k
}

func newTestDB(t *testing.T) db.Driver {
	t.Helper()
	f, err := ioutil.TempFile(os.TempDir(), t.Name())
	if err != nil {
		t.Fatal(err)
	}

	return db.NewNonCachedDriver(f)
}

func harness(t *testing.T) (*ecdsa.PrivateKey, db.Driver, *gin.Engine) {
	k := newKey(t)
	d := newTestDB(t)
	e := api.NewEngine(k, d)
	return k, d, e
}
