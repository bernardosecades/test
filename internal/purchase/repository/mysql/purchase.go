package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bernardosecades/test/internal/purchase/entity"
	_ "github.com/go-sql-driver/mysql"
)

const formatDate = "2006-01-02 15:04:05"

type MySQLPurchaseRepository struct {
	SQL *sql.DB
}

func NewMySQLPurchaseRepository(dbName string, dbUser string, dbPass string, dbHost string, dbPort string) MySQLPurchaseRepository {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
	}

	err = d.Ping() // Need to do this to check that the connection is valid
	if err != nil {
		log.Fatal(err)
	}

	return MySQLPurchaseRepository{SQL: d}
}

func (r *MySQLPurchaseRepository) GetProduct(ctx context.Context, ID string) (*entity.Product, error) {
	res := r.SQL.QueryRowContext(ctx, "SELECT *  FROM product WHERE id = ? FOR UPDATE", ID)

	var product entity.Product
	err := res.Scan(&product.ID, &product.Available, &product.CreatedAt, &product.UpdatedAt)

	if err != nil &&  err != sql.ErrNoRows {
		return nil, err
	}

	if len(product.ID) == 0 {
		return nil, nil
	}

	return &product, nil
}

func (r *MySQLPurchaseRepository) SetToNoAvailable(ctx context.Context, ID string) error {
	_, err := r.SQL.ExecContext(
		ctx,
		"UPDATE product SET available = false, updated_at = ? WHERE id = ?",
		time.Now().UTC().Format(formatDate),
		ID,
	)

	if err != nil {
		return err
	}

	return nil
}
