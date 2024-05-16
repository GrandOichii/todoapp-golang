package endpoints_test

import (
	"testing"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/user"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldRegister(t *testing.T) {
	r := setupRouter()
	w, _ := req(r, t, "POST", "/api/v1/auth/register", dto.PostUser{
		Username: "user1",
		Password: "password",
	}, "")

	assert.Equal(t, 201, w.Code)
}

func Test_ShouldNotRegisterBadRequest(t *testing.T) {
	r := setupRouter()

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

			w, _ := req(r, t, "POST", "/api/v1/auth/register", tC.user, "")

			assert.Equal(t, 400, w.Code)
		})
	}
}

func Test_ShouldNotRegisterUsernameExists(t *testing.T) {
	r := setupRouter()
	data := dto.PostUser{
		Username: "user1",
		Password: "password",
	}
	req(r, t, "POST", "/api/v1/auth/register", data, "")
	w, _ := req(r, t, "POST", "/api/v1/auth/register", data, "")

	assert.Equal(t, 400, w.Code)
}

func Test_ShouldLogin(t *testing.T) {
	r := setupRouter()

	data := dto.PostUser{
		Username: "user1",
		Password: "password",
	}
	req(r, t, "POST", "/api/v1/auth/register", data, "")
	w, _ := req(r, t, "POST", "/api/v1/auth/login", data, "")

	assert.Equal(t, 200, w.Code)
}

func Test_ShouldNotLoginWrongUsername(t *testing.T) {
	r := setupRouter()

	data := dto.PostUser{
		Username: "user1",
		Password: "password",
	}
	w, _ := req(r, t, "POST", "/api/v1/auth/login", data, "")

	assert.Equal(t, 401, w.Code)
}

func Test_ShouldNotLoginWrongPassword(t *testing.T) {
	r := setupRouter()

	data1 := dto.PostUser{
		Username: "user1",
		Password: "password",
	}

	data2 := dto.PostUser{
		Username: "user1",
		Password: "password1",
	}
	req(r, t, "POST", "/api/v1/auth/register", data1, "")
	w, _ := req(r, t, "POST", "/api/v1/auth/login", data2, "")

	assert.Equal(t, 401, w.Code)
}
