import requests
from requests import HTTPError
import logger

class Query:
    def __init__(self, serverAddr):
        self.base = serverAddr

    def get(self, path: str, params: dict) -> dict:
        """
        Query an pi with requetst module
        """
        url = self.base + path
        try:
            r = requests.get(url, params=params)
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
    def __init__(self, rpiId, serverAddr, logger: logger.Logger):
        self.query = Query(serverAddr);
        self.logger = logger
        self.rpiId = rpiId

    def getRelayStatus(self) -> dict:
        params = {'rpiId': self.rpiId}
        rawResp = self.query.get("/rpiStatus", params)
        result = {}
        self.logger.log(rawResp)
        if rawResp['valid'] == True:
            result = {"valid": rawResp['valid'], "status": rawResp['status']}
            return result
        result = {"valid": False}
        return result

    def setRelayStatus(self, status: bool) -> bool:
        payload  = {"msgId": "one", "status": status, "rpiId": self.rpiId}
        rawResp = self.query.post("/setStatus", payload)
        if rawResp['MessageId'] != None and rawResp['Received Status'] == status:
            return True
        return False
