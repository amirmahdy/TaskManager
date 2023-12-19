-- name: GetTasks :many
SELECT * FROM tasks
WHERE username = $1 LIMIT 1;

-- name: CreateTask :one
INSERT INTO tasks (
    username, title, description, due_date, priority
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;