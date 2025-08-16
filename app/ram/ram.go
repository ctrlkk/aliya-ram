package ram

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	_ "github.com/mattn/go-sqlite3"
)

type Memory struct {
	UserID  string
	Content string
}

type RAM struct {
	index   bleve.Index
	mapping *mapping.IndexMappingImpl
	db      *sql.DB
}

func NewRAM() (*RAM, error) {
	var err error
	blevePath := "./memories.bleve"
	dbPath := "./memories.db"
	mapping := bleve.NewIndexMapping()
	index, err := bleve.Open(blevePath)
	if err != nil {
		index, err = bleve.New(blevePath, mapping)
		if err != nil {
			return nil, err
		}
	}

	d, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = d.Exec(`
        CREATE TABLE IF NOT EXISTS user_memories (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id TEXT NOT NULL,
            content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
        CREATE INDEX IF NOT EXISTS idx_user_id ON user_memories(user_id);
    `)
	if err != nil {
		return nil, err
	}
	return &RAM{
		index:   index,
		mapping: mapping,
		db:      d,
	}, nil
}

func (r *RAM) Close() error {
	var err error
	err = r.db.Close()
	if err != nil {
		return err
	}
	err = r.index.Close()
	return err
}

func (r *RAM) AddMemory(userID string, content string) error {
	stmt, err := r.db.Prepare("INSERT INTO user_memories (user_id, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(userID, content)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	r.index.Index(strconv.FormatInt(id, 10), content)
	return nil
}

func (r *RAM) SearchMemory(userID, keyword string) ([]string, error) {
	query := bleve.NewQueryStringQuery(keyword)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := r.index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	var memories []string
	stmt, err := r.db.Prepare("SELECT content FROM user_memories WHERE id = ? AND user_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, hit := range searchResult.Hits {
		var message string
		row := stmt.QueryRow(hit.ID, userID)
		err := row.Scan(&message)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			return nil, err
		}
		memories = append(memories, message)
	}

	return memories, nil
}

func (r *RAM) ListMemories(userID string) ([]string, error) {
	var memories []string
	rows, err := r.db.Query("SELECT content FROM user_memories WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var content string
		err := rows.Scan(&content)
		if err != nil {
			return nil, err
		}
		memories = append(memories, content)
	}
	return memories, nil
}

func (r *RAM) DeleteAllMemories(userID string) error {
	sql := "DELETE FROM user_memories WHERE user_id = ?"
	_, err := r.db.Exec(sql, userID)
	if err != nil {
		return err
	}
	rows, err := r.db.Query("SELECT id FROM user_memories WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
		log.Println("delete id:", id)
		r.index.Delete(strconv.FormatInt(int64(id), 10))
	}
	return nil
}

func (r *RAM) CreateIndexs(userID string) error {
	rows, err := r.db.Query("SELECT id, content FROM user_memories WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var content string
		err := rows.Scan(&id, &content)
		if err != nil {
			return err
		}
		r.index.Index(strconv.FormatInt(int64(id), 10), content)
	}
	return nil
}

func (r *RAM) UpdateUserID(oldUserID string, newUserID string) (sql.Result, error) {
	result, err := r.db.Exec("UPDATE user_memories SET user_id = ? WHERE user_id = ?", newUserID, oldUserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
