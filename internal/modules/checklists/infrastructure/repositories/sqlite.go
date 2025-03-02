package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mattn/go-sqlite3"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/domainErrors"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/domain"
)

type SqliteChecklistsRepository struct {
	db *sql.DB
}

func NewSqliteChecklistsRepository(db *sql.DB) *SqliteChecklistsRepository {
	return &SqliteChecklistsRepository{db: db}
}

func (r *SqliteChecklistsRepository) List() ([]*domain.Checklist, error) {
	rows, err := r.db.Query("SELECT id, title, created_at, updated_at, tasks FROM checklists")
	if err != nil {
		return nil, r.mapSqliteError(err)
	}

	var checklists []*domain.Checklist

	for rows.Next() {
		var dbChecklist struct {
			ID        string
			Title     string
			CreatedAt time.Time
			UpdatedAt time.Time
			Tasks     string
		}

		if err := rows.Scan(&dbChecklist.ID, &dbChecklist.Title, &dbChecklist.CreatedAt, &dbChecklist.UpdatedAt, &dbChecklist.Tasks); err != nil {
			return nil, r.mapSqliteError(err)
		}

		var tasks []domain.Task
		if err := json.Unmarshal([]byte(dbChecklist.Tasks), &tasks); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
		}

		checklists = append(checklists, &domain.Checklist{
			ID:    dbChecklist.ID,
			Title: dbChecklist.Title,
			Tasks: tasks,
		})
	}

	return checklists, nil
}

func (r *SqliteChecklistsRepository) CreateOne(checklist *domain.Checklist) (*domain.Checklist, error) {
	jsonTasks, err := json.Marshal(checklist.Tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tasks: %w", err)
	}

	_, err = r.db.Exec("INSERT INTO checklists (id, title, tasks) VALUES (?, ?, ?)", checklist.ID, checklist.Title, jsonTasks)
	if err != nil {
		return nil, r.mapSqliteError(err)
	}

	return checklist, nil
}

func (r *SqliteChecklistsRepository) FindOne(id string) (*domain.Checklist, error) {
	row := r.db.QueryRow("SELECT id, title, created_at, updated_at, tasks FROM checklists WHERE id = ?", id)

	var dbChecklist struct {
		ID        string
		Title     string
		CreatedAt time.Time
		UpdatedAt time.Time
		Tasks     string
	}
	if err := row.Scan(&dbChecklist.ID, &dbChecklist.Title, &dbChecklist.CreatedAt, &dbChecklist.UpdatedAt, &dbChecklist.Tasks); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainErrors.ErrEntityNotFound
		}
		return nil, r.mapSqliteError(err)
	}

	var tasks []domain.Task
	if err := json.Unmarshal([]byte(dbChecklist.Tasks), &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return &domain.Checklist{
		ID:    id,
		Title: dbChecklist.Title,
		Tasks: tasks,
	}, nil
}

func (r *SqliteChecklistsRepository) SaveOne(checklist *domain.Checklist) (*domain.Checklist, error) {
	return nil, nil
}

func (r *SqliteChecklistsRepository) DeleteOne(id string) error {
	_, err := r.db.Exec("DELETE FROM checklists WHERE id = ?", id)
	if err != nil {
		return r.mapSqliteError(err)
	}
	return nil
}

func (r *SqliteChecklistsRepository) mapSqliteError(err error) error {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return domainErrors.ErrEntityConflict
		}
	}
	return fmt.Errorf("database error: %w", err)
}
