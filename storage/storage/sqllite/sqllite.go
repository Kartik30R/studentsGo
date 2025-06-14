package sqllite

import (
	"database/sql"
	"fmt"

	"github.com/Kartik30R/studentsGo.git/internal/config"
	"github.com/Kartik30R/studentsGo.git/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type SqlLite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*SqlLite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
 id INTEGER PRIMARY KEY AUTOINCREMENT
 namae TEXT,
 email TEXT
 age INTEGER  
   )`)

	if err != nil {
		return nil, err
	}

	return &SqlLite{Db: db}, nil
}

func (s *SqlLite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name,email, age) VALUES (?,?,?)")

	if err != nil {

		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *SqlLite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM student WHERE id = ? LIMIT 1")

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: % w", err)
	}

	return student, nil
}

func (s *SqlLite)GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM STUDENTS")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil

}
