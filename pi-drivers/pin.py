try:
    import RPi.GPIO as GPIO
except:
    import Mock.GPIO as GPIO

class pin:
    def __init__(self, pin_num, initial_state, log):
        self.pin_num = pin_num
        self.state = initial_state
        self.logger = log
        GPIO.cleanup(pin_num)
        GPIO.setmode(GPIO.BCM)
        GPIO.setup(pin_num, GPIO.OUT)
        self.off()

    def _configure_initial_state(self) -> None:
        """
        TODO we want to reset the initial state (or determine what state
        we are on when booting incase the script crashes)
        """
        currentPinState = GPIO.input(self.pin_num)
        if currentPinState == 1:
            self.off()
        elif currentPinState == 0:
            self.on()

    def toggle_state(self) -> None:
        if self.state == False:
            self.off()
        else:
            self.on

    def on(self) -> None:
        self.state = True
        GPIO.output(self.pin_num, 0)

    def off(self) -> None:
        self.state = False
        GPIO.output(self.pin_num, 1)

    def getState(self) -> bool:
        return self.state

    def determineState(self, status: bool) -> None:
        if status == False and self.getState() == True:
            self.off()
        elif status == True and self.getState() == False:
            self.on()

    def cleanup(self) -> None:
        GPIO.cleanup()

    def __del__(self) -> None:
        self.cleanup()
