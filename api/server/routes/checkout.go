package routes

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/markthub/apis/api/server/models"
	"github.com/markthub/apis/api/server/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spaghettifunk/go-mollie/mollie"
)

func AuthorizePayment(m mollie.Client, o *models.Order, orderNumber, baseUrl string) (string, string, error) {

	total := strconv.FormatFloat(o.Amount+o.Tip+o.Fee, 'f', 2, 64)

	pr := mollie.PaymentRequest{
		Amount: mollie.Amount{
			Currency: "EUR",
			Value:    total,
		},
		Method:      o.PaymentMethod,
		Description: "Payment of Order " + orderNumber,
		RedirectURL: baseUrl + "/redirect?order_id=" + orderNumber,
		WebhookURL:  baseUrl + "/webhook",
		Metadata: map[string]string{
			"orderId": orderNumber,
		},
		Issuer: o.Issuer,
	}

	resp, err := m.CreatePayment(&pr)
	if err != nil {
		return "", "", err
	}

	// TODO: Hacky! Need to be smarter here
	l := resp.Links.(map[string]interface{})

	type links struct {
		Self          map[string]string
		Checkout      map[string]string
		Documentation map[string]string
		Type          map[string]string
	}

	var result links
	err = mapstructure.Decode(l, &result)
	if err != nil {
		return "", "", err
	}

	return resp.Status, result.Checkout["href"], nil
}

// PaymentWebhook is called when checking the status of the payment
func PaymentWebhook(c *gin.Context) {

	b, _ := ioutil.ReadAll(c.Request.Body)
	body := strings.Split(string(b), "=")

	paymentID := string(body[1])

	db := c.MustGet("DB").(*gorm.DB)
	m := c.MustGet("MollieClient").(mollie.Client)

	p, err := m.GetPayment(paymentID, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	// TODO: Hacky! Need to be smarter here
	meta := p.Metadata.(map[string]interface{})
	type metadata struct {
		OrderId string
	}

	var result metadata
	err = mapstructure.Decode(meta, &result)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var order models.Order
	if err := db.Where("order_number = ?", result.OrderId).First(&order).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	order.StatusCode = models.StatusCode(strings.ToUpper(p.Status))
	if err := db.Save(order).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	switch strings.ToUpper(p.Status) {
	case string(models.Open):
		// TODO: Not sure what to do here
		c.JSON(http.StatusOK, "Open")
		return
	case string(models.Failed):
		// TODO: start back from Orders
		c.JSON(http.StatusOK, "Failed")
		return
	case string(models.Canceled):
		// TODO: start back from Orders
		c.JSON(http.StatusOK, "Canceled")
		return
	case string(models.Paid):
		// TODO: Display the Paid information in the frontend
		c.JSON(http.StatusOK, "Paid")
		return
	case string(models.Pending):
		// TODO: Not sure what to do here
		c.JSON(http.StatusOK, "Pending")
		return
	case string(models.Expired):
		// TODO: trigger an Expired page in the frontend
		c.JSON(http.StatusOK, "Expired")
		return
	default:
		// TODO: start back from Orders
		c.JSON(http.StatusOK, "Default")
		return
	}
}

// PaymentRedirectURL is called when the checkout is completed
func PaymentRedirectURL(c *gin.Context) {
	orderNumber := c.Query("order_id")
	// db := c.MustGet("DB").(*gorm.DB)
	// nc := c.MustGet("NSQClient").(*nsq.NSQClient)
	// mb := c.MustGet("MapBoxClient").(*mapbox.Mapbox)

	msg := map[string]string{
		"message":     "Thank you for your order!",
		"orderNumber": orderNumber,
	}

	// Go routine to assign the driver here
	// go AssignOrderToDriver(orderNumber, db, nc, mb)

	utils.Response(c, http.StatusOK, msg)
}
