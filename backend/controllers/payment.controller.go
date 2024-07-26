package controllers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	dbconn "github.com/wiratR/rds_server/database"
	"github.com/wiratR/rds_server/models"
	KsherGO "github.com/wiratR/rds_server/utils/KsherGo"
)

const appId = "mch38806"

// NativePay func payment connect to Ksher
// @Description payment connect to Ksher
// @Summary payment connect to Ksher
// @Tags Payment
// @Accept json
// @Produce json
// @Success 200 {array} models.ResponseBody "Ok"
// @Failure 400 {object} models.ResponseError "Error"
// @Router /payment/nativepay [post]
func NativePay(c *fiber.Ctx) error {

	if dbconn.DB == nil {
		return errors.New("database connection is nil")
	}

	var isSuccess bool = false

	// privateKey, err := utils.ReadFileKey()
	// if err != nil {
	// 	fmt.Println("get key Error:", err)
	// 	return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	// }
	// // fmt.Println("get key :", privateKey)
	// // // Initial KsherGO with appId,privateKey,publickey inside sdk

	privateKey := KsherGO.ReadPrivateKeyFromPath("./config/key/Mch38806_PrivateKey.pem")

	fmt.Println(privateKey)
	client := KsherGO.New(appId, privateKey)
	nowStr := time.Now().Format("20060102150405.000")
	//fmt.Println(nowStr)
	// mchOrderNo, feeType, channel string, totalFee int
	response, err := client.NativePay(strings.Replace(nowStr, ".", "", -1), "THB", "promptpay", 100)
	//response, err := client.QuickPay(strings.Replace(nowStr, ".", "", -1), "THB", "12345", "wechat", "", 100)
	//response, err := client.GatewayOrderQuery("999668")
	if err != nil {
		fmt.Println("QuickPay error:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(models.ApiResponse(isSuccess, fiber.Map{"message": err.Error()}))
	}

	fmt.Println("QuickPay success:", response)
	isSuccess = true
	return c.Status(fiber.StatusOK).JSON(models.ApiResponse(isSuccess, fiber.Map{"data": "ok"}))
}

// func main() {
// 	client := KsherGO.New(appId, privateKey)
// 	//s := "a我cd"
// 	//ss := string([]rune(s)[1:])
// 	//fmt.Println(ss)
// 	//
// 	//nowStr := time.Now().Format("20060102150405.000")
// 	//fmt.Println(nowStr)
// 	//response, err := client.QuickPay(strings.Replace(nowStr, ".", "", -1 ), "THB", "12345", "wechat","", 100)
// 	//response, err := client.GatewayPay("999668", "THB", "wechat,alipay,airpay", "2233", "https://www.baidu.com/",
// 	//	"https://www.baidu.com/", "test", "https://www.baidu.com/", "PC", 100)
// 	response, err := client.GatewayOrderQuery("999668")
// 	if err != nil {
// 		fmt.Println("QuickPay error:", err.Error())
// 	} else {
// 		fmt.Println("QuickPay success:", response)
// 	}
// }
