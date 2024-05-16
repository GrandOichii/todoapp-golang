package endpoints_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GrandOichii/todoapp-golang/api/config"
	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/GrandOichii/todoapp-golang/api/router"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	container, err := mongodb.RunContainer(context.Background(), testcontainers.WithImage("mongo"))
	if err != nil {
		panic(err)
	}
	conn, err := container.ConnectionString(context.Background())
	if err != nil {
		panic(err)
	}

	// !FIXME can't read from file for some reason
	config := config.Configuration{
		Port: "8080",
		Db: config.DbConfiguration{
			ConnectionUri:  conn,
			DbName:         "test_todoapp",
			TaskCollection: config.CollectionConfiguration{Name: "tasks"},
			UserCollection: config.CollectionConfiguration{Name: "users"},
		},
	}

	router := router.CreateRouter(&config)

	return router
}

func toData(o interface{}) io.Reader {
	j, _ := json.Marshal(o)
	return bytes.NewBuffer(j)
}

func req(r *gin.Engine, t *testing.T, request string, path string, data interface{}, token string) (*httptest.ResponseRecorder, []byte) {
	var reqData io.Reader = nil
	if data != nil {
		reqData = toData(data)
	}
	req, err := http.NewRequest(request, path, reqData)
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	checkErr(t, err)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	result, err := io.ReadAll(w.Body)
	checkErr(t, err)
	return w, result
}

func createUser(r *gin.Engine, t *testing.T, username string, password string) {
	req(r, t, "POST", "/api/v1/auth/register", dto.PostUser{
		Password: password,
		Username: username,
	}, "")
}

func loginAs(r *gin.Engine, t *testing.T, username string, password string) string {
	createUser(r, t, username, password)

	_, data := req(r, t, "POST", "/api/v1/auth/login", dto.PostUser{
		Username: username,
		Password: password,
	}, "")

	var res struct {
		Token string `json:"token"`
	}
	json.Unmarshal(data, &res)

	return res.Token
}
