# hightlogs
Universal Logs

Logs can be saved in MongoDB with Levels (ERROR, EVENT, DEBUG, etc..)
and sender identifier (in case goroutine > 1)


### Usage

Call function InitLogsDB once in case mongoDB used.

```bash
    import (
        "gopkg.in/mgo.v2"
        "github.com/BankEx/hightlogs/hlog"
    )

    MongoSession, err := mgo.Dial('connection_url')
	MongoDB = MongoSession.DB('database_name')

    hlog.InitLogsDB(MongoDB)
```

Call Log function.

1st param - sender ID (goroutine ID, function name, .go file name, etc)

2st param - log level (FATAL lvl will stop the application)

3, 4, 5, etc params - logging data


```bash
    import "github.com/BankEx/hightlogs/hlog"

    hlog.SaveLog("main", hlog.EVENT, "Hello", "World", 123)
```

### Log levels

const FATAL = 0

const ERROR = 1

const MAIN = 2

const EVENT = 3

const DETAILS = 4

const DEBUG = 5

const MINOR = 6

