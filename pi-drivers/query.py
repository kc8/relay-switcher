import requests
import json
from requests import HTTPError
import logger

class Query:

    def __init__(self):
        #self.base = 'http://pi.cooperkyle.com'
        self.base = 'http://localhost:8080'

    def get(self, path: str) -> dict:
        """
        Query an pi with requetst module
        """
        url = self.base + path
        try:
            r = requests.get(url)
            result = r.json()
        except HTTPError as e:
            raise e
        except ValueError as e :
            raise e
        return result

    def post(self, path: str, payload: dict) -> dict:
        """
        Query an pi with requetst module
        """
        url = self.base + path
        try:
            r = requests.post(url, json=payload)
            result = r.json()
        except HTTPError as e:
            raise e
        except ValueError as e :
            raise e
        return result

class MessageHandler:
    def __init__(self, logger: logger.Logger):
        self.query = Query();
        self.logger = logger

    def getRelayStatus(self) -> bool:
        rawResp = self.query.get("/rpiStatus")
        if rawResp['valid'] == True and rawResp['status']== True:
            return True
        return False

    def setRelayStatus(self, status: bool) -> bool:
        payload  = {"msgId": "one", "status":status, "rpiId": "0"}
        rawResp = self.query.post("/setStatus", payload)
        self.logger.log(rawResp)
        if rawResp['MessageId'] != None and rawResp['Received Status'] == status:
            return True
        return False


logger = logger.Logger()
q = MessageHandler(logger)
q.setRelayStatus(True)
q.setRelayStatus(True)
q.setRelayStatus(True)
print(q.getRelayStatus())
q.setRelayStatus(True)
