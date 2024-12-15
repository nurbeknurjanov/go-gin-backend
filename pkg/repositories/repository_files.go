package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"strings"
)

type FilesSqlRepository struct {
	db *sqlx.DB
}

func newFilesSqlRepository(sqlDb *sqlx.DB) *FilesSqlRepository {
	return &FilesSqlRepository{db: sqlDb}
}

func (r *FilesSqlRepository) Create(m *models.File) error {
	var Data interface{} = nil
	if m.Data != nil && len(*m.Data) > 0 {
		j, err := json.Marshal(m.Data)
		if err != nil {
			return err
		}
		Data = string(j)
	}

	query := "INSERT INTO files (original_file_name, ext, uuid, data, model_name, model_id) values ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	row := r.db.QueryRow(query, m.OriginalFileName, m.Ext, m.UUID, Data, m.ModelName, m.ModelId)
	return row.Scan(&m.ID, &m.CreatedAt, &m.UpdatedAt)
}

func (r *FilesSqlRepository) Update(m *models.File, data *models.FilePartial, tx *sqlx.Tx) error {
	fields := []string{}
	argIndex := 1
	values := []any{}

	if data.ModelName != nil {
		fields = append(fields, fmt.Sprintf("model_name = $%d", argIndex))
		argIndex++
		values = append(values, *data.ModelName)
	}
	if data.ModelId != nil {
		fields = append(fields, fmt.Sprintf("model_id = $%d", argIndex))
		argIndex++
		values = append(values, *data.ModelId)
	}
	if data.Data != nil {
		fields = append(fields, fmt.Sprintf("data = $%d", argIndex))
		argIndex++
		data, err := json.Marshal(*data.Data)
		if err != nil {
			return err
		}
		values = append(values, string(data))
	}

	values = append(values, *m.ID)

	query := fmt.Sprintf("UPDATE files SET %s WHERE id = $%d RETURNING model_name, model_id, updated_at", strings.Join(fields, ", "), argIndex)
	row := tx.QueryRow(query, values...)
	return row.Scan(&m.ModelName, &m.ModelId, &m.UpdatedAt)
}

func (r *FilesSqlRepository) List(p *PaginationRequest, s *Sort, f *models.FileFilter) ([]*models.File, error) {
	list := []*models.File{}

	where, values, argIndex := r.where(f)

	values = append(values, p.PageSize)
	values = append(values, p.PageSize*p.PageNumber)

	orderBy := fmt.Sprintf("ORDER BY %s %s", s.SortField, s.SortDirection)

	limit := fmt.Sprintf("LIMIT $%d OFFSET $%d", argIndex, argIndex+1)

	query := fmt.Sprintf(
		"SELECT id, original_file_name, ext, uuid, data, model_name, model_id, created_at, updated_at FROM files %s %s %s",
		where, orderBy, limit)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var DataString *string
		m := &models.File{}

		if err := rows.Scan(&m.ID, &m.OriginalFileName, &m.Ext, &m.UUID, &DataString, &m.ModelName, &m.ModelId, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
			//continue
		}

		if DataString != nil {
			if err := json.Unmarshal([]byte(*DataString), &m.Data); err != nil {
				return nil, err
			}
		}

		list = append(list, m)
	}

	return list, nil
}

func (r *FilesSqlRepository) where(f *models.FileFilter) (string, []any, int) {
	fields := []string{}
	argIndex := 1
	values := []any{}

	if f.ModelName != nil {
		fields = append(fields, fmt.Sprintf("model_name = $%d", argIndex))
		values = append(values, *f.ModelName)
		argIndex++
	}
	if f.ModelId != nil {
		fields = append(fields, fmt.Sprintf("model_id = $%d", argIndex))
		values = append(values, *f.ModelId)
		argIndex++
	}

	where := ""
	if len(fields) > 0 {
		where = fmt.Sprintf("WHERE %s", strings.Join(fields, " AND "))
	}

	return where, values, argIndex
}

func (r *FilesSqlRepository) Count(f *models.FileFilter) (*int, error) {
	where, values, _ := r.where(f)
	query := fmt.Sprintf("SELECT COUNT(*) FROM files %s", where)
	row := r.db.QueryRow(query, values...)

	var count int
	if err := row.Scan(&count); err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *FilesSqlRepository) Delete(m *models.File) error {
	query := fmt.Sprintf("DELETE FROM files WHERE id = $1")
	_, err := r.db.Exec(query, m.ID)
	return err
}

func (r *FilesSqlRepository) Find(ID string) (*models.File, error) {
	m := &models.File{}
	query := fmt.Sprintf("SELECT id, original_file_name, ext, uuid, data, model_name, model_id, created_at, updated_at FROM files WHERE id = $1")
	row := r.db.QueryRow(query, ID)

	var DataString *string
	if err := row.Scan(&m.ID, &m.OriginalFileName, &m.Ext, &m.UUID, &DataString, &m.ModelName, &m.ModelId, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return nil, errRecordNotFound
	}

	if DataString != nil {
		if err := json.Unmarshal([]byte(*DataString), &m.Data); err != nil {
			return nil, err
		}
	}

	return m, nil
}
