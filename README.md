# Relay Switcher

This is really just used as an off and on switch, hence the relay. 

## A diagram:

**TODO**

The goal here is to subscribe to a message queue so the relay switcher running on the 
raspberry pi. Each command should have some generic strcuture like the following:


```json
{
    "meta": {
        "zone": "UTC",
        "epoch": 12093812,
        "priority": 0
    },
    "data": {
        "desc": "some text",
        "pin": 10, 
        "command":"ON | OFF"
    }
}
```

## The Architecture
This does not use any formal pub/sub or message queue, which would be best so if you are looking for this. This might not 
be the answer for you. 

Instead we store our messages in a firestore database (or other document dat store). A function handler sits there listening for 
messages to add into the queue (or maybe pop off the queue). 

Another function is the function the rapsberry pi reads from, this will get all the messages in the store. Sort them, and 
then take the top one of the 'stack'. This can be determined by the date time

## Message tie breakers
This can be determined by a priority number set in the meta data block. Otherwise we will just execute them both without any gaurentee


## Authentication to the endpoint 
Currently we do not have any as there is not really any sensitive data
