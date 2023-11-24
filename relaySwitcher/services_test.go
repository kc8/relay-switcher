package main

import (
	"testing"
)

const rpiId string = "rpiId"

func TestProperInitSerices(t *testing.T) {
	s := InitServices()
    if s.piMessageSerivces == nil {
        t.Errorf("pi Messages is null")
    }
	s.piMessageSerivces[rpiId] = RpiConfigContainer{
		RpiId:          rpiId,
		MessageService: MessageServiceCreate(),
	}

	validRpiId, hasKey := s.piMessageSerivces[rpiId]

    if hasKey == false {
        t.Errorf("has key was false when it should have key")
    }
    if validRpiId.RpiId != rpiId {
        t.Errorf("Failed get set/get correct rpiid")
    }
	var msg Message = validRpiId.MessageService.Get()
    if msg.MsgId != "" {
        t.Errorf("Got a valid msg from qeue but msg id should not be set")
    }
}

func TestServiceTest(t *testing.T) {
	s := InitServices()
	_, hasKey := s.piMessageSerivces[rpiId]
    if hasKey == true {
        t.Errorf("has key was true when it should not have been")
    }
}

func TestProperAddValues(t *testing.T) {
	s := InitServices()

	s.piMessageSerivces[rpiId] = RpiConfigContainer{
		RpiId:          rpiId,
		MessageService: MessageServiceCreate(),
	}
	_, hasKey := s.piMessageSerivces[rpiId]
    if hasKey == false {
        t.Errorf("Should have key")
    }
}
