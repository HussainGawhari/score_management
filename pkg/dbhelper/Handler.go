package dbhelper

// The dbhelper package in Go application
// is responsible for interacting with the database.
//  It acts as an abstraction layer that encapsulates
//   the database operations, making it easier to manage
//  and maintain database-related code throughout your application.

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"spotbuzz-backend/models"

	"github.com/joho/godotenv"
)

var db *sql.DB

// create connection with postgres db
func CreateConnection() (*sql.DB, error) {
	// load .env file which contains configuration of database connection
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	dbc, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = dbc.Ping()

	if err != nil {
		panic(err)
	}
	db = dbc // assign the connection to the golobal variable of database

	fmt.Println("Successfully connected!")

	return db, nil

}

func CreateTable(tableName string) (string, error) {

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY,
		name VARCHAR(15) NOT NULL,
		country VARCHAR(2) NOT NULL,
		score INTEGER NOT NULL
	)`, tableName)

	// Execute the query to create the table
	_, err := db.Exec(query)
	if err != nil {
		return "", err
	}

	return "Table created successfully", nil
}

func InsertScores(player models.Player) error {
	// var player models.Player
	fmt.Println("Insert results")
	// Query executes a query that returns rows, typically a SELECT
	result, err := db.Exec("INSERT INTO player (name, country, score) VALUES ($1, $2, $3)", player.Name, player.Country, player.Score)
	if err != nil {
		fmt.Println("Err", err.Error())
		return err
	}
	fmt.Println(result)

	return nil
}

// Get all players from the database in request format
func GetPlayers() ([]models.Player, error) {

	var players []models.Player
	rows, err := db.Query("SELECT * FROM player ORDER BY name DESC")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	for rows.Next() {
		var player models.Player

		err := rows.Scan(&player.ID, &player.Name, &player.Country, &player.Score)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return players, nil

}

func UpdatePlayer(player models.Player, id string) (models.Player, error) {

	// db.Prepare("SELECT id FROM player where id = $1")
	// _, err := db.Exec(id)
	// fmt.Print(id)
	// if err != nil {
	// 	fmt.Println("Error finding the player id:", err)
	// 	return player, err
	// } else {
	// Update the player's name and score in the database
	stmt, err := db.Prepare("UPDATE player SET name = $1, score = $2 WHERE id = $3")
	if err != nil {
		fmt.Println("Update player failed", err)
		return player, err
	}
	_, err = stmt.Exec(player.Name, player.Score, id)
	if err != nil {

		return player, err
	}
	// }

	return player, nil

}

func DeletePlayer(id string) error {

	// Execute the DELETE statement
	// _, err := db.Query("select id from player where id = ?", id)
	// if err != nil {
	// 	return err
	// }

	stmt, err := db.Prepare("DELETE FROM player WHERE id = $1")
	if err != nil {
		fmt.Print(err)
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement to delete the player
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func GetByRankFromDatabase(rank int) ([]models.Player, error) {

	var players []models.Player
	rows, err := db.Query("SELECT * FROM player where score =$1", rank)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	for rows.Next() {
		var player models.Player

		err := rows.Scan(&player.ID, &player.Name, &player.Country, &player.Score)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return players, nil
}

func GetRandomePlayerfromDatabase() ([]models.Player, error) {

	var players []models.Player
	rows, err := db.Query("SELECT * FROM player ORDER BY RANDOM()")

	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	for rows.Next() {
		var player models.Player

		err := rows.Scan(&player.ID, &player.Name, &player.Country, &player.Score)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		players = append(players, player)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error retrieving data:", err)
	}
	return players, nil

}
