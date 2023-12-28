Mysql mysql.general_log table monitor

mysql variables
```SQL
SET GLOBAL general_log = 'ON';
SET GLOBAL log_output = 'TABLE';
```


```bash
go mod tidy
cd cmd
go build

./cmd -addr=127.0.0.1 -port=20701 -user=root -password=Root12#$ -eventuser=gorm
```

  -addr string
        mysql ip (default "127.0.0.1")
  -eventuser string
        monitoring user
  -password string
        mysql password
  -port string
        mysql port (default "3306")
  -user string
        mysql user (default "root")


$ ./cmd -addr=127.0.0.1 -port=20701 -user=root -password=Root12#$ -eventuser=gorm
addr: 127.0.0.1
port: 20701
user: root
password: Root12#$
eventuser: gorm
tail: []
root:Root12#$@tcp(127.0.0.1:20701)/mysql?charset=utf8mb4&parseTime=True&loc=Local
[2023-12-28 11:22:31.233178] [Query] SELECT 1
[2023-12-28 11:22:31.246793] [Query] SELECT st.* FROM performance_schema.events_statements_current st JOIN performance_schema.threads thr ON thr.thread_id = st.thread_id WHERE thr.processlist_id = 64
[2023-12-28 11:22:31.250843] [Query] SELECT st.* FROM performance_schema.events_stages_history_long st WHERE st.nesting_event_id = 87
[2023-12-28 11:22:31.252954] [Query] SELECT st.* FROM performance_schema.events_waits_history_long st WHERE st.nesting_event_id = 87
[2023-12-28 11:22:32.268787] [Query] SELECT 1
[2023-12-28 11:22:32.288996] [Query] SELECT st.* FROM performance_schema.events_statements_current st JOIN performance_schema.threads thr ON thr.thread_id = st.thread_id WHERE thr.processlist_id = 64
[2023-12-28 11:22:32.291124] [Query] SELECT st.* FROM performance_schema.events_stages_history_long st WHERE st.nesting_event_id = 89
[2023-12-28 11:22:32.293228] [Query] SELECT st.* FROM performance_schema.events_waits_history_long st WHERE st.nesting_event_id = 89
