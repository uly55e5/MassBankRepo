package server

import "github.com/uly55e5/MassBankRepo/api-server/database"

func Init() {
	database.InitMongoDB()
}

func Close() {

}
