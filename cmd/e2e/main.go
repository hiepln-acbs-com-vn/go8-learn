package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gmhafiz/go8/internal/domain/book"
	"github.com/gmhafiz/go8/internal/server"
)

// Version is injected using ldflags during build time
const Version = "v0.7.0-test"

func main() {
	log.Printf("Starting API version: %s\n", Version)
	s := server.New()
	s.Init(Version)
	s.Migrate()

	t := NewE2eTest(s)
	t.Run()
}

type E2eTest struct {
	server *server.Server
}

func NewE2eTest(server *server.Server) *E2eTest {
	return &E2eTest{
		server: server,
	}
}

func (t *E2eTest) Run() {
	testEmptyBook(t)
	id := testAddOneBook(t)
	id = testGetOneBook(t, id)
	testUpdateBook(t, id)
	testDeleteOneBook(t, id)

	log.Println("all tests passed.")
}

func testEmptyBook(t *E2eTest) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/api/v1/books",
		t.server.Config().Api.Port))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		log.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal(make([]*book.Res, 0))

	if !bytes.Equal(expected, got) {
		log.Printf("handler returned unexpected body: got %v want %v", string(got), expected)
	}

	log.Println("testEmptyBook passes")
}

func testAddOneBook(t *E2eTest) int {
	want := &book.CreateRequest{
		Title:         "test01",
		PublishedDate: "2020-02-02",
		ImageURL:      "https://example.com/image.png",
		Description:   "test01",
	}

	bR, _ := json.Marshal(want)

	resp, err := http.Post(
		fmt.Sprintf("http://localhost:%s/api/v1/books",
			t.server.Config().Api.Port),
		"Content-Type: application/json",
		bytes.NewBuffer(bR),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	gotBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	got := book.Res{}
	err = json.Unmarshal(gotBody, &got)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Printf("error code want %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	if want.Title != got.Title && want.Description != got.Description && want.
		ImageURL != got.ImageURL && want.PublishedDate != got.PublishedDate.String() {
		log.Printf("want %v, got %v\n", want, got)
	}

	log.Println("testAddOneBook passes")
	return got.ID
}

func testGetOneBook(t *E2eTest, id int) int {
	client := &http.Client{}

	url := fmt.Sprintf("http://localhost:%s/api/v1/books/%d", t.server.Config().Api.Port, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error code fail, want %d, got %d\n", http.StatusOK, resp.StatusCode)
	}

	got := book.Res{}
	err = json.Unmarshal(respBody, &got)
	if err != nil {
		log.Println(err)
	}

	log.Println("testGetBook passes")

	return got.ID
}

func testUpdateBook(t *E2eTest, bookID int) {
	newBook := book.CreateRequest{
		Title:         "updated title",
		PublishedDate: "2020-07-31T15:04:05.123499999Z",
		ImageURL:      "https://example.com/image.png",
		Description:   "test description",
	}

	client := &http.Client{}

	bR, err := json.Marshal(&newBook)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("http://localhost:%s/api/v1/books/%d", t.server.Config().Api.Port, bookID)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bR))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error code fail, want %d, got %d\n", http.StatusOK, resp.StatusCode)
	}

	got := book.Res{}
	err = json.Unmarshal(respBody, &got)
	if err != nil {
		log.Println(err)
	}

	if got.ID != bookID && got.Title != newBook.Title && got.Description != newBook.Description && got.ImageURL != newBook.ImageURL {
		if err != nil {
			log.Fatalf("returned resource does not match. want %v, got %v", respBody, got)
		}
	}

	log.Println("testUpdateBook passes")
}

func testDeleteOneBook(t *E2eTest, id int) {
	client := &http.Client{}

	req, err := http.NewRequest(
		http.MethodDelete, fmt.Sprintf("http://localhost:%s/api/v1/books/%d", t.server.Config().Api.Port, id),
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error code fail, want %d, got %d\n", http.StatusOK, resp.StatusCode)
	}
	log.Println("testDeleteOneBook passes")
}
