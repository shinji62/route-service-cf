##Description
Reverse proxy to be use with the Cloudfoundry route service.

Features:
* User proper go Reverse proxy
* Validate `X-CF-Proxy-Signature`  `X-CF-Forwarded-Url` `X-CF-Proxy-Metadata`

##Creating Route-Service
### Using User provided cf >= 229 and use edge cf cli (unrelease build)

[![asciicast](https://asciinema.org/a/bt7b4ft46en9asjm2govnury7.png)](https://asciinema.org/a/bt7b4ft46en9asjm2govnury7)

##Usage

### Build (tested with go 1.5.3)
```
$godep go build
```

###Just push to cloudfoundry
```
$cf push application-name -b go_buildpack -m 128M
```




## TBD / idea
* Add Signal handler and so on. (listenAndService is too simple)
* Embedded service-broker
* More Configuration
* Worker / pool (offload work)
