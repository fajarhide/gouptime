package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var ok bool = true

func loop() {
	for {
		if ok == true {
			checking()
			time.Sleep(time.Second * 5)
		}
	}
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "{{TOKEN}}",
		Poller: &tb.LongPoller{Timeout: 5 * time.Hour},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		// fmt.Println("Start")
		ok = true
		loop()
		b.Send(m.Sender, "Start")
	})

	b.Handle("/stop", func(m *tb.Message) {
		// fmt.Println("Stop")
		ok = false
		b.Send(m.Sender, "Stop")
	})
	b.Start()

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
			request_url := "https://api.telegram.org/bot{{TOKEN}}/sendMessage?chat_id={{USERID}}"
			client := &http.Client{}
			t1 := time.Now()
			duration := t1.Sub(t0).String()
			values := map[string]string{"text": resp.Status + "\n" + url + "\n" + "Duration time: " + duration}

			jsonStr, err := json.Marshal(values)
			if err != nil {
				fmt.Println(err)
			}
			req, err := http.NewRequest("POST", request_url, bytes.NewBuffer(jsonStr))
			if err != nil {
				fmt.Println(err)
			}
			req.Header.Set("Content-Type", "application/json")
			client.Do(req)
		}
		time.Sleep(time.Second * 2)
	}
}
