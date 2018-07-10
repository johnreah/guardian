package guardian

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"time"
	"strconv"
	"fmt"
)

const connectionString = "server=localhost;port=1433;database=Articles;user id=articles;password=articles;"

func getVerifiedConnection(t *testing.T) *RepositorySql {
	repo, err := CreateRepositorySql(connectionString)
	assert.Nil(t, err, "Couldn't connect to DB")
	version, err := repo.GetVersion()
	assert.Nil(t, err, "Couldn't read DB version")
	assert.True(t, strings.Contains(version, "Microsoft SQL Server"))
	return repo
}

func TestConnect(t *testing.T) {
	getVerifiedConnection(t)
}

func TestCount(t *testing.T) {
	repo := getVerifiedConnection(t)
	assert.Equal(t, repo.Count(), 2)
}

func TestGet(t *testing.T) {
	repo := getVerifiedConnection(t)
	article, err := repo.Get("idString4")
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err, "Get failed")
	if article != nil {
		fmt.Printf("Article: %v", article)
	} else {
		fmt.Printf("TestGet: no article found\n")
	}
}

func TestPut(t *testing.T) {
	repo := getVerifiedConnection(t)
	idString := "id" + strconv.Itoa(time.Now().Second())
	article := makeTestArticle(idString,  time.Now(), "sqlTitle","sqlBody", "sqlJson")
	err := repo.Put(article)
	assert.Nil(t, err, "Put failed")
}

func TestGetMostRecent(t *testing.T) {
	repo := getVerifiedConnection(t)
	article, err := repo.GetMostRecent()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Article: %v", article)
	assert.Nil(t, err, "Put failed")
}

