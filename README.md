# Relay Switcher
This is designed to turn on and off the outside (and maybe other Christmas lights) that 
are wired to a raspberry pi

As of 2023, this is now a self hosted solution. 

# pi-drivers
This is the driver that directly control GPIO ports on the raspberry pi turnning them on and off

The python program also queries the RelaySwticher back end to 1. update 
the current status of the lights it controls 2. turn on and off the lights it 
controls

# RelaySwitcher 
This directory holds the backend server written in Go that updates and 
shutdowns the pis. It a message queue that can maintain and update the state of multiple
raspberry pis. It has really only worked with one thus far

## Notes for future me

The goal here is to subscribe to a message queue so the relay switcher running on the 
raspberry pi. Each command should have some generic structure like the following:


```json
{
    "meta": {
        "zone": "UTC",
        "epoch": 12093812,
        "priority": 0
    },
    "data": {
        "desc": "some text",
        "pin": 10, 
        "command":"ON | OFF"
    }
}
```

## Message tie breakers
This can be determined by a priority number set in the meta data block. Otherwise we will just execute them both without any gaurentee

# Deploying 
Everything is deploying via ssh and ansible playbook. Check the /scripts directory for a 
systemd service script as well an ansible deploy script for the back end

You will need ansible playbook installed and a debian based server installed that you can use systemd on to 
start the service

## Relayswitcher
1. Setup a file in `scripts/ansible_settings.ini` containing the host and other info for ansible to 
get access to the back end server  example:
    ```ini
    [hosts]
    server ansible_host=unkubbedservices.cooperkyle.com
    server ansible_ssh_user=sslisbetter
    server ansible_ssh_pass=sslisbetter
    server ansible_sudo_pass=sslisbetter
    ```
1. Run the make file from the relaySwitcher directory with `make deploy`

