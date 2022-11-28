package main

// NOTE these must be 'exported' with a capital letter
type Message struct {
    MsgId string `json:"msgId"`
    RpiId string `json:"rpiId"`
    Status bool `json:"status" validate:"required, bool"`
    Valid bool `json:"valid" validate:"required, bool"`
}

type MessageService struct {
    messageQueue *Queue
}

func MessageServiceCreate() *MessageService {
    result := &MessageService{}
    result.messageQueue = New()
    return result
}

// Get the latest message in the queue
func (ms* MessageService) Get() Message {
    item := ms.messageQueue.Deque()
    msg, _ := item.(Message)
    return msg
}

// Get the latest message in the queue
func (ms* MessageService) Set(m Message) {
    ms.messageQueue.Enqueue(m)
}
