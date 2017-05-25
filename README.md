# log

log package from upspin. ( by the same Golang authors )

## Usage

### Default level

`Info`
> use flag or environment variable to set different default level through init.
> 
> call **log.SetLevel("debug")** to change level.

## Log

```
log.Print...
log.Debug.Print...
log.Error.Print...
```

## Log file setting (size, rotate, etc.)

See http://github.com/natefinch/lumberjack

```
log.SetOutput(&lumberjack.Logger{
    Filename:   "/var/log/myapp/foo.log",
    MaxSize:    500, // megabytes
    MaxBackups: 3,
    MaxAge:     28, //days
})
```


## Other packages considered (but not choose it) 
> (for the record here only, use above one )
* https://github.com/op/go-logging
* https://github.com/Sirupsen/logrus
* https://github.com/hashicorp/logutils
* https://github.com/golang/glog
* https://github.com/go-kit/kit/tree/master/log
* https://github.com/ScottMansfield/nanolog
* https://github.com/sirupsen/logrus
* https://github.com/apex/log
* https://github.com/uber-go/zap
* https://github.com/juju/loggo
