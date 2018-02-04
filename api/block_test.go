package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/block/test_util"
)

func TestGetBlock(t *testing.T) {
	t.Run("Not exists", func(t *testing.T) {
		_, _, e := harness(t)
		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/api/block/bloasdlfkjlaksjdflkj", nil)
		e.ServeHTTP(w, r)

		if w.Code != 400 {
			t.Errorf("Wanted 400 got %d", w.Code)
		}
	})

	t.Run("Only block", func(t *testing.T) {
		k, d, e := harness(t)
		b := block.NewBlock(nil, []*block.Transaction{test_util.GenTestTransaction(t, k)})
		err := d.Write(b)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", fmt.Sprintf("/api/block/%x", b.GetHash()), nil)
		e.ServeHTTP(w, r)

		if w.Code != 200 {
			t.Errorf("Wanted 200 got %d", w.Code)
		}

		b2 := new(block.Block)
		err = json.Unmarshal(w.Body.Bytes(), b2)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(b, b2) {
			t.Fatal(w.Body.String())
		}

	})

}
