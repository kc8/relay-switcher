try:
    import RPi.GPIO as GPIO
except:
    import Mock.GPIO as GPIO

class pin:
    def __init__(self, pin_num, initial_state):
        self.pin_num = pin_num
        self.state = initial_state
        GPIO.setmode(GPIO.BCM)
        GPIO.setup(pin_num, GPIO.OUT)

    def toggle_state(self) -> None:
        if self.state == False:
            self.off()
        else:
            self.on

    def on(self) -> None:
        self.state = True
        GPIO.output(self.pin_num, 1)

    def off(self) -> None:
        self.state = False
        GPIO.output(self.pin_num, 0)

    def getState(self) -> bool:
        return self.state

    def determineState(self, status: bool) -> None:
        """
        Pass in the state yuo want this should set the state
        """
        if status == False and self.getState() == True:
            self.off()
        elif status == True and self.getState() == False:
            self.on()

    def cleanup(self) -> None:
        GPIO.cleanup()

    def __del__(self) -> None:
        self.cleanup()
