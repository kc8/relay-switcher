import requests
from requests import HTTPError
import logger
import os

class BadEnvException(Exception): 
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)

class Query:
    def __init__(self):
        PI_RELAY_SERVER_RESOURCE = "PI_RELAY_SERVER_RESOURCE"
        uri = os.environ.get(PI_RELAY_SERVER_RESOURCE)
        if uri is not None:
            self.base = uri
        else: 
            raise BadEnvException(f'Env variable {PI_RELAY_SERVER_RESOURCE} is not defined') 

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
        Query the pi
        """
        url = self.base + path
        try:
            r = requests.post(url, json=payload)
            result = r.json()
        except Exception as e:
            raise e
        return result

class MessageHandler:
    def __init__(self, logger: logger.Logger):
        self.query = Query();
        self.logger = logger

    def getRelayStatus(self) -> dict:
        rawResp = self.query.get("/rpiStatus")
        result = {}
        self.logger.log(rawResp)
        if rawResp['valid'] == True:
            result = {"valid": rawResp['valid'], "status": rawResp['status']}
            return result
        result = {"valid": False}
        return result

    def setRelayStatus(self, status: bool) -> bool:
        payload  = {"msgId": "one", "status":status, "rpiId": "0"}
        rawResp = self.query.post("/setStatus", payload)
        if rawResp['MessageId'] != None and rawResp['Received Status'] == status:
            return True
        return False
