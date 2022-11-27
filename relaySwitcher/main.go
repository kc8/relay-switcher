package main 

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Request struct {
    Port string 
}

type Response struct {
    StatusCode int;
    Headers map[string]string
    Body string 
}

type StatusMessage struct {
}

type ServerStatus struct {
    StatusMessge string `json:"statusMsg"`
}

func getServerStatus(gc *gin.Context) {
    ss := ServerStatus{"good"}
    gc.IndentedJSON(http.StatusOK, ss)
}

func main() {
    router := gin.Default()
    router.GET("/status", getServerStatus)
    router.Run("localhost:8080")
}
