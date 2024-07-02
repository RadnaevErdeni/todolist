package repository

import (
	todo "TODO"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TodoStrDB struct {
	db *sqlx.DB
}

func NewTodoStrDB(db *sqlx.DB) *TodoStrDB {
	return &TodoStrDB{db: db}
}

func (r *TodoStrDB) Create(listId int, str todo.TodoStr) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var strId int
	createStrQuery := fmt.Sprintf("INSERT INTO %s (title, description,done) values ($1, $2, $3) RETURNING id", todoStrTable)

	row := tx.QueryRow(createStrQuery, str.Title, str.Description, str.Done)
	err = row.Scan(&strId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListStrQuery := fmt.Sprintf("INSERT INTO %s (list, str) values ($1, $2)", listsStrTable)
	_, err = tx.Exec(createListStrQuery, listId, strId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return strId, tx.Commit()
}
func (r *TodoStrDB) GetById(userId, strId int) (todo.TodoStr, error) {
	var str todo.TodoStr
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.str = ti.id
									INNER JOIN %s ul on ul.list_id = li.list WHERE ti.id = $1 AND ul.user_id = $2`, todoStrTable, listsStrTable, usersListsTable)
	if err := r.db.Get(&str, query, strId, userId); err != nil {
		return str, err
	}

	return str, nil
}

func (r *TodoStrDB) GetAll(userId, strId int) ([]todo.TodoStr, error) {
	var str []todo.TodoStr
	query := fmt.Sprintf(`SELECT ti.id,ti.title,ti.description,ti.done FROM %s ti INNER JOIN %s li ON li.str = ti.id INNER JOIN %s ul ON ul.list_id=li.list WHERE li.list = $1 AND ul.user_id = $2`, todoStrTable, listsStrTable, usersListsTable)
	if err := r.db.Select(&str, query, strId, userId); err != nil {
		return nil, err
	}
	return str, nil
}

// переделать запрос(плохо работает)
func (r *TodoStrDB) Delete(userId, strId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	deleteStrQuery := fmt.Sprintf(`DELETE FROM %s ts USING %s ls, %s ul WHERE ls.str = ts.id AND ls.list = ul.list_id AND ul.user_id = $1 AND ts.id =$2`, todoStrTable, listsStrTable, usersListsTable)
	_, err = tx.Exec(deleteStrQuery, userId, strId)
	if err != nil {
		tx.Rollback()
		return err
	}
	delete2StrQuery := fmt.Sprintf("DELETE FROM %s ls USING %s ts, %s ul WHERE ls.str = ts.id   AND ls.list = ul.list_id AND ul.user_id = $1 AND ts.id =$2 ", listsStrTable, todoStrTable, usersListsTable)
	_, err = tx.Exec(delete2StrQuery, userId, strId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// что то не так
func (r *TodoStrDB) UpdateStr(userId, strId int, input todo.UpdateStrInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	//может как то поменять тут логику на вызов чего либо, а не вот три подряд if если будет больше полей в табличке то все их так проверять не очень как то звучит(подумать/проверить)
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ts SET %s FROM %s ls, %s ul	WHERE ts.id = ls.str AND ls.list = ul.list_id AND ul.user_id = $%d AND ts.id = $%d",
		todoStrTable, setQuery, listsStrTable, usersListsTable, argId, argId+1)
	args = append(args, strId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
