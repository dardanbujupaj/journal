package entry

import (
	"time"
	"fmt"
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
)

type Entry struct {
	Id int64
	CreatedDate time.Time
	Content string
}

// return sqlite database
// creates a new database if none is found
func getDatabase() (*sql.DB, error) {
	var database *sql.DB

	expanded, err := homedir.Expand("~/.journal.db")

	if err == nil {
		database, err = sql.Open("sqlite3", expanded)
		if err == nil {

			statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS entries (id INTEGER PRIMARY KEY, date TIMESTAMP, content TEXT)")
			statement.Exec()
		}
	}

	return database, err
}

// create a new entry with given text and current timestamp
func CreateEntry(content string) {
	
	database, _ := getDatabase()

	statement, _ := database.Prepare("INSERT INTO entries (date, content) VALUES (CURRENT_TIMESTAMP, ?)")
	statement.Exec(content)

    fmt.Println("Journal entry saved.")
}

// if amount <= 0 (should be -1) all entries are printed
// else the given amount of lates entries are printed
func ListEntries(amount int) {
	database, _ := getDatabase()

    var query string

    if (amount >= 0) {
        query = fmt.Sprintf("SELECT * FROM (SELECT * FROM entries ORDER BY date DESC LIMIT %d) ORDER BY date ASC", amount)
    } else {
        query = "SELECT * FROM entries"
    }
	rows, _ := database.Query(query)


	var id int
	var date string
	var content string
    
	for rows.Next() {
		rows.Scan(&id, &date, &content)
		fmt.Println(fmt.Sprintf("#%d", id) + " - " + date)
        fmt.Println(content)
        fmt.Println()
	}
}
