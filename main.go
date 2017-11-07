package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-ini/ini"
	"github.com/goburrow/modbus"
	"github.com/metakeule/fmtdate"
	goSerial "go.bug.st/serial.v1"
	ini "gopkg.in/ini.v1"
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

func getEnergyMeter(c *gin.Context) {
	if errModbus != nil {
		c.JSON(200, gin.H{
			"error": true,
			// "userID":  claims["id"],
			"message": errModbus.Error(),
		})
		return
	}
	// claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"error":   false,
		"message": "data energy",
		// "userID":  claims["id"],
		"payload": dataEnergyMeter,
	})
}
func getAvalilablePORT(c *gin.Context) {

	setPort, err := readKonfigurasiFile()
	list, err := goSerial.GetPortsList()
	if err != nil {
		c.JSON(200, gin.H{
			"error":   true,
			"message": err.Error() + ", waktu:" + fmtdate.Format("hh-mm-ss DD-MM-YYYY", time.Now()),
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": "list port yang tersedia",
		// "userID":  claims["id"],
		"payload": map[string]interface{}{
			"portUsb": list,
			"setPort": setPort,
		},
	})
}
func setPortDariWEB(c *gin.Context) {
	portnya := c.Query("port")
	cfg, err := ini.InsensitiveLoad("setting.ini")
	if err != nil {
		c.JSON(200, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	cfg.Section("").Key("PORT").SetValue(portnya)
	err = cfg.SaveTo("setting.ini")
	if err != nil {
		c.JSON(200, gin.H{
			"error":   true,
			"message": err.Error() + ", waktu:" + fmtdate.Format("hh-mm-ss DD-MM-YYYY", time.Now()),
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": "berhasil update",
	})
}
func readKonfigurasiFile() (string, error) {
	cfg, err := ini.InsensitiveLoad("setting.ini")
	if err != nil {
		return "", err
	}
	key, err := cfg.Section("").GetKey("PORT")
	if err != nil {
		return "", err
	}
	return key.String(), nil
}
func main() {

	// http handler
	gin.SetMode(gin.ReleaseMode)

	portWeb := os.Getenv("PORT")
	r := gin.New()
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	if portWeb == "" {
		portWeb = "9000"
	}

	// jika pakai autentifikasi , sekarang tidak
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
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	/* end autentidfikasi */
	r.GET("/api/energymeter", getEnergyMeter)
	r.GET("/api/getport", getAvalilablePORT)
	r.GET("/api/setport", setPortDariWEB)

	r.Use(static.Serve("/", static.LocalFile("dist", true)))
	r.Use(static.Serve("/static", static.LocalFile("dist/static", true)))

	go http.ListenAndServe(":"+portWeb, r)

	// thread serial modbus run
	go func() {
		for {
		scanModbus:
			time.Sleep(3 * time.Second)
			COMPORT, err := readKonfigurasiFile()
			if err != nil {
				time.Sleep(3 * time.Second)
				errModbus = err
				goto scanModbus
			}
			// Modbus RTU/ASCII
			handler := modbus.NewRTUClientHandler(COMPORT)
			handler.BaudRate = 9600
			handler.DataBits = 8
			handler.Parity = "N"
			handler.StopBits = 1
			handler.SlaveId = 1
			handler.Timeout = 5 * time.Second
			err = handler.Connect()
			if err != nil {
				if strings.Contains(err.Error(), "no such file or directory") {
					err = errors.New("KONEKSI PORT :" + COMPORT + " TIDAK SESUAI , silakan ubah konfigurasi COM PORT")
				} else if strings.Contains(err.Error(), "timeout") {
					handler.Close()
					err = errors.New("Tidak ada respon dari alat , silakan cek konfigurasi COM PORT , dan koneksi ke power meter")
				}
				errModbus = err
				fmt.Println(errModbus)
				time.Sleep(3 * time.Second)
				goto scanModbus
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
			dayaAktif := float64(binary.BigEndian.Uint16(resModbusByte[0:4])) * 0.1
			tegangan := float64(binary.BigEndian.Uint16(resModbusByte[4:6])) * 0.01
			arus := float64(binary.BigEndian.Uint16(resModbusByte[6:10])) * 0.001
			daya := float64(binary.BigEndian.Uint16(resModbusByte[10:13])) * 0.01
			// fmt.Println(resModbusByte)
			// frekuensi := float64(binary.BigEndian.Uint16(resModbusByte[24:25])) * 0.01
			dataEnergyMeter.DayaAktif = fmt.Sprintf("%2f", dayaAktif)
			dataEnergyMeter.Tegangan = fmt.Sprintf("%2f", tegangan)
			dataEnergyMeter.Arus = fmt.Sprintf("%2f", arus)
			dataEnergyMeter.Daya = fmt.Sprintf("%2f", daya)
			// dataEnergyMeter.Frekuensi = fmt.Sprintf("%2f", frekuensi)
		}
	}()

	//* endless loop */
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
