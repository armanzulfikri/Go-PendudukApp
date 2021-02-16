package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// var userDatabase, passDatabase, portDatabase,
	// 	hostDatabase, nameDatabase, sslDatabase, timezoneDatabase string

	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Error Load .env file")
	// } else {
	// 	userDatabase = os.Getenv("DATABASE_USER")
	// 	passDatabase = os.Getenv("DATABASE_PASS")
	// 	portDatabase = os.Getenv("DATABASE_PORT")
	// 	hostDatabase = os.Getenv("DATABASE_HOST")
	// 	nameDatabase = os.Getenv("DATABASE_NAME")
	// 	sslDatabase = os.Getenv("DATABASE_SSLMODE")
	// 	timezoneDatabase = os.Getenv("DATABASE_TIMEZONE")

	// }
	// conn :=
	// 	" host=" + hostDatabase +
	// 		" user=" + userDatabase +
	// 		" password=" + passDatabase +
	// 		" dbname=" + nameDatabase +
	// 		" port=" + portDatabase +
	// 		" sslmode=" + sslDatabase +
	// 		" Timezone=" + timezoneDatabase

	db, errConn := gorm.Open(postgres.Open("postgres://xkloddfddainoq:5070b9f51e9b823e71bcf485efba8735fa374494263e6bd1b9f0207de95dc50e@ec2-3-213-85-90.compute-1.amazonaws.com:5432/d3457tu1vpejgm"), &gorm.Config{})
	if errConn != nil {
		panic("failed to connect to database")
	} else {
		fmt.Println("Koneksi Sukses")
	}
	return db
}
