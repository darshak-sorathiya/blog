package main

import (
	"blog/cmd/server"
	"blog/config"
	"blog/internal/repository"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// main function bootstrap server

// @title Blog
// @version 2.0
func main() {
	r := mux.NewRouter()
	appConfig := config.LoadAndGetConfig()
	// Initialize the SQL database connection
	// dsn := "root:@tcp(127.0.0.1:3306)/blog"
	dsn := appConfig.MysqlUser + ":@tcp(127.0.0.1:3306)/" + appConfig.MysqlDatabase
	db, err := sql.Open("mysql", dsn)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	articleRepo := repository.NewArticleRepository(db)

	server.SetupRoutes(r, articleRepo)

	http.Handle("/", r)
	port := appConfig.AppPort
	log.Println("server has been started on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
