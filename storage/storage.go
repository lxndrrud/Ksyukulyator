package storage

import "github.com/jmoiron/sqlx"

func InitStorage(db *sqlx.DB) error {
	_, err := db.Exec(`
		PRAGMA foreign_keys=ON;
		CREATE TABLE IF NOT EXISTS categories(
			category_id integer not null primary key,
			category_title varchar(100) not null
		);
		CREATE TABLE IF NOT EXISTS products(
			product_id integer not null primary key, 
			product_title varchar(100) not null,
			product_cost float not null,
			product_id_category integer,
			FOREIGN KEY(product_id_category) REFERENCES categories(category_id) ON DELETE CASCADE
		);
		 `)
	if err != nil {
		return err
	}
	return nil
}
