package postgres

import (
	"app/internal/errors"
	"app/tenants/internal/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CompanyRepository struct {
	tableName string
	db        *pgx.Conn
}

var _ domain.CompanyRepository = (*CompanyRepository)(nil)

func NewCompanyRepository(tableName string, db *pgx.Conn) CompanyRepository {
	return CompanyRepository{
		tableName: tableName,
		db:        db,
	}
}

// AddTenant implements domain.CompanyRepository.
func (r CompanyRepository) AddTenant(ctx context.Context, tenantID string, name string) error {
	const query = "INSERT INTO %s (id, name) VALUES ($1, $2)"

	_, err := r.db.Exec(ctx, r.table(query), tenantID, name)

	return err
}

func (r CompanyRepository) Find(ctx context.Context, tenantID string) (*domain.CompanyTenant, error) {
	const query = `SELECT name FROM %s WHERE id = $1 LIMIT 1`

	store := &domain.CompanyTenant{
		ID: tenantID,
	}

	err := r.db.QueryRow(ctx, r.table(query), tenantID).Scan(&store.Name)
	if err != nil {
		return nil, errors.Wrap(err, "scanning tenant")
	}

	return store, nil
}

func (r CompanyRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
