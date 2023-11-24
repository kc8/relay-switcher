#! /bin/python3
import query
import logger
import sys
import time
import pin
import os

class BadEnvException(Exception): 
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)

def main(args: list[str]):
    log = logger.Logger()
    log.log("rpi relay driver started")

    PI_RELAY_SERVER_RESOURCE = "PI_RELAY_SERVER_RESOURCE"
    RPI_ID = "RPI_ID"

    serverAddr = os.environ.get(PI_RELAY_SERVER_RESOURCE)
    rpiId = os.environ.get(RPI_ID)

    if serverAddr is None:
        raise BadEnvException(f'Env variable {PI_RELAY_SERVER_RESOURCE} is not set') 
    if rpiId is None:
        raise BadEnvException(f'Env variable {RPI_ID} is not set') 

    print(serverAddr)
    print(rpiId)
    q = query.MessageHandler(rpiId, serverAddr, log) 
    try:
        q.initBackendWithId()
    except Exception as err:
        log.log("Falied to init server")
        log.log(err)

    relayPin = pin.pin(14, False, log)

    while(True):
        try:
            status = q.getRelayStatus()
            if status != None and status['valid'] == True:
                status = status['status']
                relayPin.determineState(status)
            time.sleep(1.0) ## We sleep to not send so many requests
        except Exception as err:
            log.log("received error when getting status")
            log.log(err)

if __name__ == '__main__':
    main(sys.argv)
