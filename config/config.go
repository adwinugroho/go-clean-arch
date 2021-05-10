package config

import "os"

var (
	// DB URL
	DBURL string = os.Getenv("DBURL")
	// DB username
	DBUsername string = os.Getenv("DBUSERNAME")
	// DB password
	DBPassword string = os.Getenv("DBPASSWORD")
	// DB Name
	DBName string = os.Getenv("DBNAME")
	// DB Log Name
	DBLogName string = os.Getenv("DBLOGNAME")
	// DB Log URL
	DBLOGURL string = os.Getenv("DBLOGURL")
)
