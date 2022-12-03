package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// TODO we can create a library for these structs
type Message struct {
    MsgId string `json:"msgId"`
    RpiId string `json:"rpiId"`
    Status bool `json:"status" validate:"required, bool"`
    Valid bool `json:"valid" validate:"required, bool"`
}

func isBetween(val int, min int, max int) bool {
    if (val >= min) && (val <= max) {
        return true
    } else {
        return false
    }
}

func shouldTurnOff() bool {
    loc, _ := time.LoadLocation("America/New_York")
    var onHour = 16
    var offHour = 22
    var currentTime = time.Now().In(loc)
    var currentHour int = currentTime.Hour()
    var currentMinute int = currentTime.Minute()

    fmt.Printf(
        "using hour %d and minutes %d as current and compare is %d, and %d \n", 
        currentHour, 
        currentMinute,
        onHour,
        offHour)

    var betweenHours bool = isBetween(currentHour, onHour, offHour)
    var result = betweenHours

    return result
}

func Main() {
    var status = shouldTurnOff() 
    fmt.Printf("setting status %v\n", status)
    postBody, _ := json.Marshal(Message{
        MsgId: "new",
        RpiId: "0",
        Status: status,
    })
    postBodyBuffer := bytes.NewBuffer(postBody) 
    post , err := http.Post("http://pi.cooperkyle.com/setStatus", "application/json", postBodyBuffer) 
    if err != nil {
        fmt.Printf("Failed to post request to http server  %v\n", err)
    }
    resp, err := ioutil.ReadAll(post.Body)
    fmt.Printf("Response from server was %v\n", string(resp))
}
