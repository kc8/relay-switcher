# Systemd 

```sh
pi@raspberrypi:~/relay-switcher/pi-drivers $ sudo cp ../scripts/xmas.service /etc/systemd/system/
pi@raspberrypi:~/relay-switcher/pi-drivers $ sudo chmod 664 
pi@raspberrypi:~/relay-switcher/pi-drivers $ sudo chmod 664 /etc/systemd/system/xmas.service
pi@raspberrypi:~/relay-switcher/pi-drivers $ sudo systemctl daemon-reload
pi@raspberrypi:~/relay-switcher/pi-drivers $ sudo systemctl enable xmas.service 

```
