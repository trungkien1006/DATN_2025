package main

import (
	"final_project/internal/boostraps"
	"final_project/internal/pkg/helpers"
	"final_project/internal/shared/validator"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// @BasePath        /api/v1
// @title           ShareAndSave - API
// @version         1.0
// @description     Đây là tài liệu Swagger cho hệ thống.
// @termsOfService  http://swagger.io/terms/
// @contact.name    Kin
// @contact.email   trannguyentrungkien1006@gmail.com
// @license.name    Apache 2.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	//Tải các biến môi trường
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Có lỗi khi tải biến môi trường:", err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	validator.InitValidator()

	redisClient := boostraps.InitRedis()

	helpers.Redis = redisClient

	db := boostraps.GormConnection()

	route := boostraps.InitRoute(db, redisClient)

	port := os.Getenv("PORT")

	fmt.Println("Port của bạn: ", port)

	ln, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		panic(err)
	}

	_ = http.Serve(ln, route)
}
