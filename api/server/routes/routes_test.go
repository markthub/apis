package routes

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markthub/apis/api/api/middleware"
	"github.com/markthub/apis/api/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var router *gin.Engine
var db *mongo.Database

func TestMain(m *testing.M) {
	tearUp()
	c := m.Run()
	tearDown()
	os.Exit(c)
}

func tearUp() {
	gin.SetMode(gin.TestMode)

	router = gin.Default()
	router.RedirectTrailingSlash = true

	db = database.GetDatabaseClient("markthub", "markthub", "127.0.0.1", "3306", "testing")

	// Load fixtures

	router.Use(middleware.DBClient(db))
}

func tearDown() {
	router = nil
}

// MockRequest will send a request to the server. Used for testing purposes
func MockRequest(method, path string, body io.Reader) (int, *bytes.Buffer, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return -1, nil, err
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	return w.Code, w.Body, nil
}
