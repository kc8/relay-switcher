package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Services struct {
    messageService *MessageService
}

func InitServices() *Services {
    result := &Services{} 
    result.messageService = MessageServiceCreate()
    return result 
}

type Request struct {
    Port string 
}

type Response struct {
    StatusCode int;
    Headers map[string]string
    Body string 
}

type ServerStatus struct {
    StatusMessge string `json:"statusMsg"`
}

func getServerStatus(gc *gin.Context) {
    ss := ServerStatus{"good"}
    gc.IndentedJSON(http.StatusOK, ss)
}

func (s* Services) postSetStatus(gc *gin.Context) {
    var json Message 
    isValid := gc.ShouldBindJSON(&json)
    if isValid != nil {
        gc.JSON(http.StatusBadRequest, gin.H{"Invalid request for settting status": isValid.Error()})
        return
    }
    // TODO some kind of validation here also the UUID would be nice to generate and
    // let the clinet know about the UUID of their request because we want
    // to track requets in the furutre
    var msgId string = uuid.NewString()
    var message Message = Message{msgId, json.RpiId, json.Status, true}
    s.messageService.Set(message)

    gc.JSON(http.StatusOK, gin.H{"Received Status": json.Status, "MessageId": msgId})
}

// NOTE this will consumer the message within the queue
func (s* Services) getNextMessage(gc *gin.Context) {
    var msg Message = s.messageService.Get()
    gc.IndentedJSON(http.StatusOK, msg)
}

func main() {
    router := gin.Default()

    services := new(Services)
    services.messageService = MessageServiceCreate()

    router.GET("/status", getServerStatus)
    router.POST("/setStatus", services.postSetStatus)
    router.GET("/rpiStatus", services.getNextMessage)

    router.Run("localhost:8080")
}
