package repository

import (
	"database/sql"
	"session-2/params/request"
	"session-2/params/view"
)

type TodoRepository struct {
	DB *sql.DB
}

func (s TodoRepository) InsertTodo(input request.CreateTodo) error {
	query := `
    INSERT INTO todo (
        title,
        description
    ) VALUES (
        ?,?
    )
    `

	_, err := s.DB.Exec(query, input.Title, input.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s TodoRepository) GetAllTodos() (todos []view.GetTodo, err error) {
	query := `
    SELECT
        todo_id,
        title,
        description,
        created_at
    FROM todo
    `

	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo view.GetTodo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s TodoRepository) GetTodoByID(id string) (todo view.GetTodo, err error) {
	query := `
    SELECT
        todo_id,
        title,
        description,
        created_at
    FROM todo
    WHERE todo_id = ?
    `

	err = s.DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.CreatedAt,
	)
	if err != nil {
		return view.GetTodo{}, err
	}

	return todo, nil
}

func (s TodoRepository) DeleteTodoByID(id string) (err error) {
	query := `
    DELETE FROM todo
    WHERE todo_id = ?
    `

	_, err = s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s TodoRepository) UpdateTodoByID(id string, input request.UpdateTodo) (err error) {
	query := `
    UPDATE todo
    SET
        title = ?,
        description = ?
    WHERE todo_id = ?
    `

	_, err = s.DB.Exec(query, input.Title, input.Description, id)
	if err != nil {
		return err
	}

	return nil
}
