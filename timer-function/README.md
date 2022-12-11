# Timer Function 

This function executes on a Digital Ocean function. This submits a message/job to the 
relaySwitcher queue in which the respective Raspberry Pi will read. 


## To Deploy to Digital Ocean 

- Make sure that the go file has a `Main()` function with a capital 'm'

```sh 
cd ./src
doctl serverless deploy ./ --remote-build
```

If you need to auth again, get a new token 
```sh
doctl init auth -t <new api token>
```

You will also need to be in the correct digital ocean namespace which can be set 
with the doctl cli
