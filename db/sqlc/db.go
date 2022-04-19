// Code generated by sqlc. DO NOT EDIT.

package db

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
	if q.createHatcheryStmt, err = db.PrepareContext(ctx, createHatchery); err != nil {
		return nil, fmt.Errorf("error preparing query CreateHatchery: %w", err)
	}
	if q.createKuroilerStmt, err = db.PrepareContext(ctx, createKuroiler); err != nil {
		return nil, fmt.Errorf("error preparing query CreateKuroiler: %w", err)
	}
	if q.createProductionStmt, err = db.PrepareContext(ctx, createProduction); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduction: %w", err)
	}
	if q.deleteKuroilerStmt, err = db.PrepareContext(ctx, deleteKuroiler); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteKuroiler: %w", err)
	}
	if q.getHatcheryStmt, err = db.PrepareContext(ctx, getHatchery); err != nil {
		return nil, fmt.Errorf("error preparing query GetHatchery: %w", err)
	}
	if q.getKuroilerStmt, err = db.PrepareContext(ctx, getKuroiler); err != nil {
		return nil, fmt.Errorf("error preparing query GetKuroiler: %w", err)
	}
	if q.getProductionStmt, err = db.PrepareContext(ctx, getProduction); err != nil {
		return nil, fmt.Errorf("error preparing query GetProduction: %w", err)
	}
	if q.listHatcheryStmt, err = db.PrepareContext(ctx, listHatchery); err != nil {
		return nil, fmt.Errorf("error preparing query ListHatchery: %w", err)
	}
	if q.listKuroilerStmt, err = db.PrepareContext(ctx, listKuroiler); err != nil {
		return nil, fmt.Errorf("error preparing query ListKuroiler: %w", err)
	}
	if q.listProductionStmt, err = db.PrepareContext(ctx, listProduction); err != nil {
		return nil, fmt.Errorf("error preparing query ListProduction: %w", err)
	}
	if q.updateKuroilerStmt, err = db.PrepareContext(ctx, updateKuroiler); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateKuroiler: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createHatcheryStmt != nil {
		if cerr := q.createHatcheryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createHatcheryStmt: %w", cerr)
		}
	}
	if q.createKuroilerStmt != nil {
		if cerr := q.createKuroilerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createKuroilerStmt: %w", cerr)
		}
	}
	if q.createProductionStmt != nil {
		if cerr := q.createProductionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductionStmt: %w", cerr)
		}
	}
	if q.deleteKuroilerStmt != nil {
		if cerr := q.deleteKuroilerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteKuroilerStmt: %w", cerr)
		}
	}
	if q.getHatcheryStmt != nil {
		if cerr := q.getHatcheryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getHatcheryStmt: %w", cerr)
		}
	}
	if q.getKuroilerStmt != nil {
		if cerr := q.getKuroilerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getKuroilerStmt: %w", cerr)
		}
	}
	if q.getProductionStmt != nil {
		if cerr := q.getProductionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getProductionStmt: %w", cerr)
		}
	}
	if q.listHatcheryStmt != nil {
		if cerr := q.listHatcheryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listHatcheryStmt: %w", cerr)
		}
	}
	if q.listKuroilerStmt != nil {
		if cerr := q.listKuroilerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listKuroilerStmt: %w", cerr)
		}
	}
	if q.listProductionStmt != nil {
		if cerr := q.listProductionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listProductionStmt: %w", cerr)
		}
	}
	if q.updateKuroilerStmt != nil {
		if cerr := q.updateKuroilerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateKuroilerStmt: %w", cerr)
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
	db                   DBTX
	tx                   *sql.Tx
	createHatcheryStmt   *sql.Stmt
	createKuroilerStmt   *sql.Stmt
	createProductionStmt *sql.Stmt
	deleteKuroilerStmt   *sql.Stmt
	getHatcheryStmt      *sql.Stmt
	getKuroilerStmt      *sql.Stmt
	getProductionStmt    *sql.Stmt
	listHatcheryStmt     *sql.Stmt
	listKuroilerStmt     *sql.Stmt
	listProductionStmt   *sql.Stmt
	updateKuroilerStmt   *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                   tx,
		tx:                   tx,
		createHatcheryStmt:   q.createHatcheryStmt,
		createKuroilerStmt:   q.createKuroilerStmt,
		createProductionStmt: q.createProductionStmt,
		deleteKuroilerStmt:   q.deleteKuroilerStmt,
		getHatcheryStmt:      q.getHatcheryStmt,
		getKuroilerStmt:      q.getKuroilerStmt,
		getProductionStmt:    q.getProductionStmt,
		listHatcheryStmt:     q.listHatcheryStmt,
		listKuroilerStmt:     q.listKuroilerStmt,
		listProductionStmt:   q.listProductionStmt,
		updateKuroilerStmt:   q.updateKuroilerStmt,
	}
}
