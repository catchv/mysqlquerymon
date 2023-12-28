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
