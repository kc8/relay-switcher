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
        self.off()

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
            print("OFF")
        elif status == True and self.getState() == False:
            print("ON")
            self.on()

    def cleanup(self) -> None:
        GPIO.cleanup()

    def __del__(self) -> None:
        self.cleanup()
