package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	DB *sql.DB
}

var repo Repository

func Init() {
	if db, err := sql.Open("sqlite3", "/tmp/audio.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		log.Fatal("Database initialisation")
	}
}

func Create() int {
	const sql = "CREATE TABLE IF NOT EXISTS Cells" +
		"(Id TEXT PRIMARY KEY, Audio TEXT)"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Clear() int {
	const sql = "DELETE FROM Cells"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

func Read(id string) (Cell, int64) {
	const sql = "SELECT * FROM Cells WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		var c Cell
		row := stmt.QueryRow(id)
		if err := row.Scan(&c.Id, &c.Audio); err == nil {
			return c, 1
		} else {
			return Cell{}, 0
		}
	}
	return Cell{}, -1
}

func Update(c Cell) int64 {
	const sql = "UPDATE Cells SET Audio = ? " + "WHERE id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(c.Audio, c.Id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Insert(c Cell) int64 {
	const sql = "INSERT INTO Cells(Id, Audio) " + "VALUES (?, ?)"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(c.Id, c.Audio); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func Delete(id string) int64 {
	const sql = "DELETE FROM Cells WHERE Id = ?"
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

func All() ([]Cell, int64) {
	const sql = "SELECT Id FROM Cells"
	if rows, err := repo.DB.Query(sql); err == nil {
		defer rows.Close()

		var cells []Cell
		for rows.Next() {
			var cell Cell
			err = rows.Scan(&cell.Id)
			if err != nil {
				return nil, -1
			}
			cells = append(cells, cell)
		}
		return cells, 1
	}
	return nil, -1
}
