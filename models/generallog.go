package models

import "time"

type Generallog struct {
	EventTime   time.Time `gorm:"column:event_time;type:timestamp(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);NOT NULL;"` // ` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
	UserHost    string    `gorm:"column:user_host;type:mediumtext;NOT NULL;"`                                                                // ` mediumtext NOT NULL,
	ThreadId    uint64    `gorm:"column:thread_id;type:bigint unsigned;NOT NULL;"`                                                           // ` bigint unsigned NOT NULL,
	ServerId    uint32    `gorm:"column:server_id;type:int unsigned;NOT NULL;"`                                                              // ` int unsigned NOT NULL,
	CommandType string    `gorm:"column:command_type;type:varchar(64);NOT NULL;"`                                                            // ` varchar(64) NOT NULL,
	Argument    string    `gorm:"column:argument2;type:CHAR(4096);NOT NULL;"`                                                                // ` mediumblob NOT NULL
}

func (Generallog) TableName() string {
	return "general_log"
}
