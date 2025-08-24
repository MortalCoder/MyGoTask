package words

import (
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// RGetWordById ищем слово по id
func (r *Repo) RGetWordById(id int) (*Word, error) {
	var word Word
	err := r.db.QueryRow(`SELECT id, title, translation FROM ru_en WHERE id = $1`, id).
		Scan(&word.Id, &word.Title, &word.Translation)
	if err != nil {
		return nil, err
	}

	return &word, nil
}

// CreateNewWords добавляет новые переводы в базу даных
func (r *Repo) CreateNewWords(word, translate string) error {
	_, err := r.db.Exec(`INSERT INTO ru_en (title, translation) VALUES ($1, $2)`, word, translate)
	if err != nil {
		return err
	}

	return nil
}

// обновление слова по айди
func (r *Repo) UpdateWord(id int, word, translation string) error {
	_, err := r.db.Exec(
		`UPDATE ru_en SET title = $1, translation = $2 WHERE id = $3`,
		word, translation, id,
	)
	return err
}

// удаление слова по айди
func (r *Repo) DeleteWord(id int) error {
	_, err := r.db.Exec(`DELETE FROM ru_en WHERE id = $1`, id)
	return err
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
