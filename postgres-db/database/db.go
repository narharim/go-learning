// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAuthorStmt, err = db.PrepareContext(ctx, createAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAuthor: %w", err)
	}
	if q.createBookStmt, err = db.PrepareContext(ctx, createBook); err != nil {
		return nil, fmt.Errorf("error preparing query CreateBook: %w", err)
	}
	if q.deleteAuthorStmt, err = db.PrepareContext(ctx, deleteAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteAuthor: %w", err)
	}
	if q.deleteBookStmt, err = db.PrepareContext(ctx, deleteBook); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteBook: %w", err)
	}
	if q.getAuthorStmt, err = db.PrepareContext(ctx, getAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query GetAuthor: %w", err)
	}
	if q.getBookStmt, err = db.PrepareContext(ctx, getBook); err != nil {
		return nil, fmt.Errorf("error preparing query GetBook: %w", err)
	}
	if q.listAuthorsStmt, err = db.PrepareContext(ctx, listAuthors); err != nil {
		return nil, fmt.Errorf("error preparing query ListAuthors: %w", err)
	}
	if q.listBooksByAuthorStmt, err = db.PrepareContext(ctx, listBooksByAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query ListBooksByAuthor: %w", err)
	}
	if q.updateAuthorStmt, err = db.PrepareContext(ctx, updateAuthor); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateAuthor: %w", err)
	}
	if q.updateBookStmt, err = db.PrepareContext(ctx, updateBook); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateBook: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAuthorStmt != nil {
		if cerr := q.createAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAuthorStmt: %w", cerr)
		}
	}
	if q.createBookStmt != nil {
		if cerr := q.createBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createBookStmt: %w", cerr)
		}
	}
	if q.deleteAuthorStmt != nil {
		if cerr := q.deleteAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteAuthorStmt: %w", cerr)
		}
	}
	if q.deleteBookStmt != nil {
		if cerr := q.deleteBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteBookStmt: %w", cerr)
		}
	}
	if q.getAuthorStmt != nil {
		if cerr := q.getAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAuthorStmt: %w", cerr)
		}
	}
	if q.getBookStmt != nil {
		if cerr := q.getBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getBookStmt: %w", cerr)
		}
	}
	if q.listAuthorsStmt != nil {
		if cerr := q.listAuthorsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAuthorsStmt: %w", cerr)
		}
	}
	if q.listBooksByAuthorStmt != nil {
		if cerr := q.listBooksByAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listBooksByAuthorStmt: %w", cerr)
		}
	}
	if q.updateAuthorStmt != nil {
		if cerr := q.updateAuthorStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateAuthorStmt: %w", cerr)
		}
	}
	if q.updateBookStmt != nil {
		if cerr := q.updateBookStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateBookStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	createAuthorStmt      *sql.Stmt
	createBookStmt        *sql.Stmt
	deleteAuthorStmt      *sql.Stmt
	deleteBookStmt        *sql.Stmt
	getAuthorStmt         *sql.Stmt
	getBookStmt           *sql.Stmt
	listAuthorsStmt       *sql.Stmt
	listBooksByAuthorStmt *sql.Stmt
	updateAuthorStmt      *sql.Stmt
	updateBookStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		createAuthorStmt:      q.createAuthorStmt,
		createBookStmt:        q.createBookStmt,
		deleteAuthorStmt:      q.deleteAuthorStmt,
		deleteBookStmt:        q.deleteBookStmt,
		getAuthorStmt:         q.getAuthorStmt,
		getBookStmt:           q.getBookStmt,
		listAuthorsStmt:       q.listAuthorsStmt,
		listBooksByAuthorStmt: q.listBooksByAuthorStmt,
		updateAuthorStmt:      q.updateAuthorStmt,
		updateBookStmt:        q.updateBookStmt,
	}
}
