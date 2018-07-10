package guardian

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
	"log"
	"time"
)

// RepositorySql is an implementation of the Repository interface for SQL Server, It uses
// denisenkom's go-mssqldb package, which appears to be the most popular MSSQL driver, but this
// has numerous problems, especially when used on a Windows client:
// 1) I couldn't make integrated authentication work, and had to fall back to using SQL auth.
// 2) Parameter substitution doesn't handle inserting time.Time values into a DateTimeOffset
//    field, interpreting it as money. The workaround was to construct the entire query string,
//    which would generally be a SQL injection risk. This may be an sp_executesql problem rather
//    than go-mssqldb. UPDATE: change $1 from literature to @p1 for sp_executesql.
// 3) LastInsertId doesn't work. Apparently it used to but is no longer supported. The
//    workaround would be to use an OUTPUT clause or call SCOPE_IDENTITY().
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

func (repo *RepositorySql) Get(articleId string) (Article, error) {
	rows, err := repo.db.Query("select top 1 articleId, articleDate, title, body, json from articles where articleId=@p1", articleId)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var articleId, title, body, json string
		var articleDate time.Time
		err := rows.Scan(&articleId, &articleDate, &title, &body, &json)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
			return nil, err
		}
		fmt.Printf("id=%s\n", articleId)
		return &TestArticle{articleId, articleDate, title, body, json}, nil
	}
	return nil, nil
}

func (repo *RepositorySql) Put(article Article) error {
	result, err := repo.db.Exec("insert into Articles(articleId, articleDate, title, body, json) values (@p1, @p2, @p3, @p4, @p5)",
		article.ArticleId(), article.ArticleDate(), article.Title(), article.Body(), article.Json())
	//s := article.ArticleDate().Format("2006-01-02 15:04:05.999999 -07:00")
	//query := fmt.Sprintf("insert into Articles(idString, title, articleDate, bodyText, json) values ('%s', '%s', '%s', '%s', '%s')",
	//	article.ArticleId(), article.Title(), s, article.Body(), article.Json())
	//result, err := repo.db.Exec(query)
	if err != nil {
		return err
	}
	rowsAffected, rowsErr := result.RowsAffected()
	if rowsErr != nil {
		fmt.Printf("Failed to get rowsAffected: %v\n", rowsErr)
	}
	if rowsAffected != 1 {
		return fmt.Errorf("Insert succeeded but rowsAffected = %d instead of 1", rowsAffected)
	}
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

func (repo *RepositorySql) GetMostRecent() (Article, error) {
	rows, err := repo.db.Query("select top 1 articleId, articleDate, title, body, json from articles order by articleDate desc")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var articleId, title, body, json string
		var articleDate time.Time
		err := rows.Scan(&articleId, &articleDate, &title, &body, &json)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
			return nil, err
		}
		fmt.Printf("id=%s\n", articleId)
		return &TestArticle{articleId, articleDate, title, body, json}, nil
	}
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

