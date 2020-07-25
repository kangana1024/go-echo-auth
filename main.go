package main

import (
	"net/url"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	dbmodels "github.com/kangana1024/go-echo-auth/dbmodels"
)

func init() {

}
func main() {
	err := godotenv.Load()
	dsn := url.URL{
		User:     url.UserPassword("kandev", ""),
		Scheme:   "postgres",
		Host:     "127.0.0.1:5432",
		Path:     "usertest",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	// fmt.Println("conected")

	// err = router.Run(":8081")
	// if err != nil {
	// 	log.Fatalln(fmt.Errorf("faild to start Gin application: %w", err))
	// }
	db.AutoMigrate(&dbmodels.CasbinRole{}, &dbmodels.Users{})

}
