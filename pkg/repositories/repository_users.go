package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/helpers"
	"github.com/nurbeknurjanov/go-gin-backend/pkg/models"
	"reflect"
	"strings"
)

type UsersSqlRepository struct {
	db *sqlx.DB
}

func newUsersSqlRepository(sqlDb *sqlx.DB) *UsersSqlRepository {
	return &UsersSqlRepository{db: sqlDb}
}

func (r *UsersSqlRepository) Create(u *models.User) error {
	query := "INSERT INTO users (email, encrypted_password, name, age, sex, status) values ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	row := r.db.QueryRow(query, u.Email, u.Password, u.Name, u.Age, u.Sex, u.Status)
	return row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}

func (r *UsersSqlRepository) Update(u *models.User, data *models.UserPartial) error {
	updateFields := []string{}
	argIndex := 1
	values := []any{}

	Object := reflect.ValueOf(data).Elem()
	for i := 0; i < Object.NumField(); i++ {
		Field := Object.Field(i)
		if Field.Kind() == reflect.Ptr && !Field.IsNil() {
			updateFields = append(updateFields, fmt.Sprintf("%s = $%d", helpers.ToSnakeCase(Object.Type().Field(i).Name), argIndex))
			argIndex++
			values = append(values, Field.Elem().Interface())
		}
		if Field.Kind() == reflect.Struct {
			for j := 0; j < Field.NumField(); j++ {
				SubField := Field.Field(j)
				if SubField.Kind() == reflect.Ptr && !SubField.IsNil() {
					updateFields = append(updateFields, fmt.Sprintf("%s = $%d", helpers.ToSnakeCase(Field.Type().Field(j).Name), argIndex))
					argIndex++
					values = append(values, SubField.Elem().Interface())
				}
			}
		}
	}

	values = append(values, u.ID)

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d RETURNING email, name, age, sex, status, updated_at",
		strings.Join(updateFields, ", "), argIndex)
	row := r.db.QueryRow(query, values...)

	return row.Scan(&u.Email, &u.Name, &u.Age, &u.Sex, &u.Status, &u.UpdatedAt)
}

func (r *UsersSqlRepository) Delete(u *models.User) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id = $1")
	_, err := r.db.Exec(query, u.ID)
	return err
}

func (r *UsersSqlRepository) Find(ID string) (*models.User, error) {
	u := &models.User{}
	query := "SELECT id, email, name, age, sex, status, created_at, updated_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, ID)
	if err := row.Scan(&u.ID, &u.Email, &u.Name, &u.Age, &u.Sex, &u.Status, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, errRecordNotFound
	}

	return u, nil
}

func (r *UsersSqlRepository) List(p *PaginationRequest, s *Sort, f *models.UserFilter) ([]*models.User, error) {
	list := []*models.User{}

	where, values, argIndex := r.where(f)

	values = append(values, p.PageSize)
	values = append(values, p.PageSize*p.PageNumber)

	orderBy := fmt.Sprintf("ORDER BY %s %s", s.SortField, s.SortDirection)

	limit := fmt.Sprintf("LIMIT $%d OFFSET $%d", argIndex, argIndex+1)

	query := fmt.Sprintf(
		"SELECT id, email, name, age, sex, status, created_at, updated_at FROM users %s %s %s",
		where, orderBy, limit)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := &models.User{}
		err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.Age, &u.Sex, &u.Status, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
			//continue
		}
		list = append(list, u)
	}

	return list, nil
}

func (r *UsersSqlRepository) where(f *models.UserFilter) (string, []any, int) {
	whereFields := []string{}
	argIndex := 1
	values := []any{}

	if f.Name != nil {
		whereFields = append(whereFields, fmt.Sprintf("name LIKE $%d", argIndex))
		values = append(values, "%"+*f.Name+"%")
		argIndex++
	}
	if f.Email != nil {
		whereFields = append(whereFields, fmt.Sprintf("email LIKE $%d", argIndex))
		values = append(values, "%"+*f.Email+"%")
		argIndex++
	}
	if f.Age != nil {
		whereFields = append(whereFields, fmt.Sprintf("age = $%d", argIndex))
		values = append(values, *f.Age)
		argIndex++
	}
	if f.Sex != nil {
		sexConditions := []string{}
		for _, val := range *f.Sex {
			sexConditions = append(sexConditions, fmt.Sprintf("sex = $%d", argIndex))
			argIndex++
			values = append(values, val)
		}
		whereFields = append(whereFields, "("+strings.Join(sexConditions, " OR ")+")")
	}
	if f.Status != nil {
		statusConditions := []string{}
		for _, val := range *f.Status {
			statusConditions = append(statusConditions, fmt.Sprintf("status = $%d", argIndex))
			argIndex++
			values = append(values, val)
		}
		whereFields = append(whereFields, "("+strings.Join(statusConditions, " OR ")+")")
	}

	if f.CreatedAtFrom != nil {
		whereFields = append(whereFields, fmt.Sprintf("created_at >= $%d", argIndex))
		values = append(values, *f.CreatedAtFrom)
		argIndex++
	}
	if f.CreatedAtTo != nil {
		whereFields = append(whereFields, fmt.Sprintf("created_at <= $%d", argIndex))
		values = append(values, *f.CreatedAtTo)
		argIndex++
	}
	if f.UpdatedAtFrom != nil {
		whereFields = append(whereFields, fmt.Sprintf("updated_at >= $%d", argIndex))
		values = append(values, *f.UpdatedAtFrom)
		argIndex++
	}
	if f.UpdatedAtTo != nil {
		whereFields = append(whereFields, fmt.Sprintf("updated_at <= $%d", argIndex))
		values = append(values, *f.UpdatedAtTo)
		argIndex++
	}

	where := ""
	if len(whereFields) > 0 {
		where = fmt.Sprintf("WHERE %s", strings.Join(whereFields, " AND "))
	}

	return where, values, argIndex
}
func (r *UsersSqlRepository) Count(f *models.UserFilter) (*int, error) {
	where, values, _ := r.where(f)
	query := fmt.Sprintf("SELECT COUNT(*) FROM users %s", where)
	row := r.db.QueryRow(query, values...)

	var count int
	if err := row.Scan(&count); err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *UsersSqlRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	query := "SELECT id, email, encrypted_password AS password, name, age, sex, status, created_at, updated_at FROM users WHERE email=$1"
	//err := r.db.Get(&u, query, email)

	row := r.db.QueryRow(query, email)
	if err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Name, &u.Age, &u.Sex, &u.Status, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UsersSqlRepository) ChangeUserPassword(u *models.User, password string) error {
	query := "UPDATE users SET encrypted_password = $1 WHERE id = $2"
	return r.db.QueryRow(query, password, u.ID).Err()
}
