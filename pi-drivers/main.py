#! /bin/python3
import query
import logger
import sys
import time
import pin

def main(args: list[str]):
    l = logger.Logger()
    l.log("rpi relay driver started")
    q = query.MessageHandler(l)
    relayPin = pin.pin(14, False)

    while(True):
        try:
            status = q.getRelayStatus()
            if status != None and status['valid'] == True:
                status = status['status']
                relayPin.determineState(status)
            time.sleep(1.0)
        except Exception as err:
            l.log("received error when getting status")
            l.log(err)

if __name__ == '__main__':
    main(sys.argv)
