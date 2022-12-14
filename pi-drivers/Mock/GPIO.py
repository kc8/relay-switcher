# NOTE these values do not reflect what they should be
# Also this entire module is crudely done
BCM = 10
OUT = 10

def output(pinNum, state):
    print(f'recevied output with {pinNum} and state {state}')

def setup(pn, gpio_type):
    print(f'recevied setup with pin {pn} and type {gpio_type}')

def setmode(gpio_type):
    print(f'recevied set mode {gpio_type}')

def cleanup(pn=None):
    print("gpio clean up called")

def input(pn) -> int:
    return 1

