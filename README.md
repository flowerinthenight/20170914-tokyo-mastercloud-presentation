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

Run demoapp in a container using ETW as logging driver (Windows)

* First, run `mftrace` to read real-time logs from ETW. Tool is already provided, along with the config file.

```
# do this in a separate cmd/powershell window
$ cd mftrace\x86\
$ mftrace -c config.xml
```

* Then run our demoapp in a container.

```
$ docker build -t demoapp -f Dockerfile-win .
$ docker run -d --log-driver=etwlogs --name demoapp demoapp:latest
```
