package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for {
		checking()
	}
}

func checking() {
	fHandle, err := os.Open("url")
	if err != nil {
		fmt.Println(err)
	}
	defer fHandle.Close()
	fScanner := bufio.NewScanner(fHandle)
	for fScanner.Scan() {
		url := fScanner.Text()
		t0 := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != 200 && resp.StatusCode != 301 && resp.StatusCode != 302 && resp.StatusCode != 000 && resp.StatusCode != 405 {
			token := os.Getenv("TOKEN")
			uid := os.Getenv("USERID")
			webhook := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s", token, uid)
			client := &http.Client{}
			t1 := time.Now()
			duration := t1.Sub(t0).String()
			values := map[string]string{"text": resp.Status + "\n" + url + "\n" + "Duration time: " + duration}

			jsonStr, err := json.Marshal(values)
			if err != nil {
				fmt.Println(err)
			}
			req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(jsonStr))
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Set("Content-Type", "application/json")
			client.Do(req)
		}
		sec := os.Getenv("INTERVAL")
		i, err := strconv.Atoi(sec)
		time.Sleep(time.Second * time.Duration(i))
	}
}
