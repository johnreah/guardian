package guardian

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
)

type RepositorySql struct {
	db *sql.DB
}

func (repo *RepositorySql) connect(connString string) error{
	var err error
	db, err := sql.Open("sqlserver", connString)
	if err == nil {
		repo.db = db
	}
	return err
}

func (*RepositorySql) Get(id string) (*Article, error) {
	return nil, nil
}

func (repo *RepositorySql) Put(article Article) error {
	s := article.ArticleDate().Format("2006-01-02 15:04:05.999999 -07:00")
	//fmt.Printf("Formatted time: %s\n", s)
	//result, err := repo.db.Exec("insert into Articles(idString, title, articleDate, bodyText, json) values ($1, $2, $3, $4, $5)",
	//	article.IdString(), article.Title(), s, article.Body(), article.Json())
	query := fmt.Sprintf("insert into Articles(idString, title, articleDate, bodyText, json) values ('%s', '%s', '%s', '%s', '%s')",
		article.IdString(), article.Title(), s, article.Body(), article.Json())
	result, err := repo.db.Exec(query)
	if err != nil {
		return err
	}
	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Result: LastInsertId=%d, RowsAffected=%d\n", lastInsertId, rowsAffected)
	return nil
}

func (repo *RepositorySql) Count() int {
	rows, err := repo.db.Query("select count(*) from articles")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var numRows int
	if rows.Next() {
		err := rows.Scan(&numRows)
		if err != nil {
			panic(err)
		}
	}
	return numRows
}

func (*RepositorySql) GetMostRecent() (*Article, error) {
	return nil, nil
}

func CreateRepositorySql(connString string) (*RepositorySql, error) {
	repo := RepositorySql{}
	err := repo.connect(connString)
	if err != nil {
		panic(err)
	}
	return &repo, err
}

func (repo *RepositorySql) GetVersion() (string, error) {
	rows, err := repo.db.Query("select @@version")
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var version string
	if rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			return"",  err
		}
	}
	return version, nil
}

