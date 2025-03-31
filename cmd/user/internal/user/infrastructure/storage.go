package infrastructure

import (
	"context"
	"fmt"
	userProto "github.com/Errera11/user/internal/protogen"
	"github.com/Errera11/user/internal/user/repository"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUserRepo(conn *pgx.Conn) repository.UserRepo {
	createUsersTable(conn)

	return &UserRepo{conn: conn}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *userProto.CreateUserRequest) (int32, error) {
	var userID int32
	query := `INSERT INTO users (username, email, password, image) 
	          VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.conn.QueryRow(ctx, query, user.Username, user.Email, user.Password, user.Image).
		Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("could not create user: %v", err)
	}

	return userID, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id int32) (*userProto.GetUserByIdResponse, error) {
	query := `SELECT id, username, email, created_at, image
			  FROM users WHERE id = $1`

	var user userProto.GetUserByIdResponse
	err := r.conn.QueryRow(ctx, query, id).Scan(&user.Id, &user.Username, &user.Email,
		&user.CreatedAt, &user.Image)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, fmt.Errorf("could not retrieve user: %v", err)
	}
	return &user, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*userProto.GetUserByIdResponse, error) {
	query := `SELECT id, username, email, created_at 
			  FROM users WHERE email = $1`

	var user userProto.GetUserByIdResponse
	err := r.conn.QueryRow(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email,
		&user.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("could not retrieve user: %v", err)
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *userProto.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, image = $4
	          WHERE id = $5`

	_, err := r.conn.Exec(ctx, query, user.Username, user.Email, user.Password, user.Image)
	if err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}

	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int32) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.conn.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}

	return nil
}

func (r *UserRepo) GetAllUsers(ctx context.Context) (*userProto.GetAllUsersResponse, error) {
	query := "SELECT id, username, email, COALESCE(image, 'null'), created_at::text FROM users"

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch users: %v", err)
	}
	defer rows.Close()

	var users []*userProto.GetUserByIdResponse

	for rows.Next() {
		var user userProto.GetUserByIdResponse
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Image, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}
		users = append(users, &user)
	}

	// Check for any errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over rows: %v", err)
	}

	return &userProto.GetAllUsersResponse{Users: users}, nil
}

func createUsersTable(conn *pgx.Conn) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password TEXT NOT NULL,
	    image TEXT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`

	createIndexQuery := `
	CREATE INDEX IF NOT EXISTS idx_email ON users(email);
	`

	tx, err := conn.Begin(context.Background())
	if err != nil {
		fmt.Errorf("could not begin transaction: %v", err)
		return err
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), createTableQuery)
	if err != nil {
		fmt.Errorf("could not create users table: %v", err)
		return err
	}

	_, err = tx.Exec(context.Background(), createIndexQuery)
	if err != nil {
		fmt.Errorf("could not create index on email: %v", err)
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		fmt.Errorf("could not commit transaction: %v", err)
		return err
	}

	return nil
}
