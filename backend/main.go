package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/storage_handler"
	"net/http"
	"os"

	"github.com/gookit/ini/v2"

	_ "github.com/lib/pq"
)

func startServer(port string, db *sql.DB) {
	handler := storage_handler.NewStorageHandler(db)

	http.HandleFunc("/", handler.Root)
	http.HandleFunc("/api/v1/sign_up", handler.SignUp)
	http.HandleFunc("/api/v1/auth", handler.Auth)
	http.HandleFunc("/api/v1/logout", handler.LogOut)
	fmt.Println("starting server at :" + port)
	http.ListenAndServe(":"+port, nil)
}

/*
func runGetRoot() {
	url := "http://127.0.0.1:8080/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close() // важный пункт!

	respBody, err := ioutil.ReadAll(resp.Body)

	fmt.Printf(string(respBody))
}

func createUser() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	data := `{"username": "Dima", "password": "password"}`
	body := bytes.NewBufferString(data)

	url := "http://127.0.0.1:8080/api/v1/sign_up"
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error happened", err)
		return
	}
	defer resp.Body.Close() // важный пункт!

	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("runTransport %#v\n", string(respBody))
}
*/

func main() {
	err := ini.LoadExists("./config.ini")
	if err != nil {
		fmt.Println("config file is missing")
		os.Exit(1)
	}

	var (
		serverPort = ini.String("server.port")
		host       = ini.String("postgres.host")
		port       = ini.String("postgres.port")
		user       = os.Getenv("POSTGRES_USER")
		password   = os.Getenv("POSTGRES_PASSWORD")
		dbname     = os.Getenv("POSTGRES_DB")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("error while starting database %e", err)
	}
	defer db.Close()

	fmt.Println("we ready to ping")

	err = db.Ping()
	if err != nil {
		log.Fatalf("error database ping %e", err)
	}
	fmt.Println("Successfully connected to database!")

	startServer(serverPort, db)
}
