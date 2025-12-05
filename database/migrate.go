package database

import "log"

func AutoMigrate() {

	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Error creating categories table:", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(200) NOT NULL,
			author VARCHAR(100),
			category_id INT REFERENCES categories(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		log.Fatal("Error creating books table:", err)
	}

	log.Println("Auto migration completed")
}
