-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, email, created_at, updated_at;

-- name: GetUserById :one
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUsers :many
SELECT id, username, email, created_at, updated_at
FROM users
ORDER BY id DESC;

-- name: CreateBlog :one
INSERT INTO blogs (user_id, title, content, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, title, content, created_at, updated_at;

-- name: GetBlogById :one
SELECT id, user_id, title, content, created_at, updated_at
FROM blogs
WHERE id = $1;

-- name: GetBlogs :many
SELECT id, user_id, title, content, created_at, updated_at
FROM blogs
ORDER BY id DESC;