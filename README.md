### demoapp

Run demoapp as native, out to stdout.

```
$ cd demoapp
$ go build -v
$ ./demoapp
CTRL+C
```

Run demoapp in a container, view logs via docker logs.

```
$ docker build -t demoapp .
$ docker run -d --name demoapp demoapp:latest
$ docker logs -f demoapp
CTRL+C
$ docker rm -f demoapp
```

Run demoapp in a container using syslog as logging driver.

```
$ docker run -d --log-driver=syslog --log-opt tag=demoapp --name demoapp demoapp:latest
$ tail -f /var/log/syslog | grep -i 'demoapp'
```
