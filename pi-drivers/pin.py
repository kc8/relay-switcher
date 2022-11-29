import RPi.GPIO as GPIO
import time

class pin:
    def __init__(self, pin_num, initial_state):
        self.pin_num = pin_num
        self.state = initial_state
        ## TODO setup should be more dynamic
        GPIO.setup(pin_num, GPIO.OUT)

    def toggle_state(self):
        if self.state == False:
            GPIO.output(self.pin_num, 0)
        else:
            GPIO.output(self.pin_num, 1)

    def cleanup(self):
        GPIO.cleanup()

    def __del__(self):
        self.cleanup()

def set_gpio_state():
    GPIO.setmode(GPIO.BCM)

set_gpio_state()
output_pin = pin(14, False)
output_pin.toggle_state()
time.sleep(5)
output_pin.toggle_state()
