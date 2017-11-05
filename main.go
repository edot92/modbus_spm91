package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/goburrow/modbus"
)

type structEnergyMeter struct {
	Tegangan  string
	Arus      string
	DayaAktif string
	Daya      string
	Frekuensi string
}

var register [15]uint16
var errModbus error
var dataEnergyMeter structEnergyMeter

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}
func getEnergyMeter(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID":  claims["id"],
		"payload": dataEnergyMeter,
	})
}
func main() {

	// http handler
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			if (userId == "admin" && password == "admin") || (userId == "test" && password == "test") {
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			if userId == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"error":   true,
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/api")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	r.GET("energymeter", getEnergyMeter)
	go http.ListenAndServe(":"+port, r)

	go func() {
		for {
		scanModbus:
			time.Sleep(1 * time.Second)
			// Modbus RTU/ASCII
			handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
			handler.BaudRate = 9600
			handler.DataBits = 8
			handler.Parity = "N"
			handler.StopBits = 1
			handler.SlaveId = 1
			handler.Timeout = 10 * time.Second

			err := handler.Connect()
			if err != nil {
				errModbus = err
				log.Panic(err)
			}
			defer handler.Close()
			client := modbus.NewClient(handler)
			resModbusByte, err := client.ReadHoldingRegisters(0, 20)
			if err != nil {
				errModbus = err
				fmt.Println(err.Error())
				handler.Close()
				goto scanModbus
			}
			dayaAktif := float64(float64(binary.LittleEndian.Uint16(resModbusByte[0:4])) * 0.1)
			tegangan := float64(binary.BigEndian.Uint16(resModbusByte[4:6])) * 0.01
			arus := float64(binary.BigEndian.Uint16(resModbusByte[6:10])) * 0.001
			daya := float64(binary.BigEndian.Uint16(resModbusByte[10:13])) * 0.01
			fmt.Println(resModbusByte)
			// frekuensi := float64(binary.BigEndian.Uint16(resModbusByte[24:25])) * 0.01
			dataEnergyMeter.DayaAktif = fmt.Sprintf("%2f", dayaAktif)
			dataEnergyMeter.Tegangan = fmt.Sprintf("%2f", tegangan)
			dataEnergyMeter.Arus = fmt.Sprintf("%2f", arus)
			dataEnergyMeter.Daya = fmt.Sprintf("%2f", daya)
			// dataEnergyMeter.Frekuensi = fmt.Sprintf("%2f", frekuensi)
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
