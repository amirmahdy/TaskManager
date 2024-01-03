package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTask(t *testing.T) {
	username, _, err := createFakeDBUser()
	require.NoError(t, err)

	// Prepare test data
	due, err := time.Parse("2006-01-02T15:04:05", "2021-01-01T00:00:00")
	require.NoError(t, err)

	params := CreateTaskParams{
		Username:    username,
		Title:       sql.NullString{String: "Test Task", Valid: true},
		Description: sql.NullString{String: "This is a test task", Valid: true},
		DueDate:     sql.NullTime{Time: due, Valid: true},
		Priority:    sql.NullInt32{Int32: 1, Valid: true},
	}

	// Call the CreateTask function
	task, err := testStore.CreateTask(context.Background(), params)
	require.NoError(t, err)

	// Assert the task properties
	require.Equal(t, params.Username, task.Username, "expected username %s, got %s", params.Username, task.Username)
	require.Equal(t, params.Title.String, task.Title.String, "expected title %s, got %s", params.Title.String, task.Title.String)
	require.Equal(t, params.Description.String, task.Description.String, "expected description %s, got %s", params.Description.String, task.Description.String)
	require.Equal(t, params.DueDate.Time.Local(), task.DueDate.Time.Local(), "expected due date %s, got %s", params.DueDate.Time, task.DueDate.Time)
	require.Equal(t, params.Priority.Int32, task.Priority.Int32, "expected priority %d, got %d", params.Priority.Int32, task.Priority.Int32)
}
