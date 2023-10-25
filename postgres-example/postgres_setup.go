package main

import (
  "database/sql"
   "log"
   "fmt"

   _ "github.com/lib/pq"
)

// Ubuntu installation was as simple as:
//      sudo apt install postgresql
//      sudo systemctl restart postgresql.service
 
// Make sure your 'postgres' user has a password setup
//      sudo -u postgres psql
//      ALTER USER postgres WITH PASSWORD 'new_password';

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "new_password"
    dbname   = "testdb"
    tablename = "example"

    TABLE_EXIST = "SELECT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'example') AS table_existence;"
)

func drop_table(conninfo string) {
    conninfo = conninfo + fmt.Sprintf(" dbname=%s", dbname)
    db, err := sql.Open("postgres", conninfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec(fmt.Sprintf("DROP TABLE %s", tablename))
    if err != nil {
        log.Fatal(err)
    }
}

func create_table(conninfo string) {
    conninfo = conninfo + fmt.Sprintf(" dbname=%s", dbname)
    db, err := sql.Open("postgres", conninfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec("CREATE TABLE %s ( id integer, username varchar(255) )", tablename)
    if err != nil {
        log.Fatal(err)
    }
}

func create_db(conninfo string) {
    db, err := sql.Open("postgres", conninfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec("create database " + dbname)
    if err != nil {
        //handle the error
        log.Fatal(err)
    }
}

func query_table(conninfo string) {
    conninfo = conninfo + fmt.Sprintf(" dbname=%s", dbname)
    db, err := sql.Open("postgres", conninfo)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //r, err := db.Exec(fmt.Sprintf("SELECT * FROM %s;", tablename))
    r, err := db.Exec("SELECT * FROM fake_table")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("r = %v\n", r)
}
 
func main() {
    conninfo := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", user, password, host)
    // create_db(conninfo)
    //create_table(conninfo)
    //drop_table(conninfo)

    query_table(conninfo)
}
