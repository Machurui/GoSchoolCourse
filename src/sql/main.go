package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "1234",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}

	fmt.Println("Connexion réussie")

	albums, err := albumByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Albums :", albums)

	album, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Album found :", album)

	newAlbum := Album{
		Title:  "Foo",
		Artist: "Bar",
		Price:  78.99,
	}

	id, err := addAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New album id :", id)
}

func albumByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q: %v", name, err)
		}

		albums = append(albums, album)
	}

	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var album Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("no album with id=%d", id)
		}
		return album, fmt.Errorf("albumByID %d: %v", id, err)
	}

	return album, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum %v", err)
	}

	return id, nil
}
