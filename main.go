package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var APIKey = os.Getenv("RZP_API_KEY")
var APISecret = os.Getenv("RZP_API_SECRET")

var client *razorpay.Client

func init() {
	client = razorpay.NewClient(APIKey, APISecret)
}

func main() {
	router := gin.Default()
	router.GET("/order", CreateOrder)
	router.GET("/customer", CreateCustomer)
	router.GET("/payment/:id", FetchPaymentById)
	err := router.Run("localhost:8081")
	if err != nil {
		panic(err)
	}

}

func CreateOrder(ctx *gin.Context) {
	orderParams := map[string]interface{}{
		"amount":   5000,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	response, err := client.Order.Create(orderParams, nil)
	if err != nil {
		ctx.AbortWithStatus(500)
	}

	var order Order
	responseBytes, err := json.MarshalIndent(response, "", "    ")
	fmt.Println(string(responseBytes))

	err = json.Unmarshal(responseBytes, &order)
	if err != nil {
		ctx.AbortWithStatus(500)
	}

	ctx.IndentedJSON(http.StatusOK, order)
}

func CreateCustomer(c *gin.Context) {
	customerParams := map[string]interface{}{
		"name":          "Mehul Sharma",
		"contact":       "8769883659",
		"email":         fmt.Sprintf("mehulsharma%d.me@gmail.com", rand.New(rand.NewSource(int64(time.Now().Second()))).Int()),
		"fail_existing": "0",
		"gstin":         "05BPNPS8985R1ZD",
	}
	response, err := client.Customer.Create(customerParams, nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var customer Customer
	responseBytes, err := json.MarshalIndent(response, "", "    ")
	fmt.Println(string(responseBytes))

	err = json.Unmarshal(responseBytes, &customer)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, customer)
}

func FetchPaymentById(c *gin.Context) {
	paymentId := c.Param("id")

	response, err := client.Payment.Fetch(paymentId, nil, nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var payment Payment
	responseBytes, err := json.MarshalIndent(response, "", "    ")
	fmt.Println(string(responseBytes))

	err = json.Unmarshal(responseBytes, &payment)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.IndentedJSON(http.StatusOK, payment)
}
