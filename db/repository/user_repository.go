package repository

import (
	"database/sql"
	"gofiber/db/repository/models"
	"strings"
)

type IUserRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (data []models.User, count int, err error)

	Read(ID string) (data models.User, err error)

	Edit(model models.User) (res string, err error)

	Add(model models.User) (res string, err error)

	Delete(model models.User) (res string, err error)

	CountByEmail(ID,email string) (res int, err error)
}

type UserRepository struct {
	DB *sql.DB
}

const (
	userSelectStatement = `select id,email,password,created_at,updated_at`
	userWhereStatement  = `where lower(email) like $1 and deleted_at is null`
)

//new user repository
func NewUserRepository(DB *sql.DB) IUserRepository {
	return &UserRepository{DB: DB}
}

//scan rows
func (repository UserRepository) scanRows(rows *sql.Rows) (res models.User, err error) {
	err = rows.Scan(&res.ID, &res.Email, &res.Password, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//scan row
func (repository UserRepository) scanRow(row *sql.Row) (res models.User, err error) {
	err = row.Scan(&res.ID, &res.Email, &res.Password, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

//browse
func (repository UserRepository) Browse(search, orderBy, sort string, limit, offset int) (data []models.User, count int, err error) {
	statement := userSelectStatement + ` from "users" ` + userWhereStatement + ` order by ` + orderBy + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}

	statement = `select count(id) from "users" ` + userWhereStatement
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

//read
func (repository UserRepository) Read(ID string) (data models.User, err error) {
	statement := userSelectStatement + ` from "users" where id=$1 and deleted_at is null`
	row := repository.DB.QueryRow(statement, ID)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

//edit
func (repository UserRepository) Edit(model models.User) (res string, err error) {
	updateStatement := `set email=$1,updated_at=$2 where id=$3`
	updatedParams := []interface{}{model.Email, model.UpdatedAt, model.ID}
	if model.Password != "" {
		updateStatement = `set email=$1,updated_at=$2,password=$4 where id=$3`
		updatedParams = append(updatedParams, model.Password)
	}

	statement := `update "users" ` + updateStatement
	err = repository.DB.QueryRow(statement, updatedParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

//add
func (repository UserRepository) Add(model models.User) (res string, err error) {
	statement := `insert into "users" (email,password,created_at,updated_at) values($1,$2,$3,$4) returning id`
	err = repository.DB.QueryRow(statement, model.Email, model.Password, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

//delete
func (repository UserRepository) Delete(model models.User) (res string, err error) {
	statement := `update "users" set updated_at=$1, deleted_at=$2 where id=$3`
	err = repository.DB.QueryRow(statement, model.UpdatedAt, model.DeletedAt.Time, model.ID).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

//count
func (repository UserRepository) CountByEmail(ID,email string) (res int, err error) {
	whereStatement := `where email=$1 and deleted_at is null`
	whereParams := []interface{}{email}
	if ID != ""{
		whereStatement = `where (email=$1 and deleted_at is null) and id <>$2`
		whereParams = append(whereParams,ID)
	}
	statement := `select count(id) from "users" `+whereStatement
	err = repository.DB.QueryRow(statement, whereParams...).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
