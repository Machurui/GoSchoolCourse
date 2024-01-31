package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

const server = "http://localhost:8080/albums"

func main() {
	albums := GetAll()
	fmt.Printf("Albums: %v\n", albums)

	album := GetOne(2)
	fmt.Printf("Albums: %v\n", album)

	// album = GetOne(10)
	// fmt.Printf("Albums: %v\n", album)

	postAlbum := Album{ID: "5", Title: "Test", Artist: "AHAH"}
	postAlbum = Post(postAlbum)
	fmt.Println("Album:", postAlbum)
}

func GetAll() []Album {
	url := server
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Status code:", resp.StatusCode)
	}

	var albums []Album
	bodyBytes, _ := io.ReadAll(resp.Body)

	json.Unmarshal(bodyBytes, &albums)

	return albums
}

func GetOne(id int) Album {
	url := fmt.Sprintf("%s/%d", server, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Status code:", resp.StatusCode)
	}

	var album Album
	bodyBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &album)

	return album
}

func Post(album Album) Album {
	jsonReq, _ := json.Marshal(album)
	resp, err := http.Post(server, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalln("Status code:", resp.StatusCode)
	}

	var posted Album
	bodyBytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &posted)

	return posted
}
