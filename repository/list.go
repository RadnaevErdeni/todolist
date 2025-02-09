package repository

import (
	todo "TODO"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TodoListDB struct {
	db *sqlx.DB
}

func NewTodoListDB(db *sqlx.DB) *TodoListDB {
	return &TodoListDB{db: db}
}

func (r *TodoListDB) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1,$2) RETURNING id`, todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, nil
	}

	createUserListQuery := fmt.Sprintf(`INSERT INTO %s (user_id, list_id) VALUES ($1,$2)`, usersListsTable)
	_, err = tx.Exec(createUserListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}
	return id, tx.Commit()
}

func (r *TodoListDB) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf(`SELECT tl.id,tl.title,tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1`,
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}
func (r *TodoListDB) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListDB) Delete(userId, listId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	deleteStr := fmt.Sprintf(`DELETE FROM %s WHERE id IN (SELECT ts.id FROM %s ts INNER JOIN %s ls ON ts.id = ls.str INNER JOIN %s ul ON ls.list = ul.list_id
    WHERE ul.user_id = $1 AND ls.list =$2)`, todoStrTable, todoStrTable, listsStrTable, usersListsTable)
	_, err = tx.Exec(deleteStr, userId, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteListsStr := fmt.Sprintf(`DELETE FROM %s WHERE id IN (SELECT ls.id FROM %s ls INNER JOIN %s ul ON ls.list = ul.list_id
    WHERE ul.user_id = $1 AND ls.list = $2)`, listsStrTable, listsStrTable, usersListsTable)
	_, err = tx.Exec(deleteListsStr, userId, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteUserzList := fmt.Sprintf(`DELETE FROM %s tl WHERE tl.user_id=$1 AND tl.list_id=$2`, usersListsTable)
	_, err = tx.Exec(deleteUserzList, userId, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteTodoList := fmt.Sprintf(`DELETE FROM %s tl WHERE tl.id = $1`, todoListsTable)
	_, err = tx.Exec(deleteTodoList, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *TodoListDB) Update(userId, listId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, listId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
