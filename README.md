# log

log package from upspin.

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
