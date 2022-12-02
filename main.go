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

var client *razorpay.Client

func init() {
	var APIKey = os.Getenv("RZP_KEY_ID")
	var APISecret = os.Getenv("RZP_KEY_SECRET")
	client = razorpay.NewClient(APIKey, APISecret)
}

func main() {
	router := gin.Default()
	router.GET("/order", CreateOrder)
	router.GET("/customer", CreateCustomer)
	router.GET("/payment/:id", FetchPaymentById)
	router.GET("/payments", FetchAllPayments)
	err := router.Run("localhost:8081")
	if err != nil {
		panic(err)
	}

}

func CreateOrder(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	orderParams := map[string]interface{}{
		"amount":   queryParams.Get("amount"),
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	response, err := client.Order.Create(orderParams, nil)
	if err != nil {
		c.AbortWithStatus(500)
	}

	var order Order
	responseBytes, err := json.MarshalIndent(response, "", "    ")
	fmt.Println(string(responseBytes))

	err = json.Unmarshal(responseBytes, &order)
	if err != nil {
		c.AbortWithStatus(500)
	}

	responseModifiers(c)
	c.IndentedJSON(http.StatusOK, order)
}

func CreateCustomer(c *gin.Context) {
	customerParams := map[string]interface{}{
		"name":          "Mehul Sharma",
		"contact":       "8769888888",
		"email":         fmt.Sprintf("mehulsharma%d.me@gmail.com", rand.New(rand.NewSource(int64(time.Now().Second()))).Int()),
		"fail_existing": "0",
		"gstin":         "05BPNPS8989Q2AB",
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

	responseModifiers(c)
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

	responseModifiers(c)
	c.IndentedJSON(http.StatusOK, payment)
}

func FetchAllPayments(c *gin.Context) {
	response, err := client.Payment.All(map[string]interface{}{
		"count": 2,
	}, nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var fetchAllPaymentResponse FetchAllPaymentsResponse
	responseBytes, err := json.MarshalIndent(response, "", "    ")
	fmt.Println(string(responseBytes))

	err = json.Unmarshal(responseBytes, &fetchAllPaymentResponse)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	responseModifiers(c)
	c.IndentedJSON(http.StatusOK, fetchAllPaymentResponse.Items)
}

func responseModifiers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
