package main

import (
	"flag"
	"fmt"
	"mysqlquerymon/models"
	"mysqlquerymon/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ./cmd -addr=127.0.0.1 -port=20701 -user=root -password=Root12#$ -eventuser=gorm

func main() {
	addrPtr := flag.String("addr", "127.0.0.1", "mysql ip")
	portPtr := flag.String("port", "3306", "mysql port")
	userPtr := flag.String("user", "root", "mysql user")
	passwordPtr := flag.String("password", "", "mysql password")
	eventuserPtr := flag.String("eventuser", "", "monitoring user")

	flag.Parse()

	fmt.Println("addr:", *addrPtr)
	fmt.Println("port:", *portPtr)
	fmt.Println("user:", *userPtr)
	fmt.Println("password:", *passwordPtr)
	fmt.Println("eventuser:", *eventuserPtr)
	fmt.Println("tail:", flag.Args())

	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local", *userPtr, *passwordPtr, *addrPtr, *portPtr)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err)
	}

	utils.ColorPrintln(utils.ColorRed, dsn)

	var currentTime time.Time = time.Now()

	for {
		var logs = []models.Generallog{}

		db.Where("event_time > ? AND user_host LIKE ?", currentTime, *eventuserPtr+"%").Order("event_time asc").Select(
			"event_time",
			"user_host",
			"thread_id",
			"server_id",
			"command_type",
			"cast(argument as CHAR(4096)) as argument2",
		).Find(&logs)

		for i := 0; i < len(logs); i++ {
			fmt.Printf("\u001b[32m[%10s] [%s] \u001b[35m%s\u001b[0m\n",
				logs[i].EventTime.Format("2006-01-02 15:04:05.999999"),
				logs[i].CommandType,
				logs[i].Argument,
			)
		}

		if len(logs) > 0 {
			// utils.ColorPrintln(utils.ColorGreen,
			// 	fmt.Sprintf("%+v %+v", len(logs), logs[len(logs)-1].EventTime))
			currentTime = logs[len(logs)-1].EventTime
		}

		time.Sleep(1 * time.Second)
	}
}
