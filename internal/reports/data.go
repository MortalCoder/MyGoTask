package reports

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// Добавить репорт
func (r *Repo) CreateReport(title, description string) error {
	_, err := r.db.Exec(
		`INSERT INTO reports (title, description) VALUES ($1, $2)`,
		title, description,
	)
	return err
}

// Получить репорт по id
func (r *Repo) GetReportById(id int) (*Report, error) {
	row := r.db.QueryRow(`SELECT id, title, description, created_at, updated_at FROM reports WHERE id = $1`, id)

	var report Report
	err := row.Scan(&report.ID, &report.Title, &report.Description, &report.CreatedAt, &report.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &report, nil
}

// Обновить репорт
func (r *Repo) UpdateReport(id int, title, description string) error {
	_, err := r.db.Exec(
		`UPDATE reports SET title = $1, description = $2, updated_at = NOW() WHERE id = $3`,
		title, description, id,
	)
	return err
}

// Удалить репорт
func (r *Repo) DeleteReport(id int) error {
	_, err := r.db.Exec(`DELETE FROM reports WHERE id = $1`, id)
	return err
}
