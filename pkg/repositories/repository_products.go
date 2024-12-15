package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"reflect"
	"strings"
)

type ProductsSqlRepository struct {
	db *sqlx.DB
}

func newProductsSqlRepository(sqlDb *sqlx.DB) *ProductsSqlRepository {
	return &ProductsSqlRepository{db: sqlDb}
}

func (r *ProductsSqlRepository) Create(m *models.Product) error {
	if m.Description == nil {
		m.Description = new(string)
		*m.Description = ""
	}
	query := fmt.Sprintf("INSERT INTO products (name, description) values ($1, $2) RETURNING id, created_at, updated_at")
	row := r.db.QueryRow(query, m.Name, m.Description)

	return row.Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
}
func (r *ProductsSqlRepository) CreateWithImage(m *models.Product, fm *models.File, filesRepo Files) error {
	if m.Description == nil {
		m.Description = new(string)
		*m.Description = ""
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("INSERT INTO products (name, description) values ($1, $2) RETURNING id, created_at, updated_at")
	row := tx.QueryRow(query, m.Name, m.Description)
	if err = row.Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt); err != nil {
		_ = tx.Rollback()
		return err
	}

	fp := &models.FilePartial{}
	modelName := "Product"
	fp.ModelName = &modelName
	fp.ModelId = m.ID
	fp.Data = new(map[string]string)
	*fp.Data = map[string]string{"type": "image"}

	fmt.Println("tx->", reflect.TypeOf(tx))
	/*if err = filesRepo.Update(fm, fp, tx); err != nil {
		_ = tx.Rollback()
		return err
	}*/

	m.Image = &models.File{ID: fm.ID, Ext: fm.Ext, UUID: fm.UUID, Data: fp.Data}

	return tx.Commit()

}

func (r *ProductsSqlRepository) Update(m *models.Product, data *models.ProductPartial) error {
	fields := []string{}
	argIndex := 1
	values := []any{}

	if data.Name != nil {
		fields = append(fields, fmt.Sprintf("name = $%d", argIndex))
		argIndex++
		values = append(values, *data.Name)
	}
	if data.Description != nil {
		fields = append(fields, fmt.Sprintf("description = $%d", argIndex))
		argIndex++
		values = append(values, *data.Description)
	}

	values = append(values, m.ID)

	query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d RETURNING name, description, updated_at", strings.Join(fields, ", "), argIndex)
	row := r.db.QueryRow(query, values...)
	return row.Scan(&m.Name, &m.Description, &m.UpdatedAt)
}

func (r *ProductsSqlRepository) Delete(m *models.Product) error {
	query := fmt.Sprintf("DELETE FROM products WHERE id = $1")
	_, err := r.db.Exec(query, m.ID)
	return err
}

func (r *ProductsSqlRepository) Find(ID string) (*models.Product, error) {
	m := &models.Product{}
	fm := &models.File{}
	var fDataString *string

	query := fmt.Sprintf("SELECT p.id, p.name, p.description, p.created_at, p.updated_at, f.id, f.ext, f.uuid, f.data FROM products p LEFT JOIN files f ON p.id = f.model_id WHERE p.id = $1")
	row := r.db.QueryRow(query, ID)

	if err := row.Scan(&m.ID, &m.Name, &m.Description, &m.CreatedAt, &m.UpdatedAt, &fm.ID, &fm.Ext, &fm.UUID, &fDataString); err != nil {
		return nil, errRecordNotFound
	}

	if fm.ID != nil {
		if fDataString != nil {
			if err := json.Unmarshal([]byte(*fDataString), &fm.Data); err != nil {
				return nil, err
			}
		}
		m.Image = fm
	}

	return m, nil
}

/*func (r *ProductsSqlRepository) FindProduct(ID string) (*models.Product, error) {
	m := &models.Product{}
	query := fmt.Sprintf("SELECT p.id, p.name, p.description, p.created_at, p.updated_at FROM products p WHERE p.id = $1")
	row := r.db.QueryRow(query, ID)
	if err := row.Scan(&m.ID, &m.Name, &m.Description, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return nil, errRecordNotFound
	}
	return m, nil
}*/

func (r *ProductsSqlRepository) List(p *PaginationRequest, s *Sort, f *models.ProductFilter) ([]*models.Product, error) {
	list := []*models.Product{}

	where, values, argIndex := r.where(f)

	values = append(values, p.PageSize)
	values = append(values, p.PageSize*p.PageNumber)

	orderBy := fmt.Sprintf("ORDER BY p.%s %s", s.SortField, s.SortDirection)

	limit := fmt.Sprintf("LIMIT $%d OFFSET $%d", argIndex, argIndex+1)

	query := fmt.Sprintf(
		"SELECT p.id, p.name, p.description, p.created_at, p.updated_at, f.id, f.ext, f.uuid, f.data FROM products p LEFT JOIN files f ON p.id = f.model_id %s %s %s",
		where, orderBy, limit)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		m := &models.Product{}
		fm := &models.File{}
		var fDataString *string

		err := rows.Scan(&m.ID, &m.Name, &m.Description, &m.CreatedAt, &m.UpdatedAt, &fm.ID, &fm.Ext, &fm.UUID, &fDataString)
		if err != nil {
			return nil, err
			//continue
		}

		if fm.ID != nil {
			if fDataString != nil {
				if err := json.Unmarshal([]byte(*fDataString), &fm.Data); err != nil {
					return nil, err
				}
			}
			m.Image = fm
		}

		list = append(list, m)
	}

	return list, nil
}

func (r *ProductsSqlRepository) where(f *models.ProductFilter) (string, []any, int) {
	fields := []string{}
	argIndex := 1
	values := []any{}

	if f.Name != nil {
		fields = append(fields, fmt.Sprintf("name LIKE $%d", argIndex))
		values = append(values, "%"+*f.Name+"%")
		argIndex++
	}
	if f.Description != nil {
		fields = append(fields, fmt.Sprintf("description LIKE $%d", argIndex))
		values = append(values, "%"+*f.Description+"%")
		argIndex++
	}

	where := ""
	if len(fields) > 0 {
		where = fmt.Sprintf("WHERE %s", strings.Join(fields, " AND "))
	}
	return where, values, argIndex
}

func (r *ProductsSqlRepository) Count(f *models.ProductFilter) (*int, error) {
	where, values, _ := r.where(f)
	query := fmt.Sprintf("SELECT COUNT(*) FROM products %s", where)
	row := r.db.QueryRow(query, values...)

	var count int
	if err := row.Scan(&count); err != nil {
		return nil, err
	}

	return &count, nil
}
