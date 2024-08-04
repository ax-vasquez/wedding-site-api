//go:build unit
// +build unit

package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/ax-vasquez/wedding-site-api/models"
	"github.com/stretchr/testify/assert"
)

func TestEntreeControllerUnit(t *testing.T) {
	os.Setenv("USE_MOCK_DB", "true")
	_, mock, _ := models.Setup()
	assert := assert.New(t)
	router := paveRoutes()
	t.Run("GET /api/v1/entrees - internal server error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "entrees" WHERE "entrees"."deleted_at" IS NULL`)).WillReturnError(fmt.Errorf("arbitrary database error"))

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/entrees", nil)
		router.ServeHTTP(w, req)
		assert.Equal(nil, err)
		assert.Equal(http.StatusInternalServerError, w.Code)
	})
	t.Run("GET /api/v1/user/:id/entrees - internal server error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT "entrees"."created_at","entrees"."updated_at","entrees"."deleted_at","entrees"."id","entrees"."option_name" FROM "entrees" JOIN users ON entrees.id = users.entree_selection_id AND users.id = $1 WHERE "entrees"."deleted_at" IS NULL`)).WithArgs(models.NilUuid).WillReturnError(fmt.Errorf("arbitrary database error"))

		route := fmt.Sprintf("/api/v1/user/%s/entrees", models.NilUuid)
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", route, nil)
		router.ServeHTTP(w, req)
		assert.Equal(nil, err)
		assert.Equal(http.StatusInternalServerError, w.Code)
	})
	// t.Run("POST /api/v1/entrees - internal server error", func(t *testing.T) {

	// })
	// t.Run("DELETE /api/v1/entrees - internal server error", func(t *testing.T) {

	// })
}
