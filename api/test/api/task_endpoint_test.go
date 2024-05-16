package endpoints_test

import (
	"encoding/json"
	"testing"

	dto "github.com/GrandOichii/todoapp-golang/api/dto/task"
	"github.com/go-playground/assert/v2"
)

func Test_ShouldFetchAll(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	w, _ := req(r, t, "GET", "/api/v1/task", nil, token)
	assert.Equal(t, 200, w.Code)
}

func Test_ShouldCreate(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	w, _ := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	assert.Equal(t, 201, w.Code)
}

func Test_ShouldNotCreateNotLoggedIn(t *testing.T) {
	r := setupRouter()
	w, _ := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, "")
	assert.Equal(t, 401, w.Code)
}

func Test_ShouldNotCreate(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")

	testCases := []struct {
		desc string
		task dto.CreateTask
	}{
		{
			desc: "Title empty",
			task: dto.CreateTask{
				Title: "",
				Text:  "task description",
			},
		},
		{
			desc: "Title too short",
			task: dto.CreateTask{
				Title: "ti",
				Text:  "task description",
			},
		},
		{
			desc: "Title too long",
			task: dto.CreateTask{
				Title: "titletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitle",
				Text:  "task description",
			},
		},
		{
			desc: "Description too long",
			task: dto.CreateTask{
				Title: "task title",
				Text:  "task descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask descriptiontask description",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w, _ := req(r, t, "POST", "/api/v1/task", tC.task, token)
			assert.Equal(t, 400, w.Code)
		})
	}
}

func Test_ShouldFetchById(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	_, data := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	var task dto.GetTask
	err := json.Unmarshal(data, &task)
	if err != nil {
		panic(err)
	}

	w, _ := req(r, t, "GET", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 200, w.Code)
}

func Test_ShouldNotFetchById(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")

	w, _ := req(r, t, "GET", "/api/v1/task/invalidid", nil, token)

	assert.Equal(t, 404, w.Code)
}

func Test_ShouldNotFetchByStolenId(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	_, data := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	var task dto.GetTask
	err := json.Unmarshal(data, &task)
	if err != nil {
		panic(err)
	}

	token = loginAs(r, t, "otheruser", "password")
	w, _ := req(r, t, "GET", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 404, w.Code)
}

func Test_ShouldToggleCompleted(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	_, data := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	var task dto.GetTask
	err := json.Unmarshal(data, &task)
	if err != nil {
		panic(err)
	}

	w, _ := req(r, t, "PATCH", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 200, w.Code)
}

func Test_ShouldNotToggleCompleted(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")

	w, _ := req(r, t, "PATCH", "/api/v1/task/invalidid", nil, token)

	assert.Equal(t, 404, w.Code)
}

func Test_ShouldNotToggleCompletedStolenId(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	_, data := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	var task dto.GetTask
	err := json.Unmarshal(data, &task)
	if err != nil {
		panic(err)
	}

	token = loginAs(r, t, "otheruser", "password")
	w, _ := req(r, t, "PATCH", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 404, w.Code)
}

func Test_ShouldDelete(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	_, data := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	var task dto.GetTask
	err := json.Unmarshal(data, &task)
	if err != nil {
		panic(err)
	}

	w, _ := req(r, t, "DELETE", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 200, w.Code)

	w, _ = req(r, t, "GET", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 404, w.Code)
}

func Test_ShouldNotDelete(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")

	w, _ := req(r, t, "DELETE", "/api/v1/task/invalidid", nil, token)

	assert.Equal(t, 404, w.Code)
}

func Test_ShouldNotDeleteStolenId(t *testing.T) {
	r := setupRouter()
	token := loginAs(r, t, "user", "password")
	_, data := req(r, t, "POST", "/api/v1/task", dto.CreateTask{
		Title: "task title",
		Text:  "task description",
	}, token)
	var task dto.GetTask
	err := json.Unmarshal(data, &task)
	if err != nil {
		panic(err)
	}

	token = loginAs(r, t, "otheruser", "password")
	w, _ := req(r, t, "DELETE", "/api/v1/task/"+task.Id, nil, token)

	assert.Equal(t, 404, w.Code)
}
