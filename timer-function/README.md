# Timer Function 

This function executes on a Digital Ocean function at X and Y time, which is a 
trigger set through the web gui

It looks if the current time is between the hard coded value and 
triggers the relay to switch on or off, this part is just a REST requst to the 
relaySwitcher queue.

## To Deploy to Digital Ocean 

- Make sure that the go file has a `Main()` function

```sh 
cd ./src
doctl serverless deploy ./ --remote-build
```

If you need to auth again, get a new token 
```sh
doctl init auth -t <new api token>
```

You will also need to be in the correct namespace for function suauge
