package endpoints_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GrandOichii/todoapp-golang/api/config"
	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/GrandOichii/todoapp-golang/api/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

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

func Test_ShouldRegister(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()

	userData := dto.PostUser{
		Username: "user1",
		Password: "password",
	}
	data, _ := json.Marshal(userData)
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(string(data)))

	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func Test_ShouldNotRegisterBadRequest(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()

	testCases := []struct {
		desc string
		user dto.PostUser
	}{
		{
			desc: "No username",
			user: dto.PostUser{
				Username: "",
				Password: "password",
			},
		},
		{
			desc: "Short username",
			user: dto.PostUser{
				Username: "u",
				Password: "password",
			},
		},
		{
			desc: "Long username",
			user: dto.PostUser{
				Username: "usernameusernameusernameusernameusernameusernameusernameusernameusernameusernameusernameusernameusernameusernameusernameusername",
				Password: "password",
			},
		},
		{
			desc: "No password",
			user: dto.PostUser{
				Username: "user",
				Password: "",
			},
		},
		{
			desc: "Short password",
			user: dto.PostUser{
				Username: "user",
				Password: "passwor",
			},
		},
		{
			desc: "Long password",
			user: dto.PostUser{
				Username: "user",
				Password: "passwordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpasswordpassword",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, _ := json.Marshal(tC.user)
			req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(string(data)))

			r.ServeHTTP(w, req)

			assert.Equal(t, 400, w.Code)
		})
	}
}

func Test_ShouldNotRegisterUsernameExists(t *testing.T) {
	r := setupRouter()
	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	userData := dto.PostUser{
		Username: "user1",
		Password: "password",
	}
	data, _ := json.Marshal(userData)
	req1, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(string(data)))
	r.ServeHTTP(w1, req1)
	req2, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(string(data)))
	r.ServeHTTP(w2, req2)

	assert.Equal(t, 400, w2.Code)
}
