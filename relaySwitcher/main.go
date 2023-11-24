package main

import (
	"fmt"
	"net/http"
    "html/template"
    "bytes"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Services struct {
	piMessageSerivces map[string]RpiConfigContainer
}

func InitServices() *Services {
	result := &Services{}
	result.piMessageSerivces = make(map[string]RpiConfigContainer)
	return result
}

type Request struct {
	Port string
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

type ServerStatus struct {
	Health string `json:"statusMsg"`
    NumInQueue int `json:"numInQueue"`
    NumPisConnected int `json:"numPisConnected"`
}

// TODO This functiokn function should return stats 
func (s *Services) getServerStatus(gc *gin.Context) {
	ss := ServerStatus{
        Health: "TBD",
        NumPisConnected: len(s.piMessageSerivces),
        NumInQueue: 0,
    }
	gc.IndentedJSON(http.StatusOK, ss)
}

const MAX_QUEUE_SIZE = 32 // 32 rpis can connect
func (s *Services) initRpiMessageService(gc *gin.Context) {
	var rpiId = gc.Query("rpiId")
	_, hasKey := s.piMessageSerivces[rpiId]
	if hasKey == true {
		resp := gin.H{"{Failure_Reason": fmt.Sprintf("RpiId Already exists %s}", rpiId)}
		gc.JSON(http.StatusBadRequest, resp)
	}
	s.piMessageSerivces[rpiId] = RpiConfigContainer{
		RpiId:          rpiId,
		MessageService: MessageServiceCreate(),
	}
}

func (s *Services) postSetStatus(gc *gin.Context) {
	var json Message
	isValid := gc.ShouldBindJSON(&json)
	if isValid != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"Invalid request for settting status": isValid.Error()})
		return
	}
	var msgId string = uuid.NewString()

	var rpiId string = json.RpiId
	validRpiId, hasKey := s.piMessageSerivces[rpiId]
	if hasKey == false {
		gc.JSON(http.StatusBadRequest, gin.H{fmt.Sprintf("Invalid or unknown: %s RpId. You /initRpi to create a queue", rpiId): isValid.Error()})
	}

	var message Message = Message{msgId, json.RpiId, json.Status, true}
	validRpiId.MessageService.Set(message)
	gc.JSON(http.StatusOK, gin.H{"Received Status": json.Status, "MessageId": msgId})
}

func (s *Services) getNextMessage(gc *gin.Context) {
	var rpiId = gc.Query("rpiId")
	validRpiId, hasKey := s.piMessageSerivces[rpiId]
	if hasKey == false {
		resp := gin.H{"{Failure_Reason": fmt.Sprintf("Invalid or unknow RpiId %s}", rpiId)}
		gc.JSON(http.StatusBadRequest, resp)
	}
	var msg Message = validRpiId.MessageService.Get()
	gc.IndentedJSON(http.StatusOK, msg)
}

var endpoints = []string{
	"/status", 
	"/setStatus",
	"/initRpi", 
	"/rpiStatus",
};

func (s* Services) landing(gc *gin.Context) {
    const indexTmpl string = `
    <html>
        <head>
        <title>Raspi Queue Controller</title>
        </head>
        <body>
            <h1>Raspi Queue Controller</h1>
            A list of all endpoints 
            {{range .}}
                {{.}}
            {{end}}
        </body>
    </html>
    `

    const errorPage = "<html><body>500 - Internal Server Error</body></html>"
    tmpl, err := template.New("template").Parse(indexTmpl)
    if err != nil {
        gc.Data(http.StatusOK, "text/html charest=utf-8", []byte(errorPage))
    }
    var data bytes.Buffer
    tmpl.Execute(&data, endpoints)
    gc.Data(http.StatusOK, "text/html charest=utf-8", data.Bytes())
}

func main() {
	router := gin.Default()

	services := new(Services)
	router.GET("/status", services.getServerStatus)
	router.POST("/setStatus", services.postSetStatus)
	router.GET("/initRpi", services.initRpiMessageService)
	router.GET("/rpiStatus", services.getNextMessage)

    // TODO a custom swagger like doc, but much more boring and simple
	router.GET("/", services.landing)

	router.Run("localhost:8080")
}
