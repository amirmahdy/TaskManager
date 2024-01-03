package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	mockdb "taskmanager/db/mock"
	db "taskmanager/db/model"
	"taskmanager/token"
	"taskmanager/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func randomTask(user string) db.Task {
	title := sql.NullString{String: utils.CreateRandomString(10), Valid: true}
	description := sql.NullString{String: utils.CreateRandomString(10), Valid: true}
	dDate, _ := time.Parse("2006-01-02T15:04:05", "2021-01-01T15:04:05")
	dueDate := sql.NullTime{Time: dDate, Valid: true}
	curDate := sql.NullTime{Time: time.Now(), Valid: true}
	priority := sql.NullInt32{Int32: int32(utils.CreateRandomInt(2) + 1), Valid: true}

	return db.Task{
		ID:          int64(utils.CreateRandomInt(10)),
		Username:    user,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Priority:    priority,
		CreatedAt:   curDate,
	}
}

func TestCreateTask(t *testing.T) {
	user, _ := randomUser(t)
	task := randomTask(user.Username)

	tasks := []struct {
		name          string
		body          createTaskRequest
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			body: createTaskRequest{
				Title:       task.Title.String,
				Description: task.Description.String,
				DueDate:     task.DueDate.Time,
				Priority:    task.Priority.Int32,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				// DB input arguments
				arg := db.CreateTaskParams{
					Username:    user.Username,
					Title:       task.Title,
					Description: task.Description,
					DueDate:     task.DueDate,
					Priority:    task.Priority,
				}
				store.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)

				store.EXPECT().CreateTask(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(task, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTask(t, recorder.Body, task)
			},
		},
	}

	for _, tc := range tasks {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock store
			store := mockdb.NewMockStore(ctrl)
			// Create test server using mock store
			server := newTestServer(t, store)
			tc.buildStubs(store)

			recorder := httptest.NewRecorder()
			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/task"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.token)
			// Serving requests
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func requireBodyMatchTask(t *testing.T, body *bytes.Buffer, task db.Task) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTask db.Task
	err = json.Unmarshal(data, &gotTask)

	require.NoError(t, err)
	require.Equal(t, task.Title, gotTask.Title)
	require.Equal(t, task.Description, gotTask.Description)
	require.Equal(t, task.DueDate, gotTask.DueDate)
	require.Equal(t, task.Priority, gotTask.Priority)
	require.Equal(t, task.CreatedAt.Time.Format("2006-01-02T15:04:05"),
		gotTask.CreatedAt.Time.Format("2006-01-02T15:04:05"))

}
