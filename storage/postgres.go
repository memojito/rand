package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/memojito/igapi/types"
	"github.com/memojito/igapi/utils"
)

type Storage interface {
	Store(context.Context, *types.CreateUserRequest) error
	List(context.Context, []int) ([]types.User, error)
	Delete(context.Context, int) error
}

type PostgresStorage struct {
	Conn *pgx.Conn
}

func New(url string) (*PostgresStorage, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Panic(err)
	}
	return &PostgresStorage{Conn: conn}, nil
}

func (ps *PostgresStorage) Store(ctx context.Context, u *types.CreateUserRequest) error {
	q := "INSERT INTO iguser (name, email) VALUES ($1, $2)"
	log.Printf("query: %s", q)

	log.Printf("name: %s, email: %s", u.Name, u.Email)

	_, err := ps.Conn.Exec(ctx, q, u.Name, u.Email)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStorage) List(ctx context.Context, ids []int) ([]types.User, error) {
	placeholders := utils.PrepareWhereINString(len(ids))

	q := "SELECT * FROM iguser WHERE id IN " + placeholders + ")"
	log.Printf("query: %s", q)

	rows, err := ps.Conn.Query(ctx, q, ids)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[types.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ps *PostgresStorage) Delete(ctx context.Context, id int) error {
	return nil
}
