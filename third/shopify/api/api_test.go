package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	goshopify "github.com/bold-commerce/go-shopify/v4"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	app     = goshopify.App{}
	shop    = "test20240319.myshopify.com"
	token   = os.Getenv("SHOPIFY_APP_ACCESS_TOKEN")
	version = "2024-04"

	ctx = context.Background()

	sc *goshopify.Client
)

func initClient() {
	var err error
	client, err := goshopify.NewClient(app, shop, token, goshopify.WithVersion(version))
	if err != nil {
		panic(err)
	}
	sc = client
}

func TestAccessScope(t *testing.T) {
	initClient()

	// // wrong
	// scopes, err := sc.AccessScopes.List(ctx, nil)
	// if err != nil {
	// 	panic(err) // Not Found
	// }
	// fmt.Println(toJsonStr(scopes, true))

	type Result struct {
		AccessScopes []*goshopify.AccessScope `json:"access_scopes"`
	}
	url := fmt.Sprintf("https://%s/admin/oauth/access_scopes.json", shop)
	var result Result
	resp, err := resty.New().R().
		SetHeader("X-Shopify-Access-Token", token).
		SetResult(&result).
		Get(url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		log.Fatalf("request err: %v, %d", err, resp.StatusCode())
	}
	fmt.Println(len(result.AccessScopes), toJsonStr(result.AccessScopes, true))
}

func TestProduct(t *testing.T) {
	initClient()

	var total int

	// list products (default limit 50)
	products, err1 := sc.Product.List(ctx, nil)
	if err1 != nil {
		log.Fatal(err1)
	}
	total = len(products)
	fmt.Println(total, toJsonStr(products, false))

	// list all products page by page
	opts := goshopify.ProductListOptions{
		ListOptions: goshopify.ListOptions{
			Limit: 10,
		},
	}
	list, page, err := sc.Product.ListWithPagination(ctx, opts)
	if err != nil {
		panic(err)
	}
	for page.NextPageOptions != nil {
		opts = goshopify.ProductListOptions{
			ListOptions: goshopify.ListOptions{
				Limit:    10,
				PageInfo: page.NextPageOptions.PageInfo,
			},
		}
		var next []goshopify.Product
		next, page, err = sc.Product.ListWithPagination(ctx, opts)
		if err != nil {
			panic(err)
		}
		list = append(list, next...)
	}
	assert.Equal(t, total, len(list))

	// a single product
	if len(products) > 0 {
		id := products[0].Id
		product, err2 := sc.Product.Get(ctx, id, nil)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(toJsonStr(product, true))
	}

	// a count of products
	count, err3 := sc.Product.Count(ctx, nil)
	if err3 != nil {
		log.Fatal(err3)
	}
	assert.Equal(t, total, count)
}

func TestOrder(t *testing.T) {
	initClient()

	var total int

	// list orders
	orders, err1 := sc.Order.List(ctx, nil)
	if err1 != nil {
		log.Fatal(err1)
	}
	total = len(orders)
	fmt.Println(total, toJsonStr(orders, false))

	// a single order
	if len(orders) > 0 {
		id := orders[0].Id
		order, err2 := sc.Order.Get(ctx, id, nil)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println(toJsonStr(order, true))
	}

	// a count of orders
	count, err3 := sc.Order.Count(ctx, nil)
	if err3 != nil {
		log.Fatal(err3)
	}
	assert.Equal(t, total, count)
}

func TestAbandonedCheckout(t *testing.T) {
	initClient()

	abandonedCheckouts, err := sc.AbandonedCheckout.List(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(abandonedCheckouts), toJsonStr(abandonedCheckouts, false))
}

func TestWebhook(t *testing.T) {
	initClient()

	// list webhooks
	webhooks, err := sc.Webhook.List(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(webhooks), toJsonStr(webhooks, false))

	// create webhook
	addWebhook := goshopify.Webhook{
		Address: "https://0694-183-240-253-203.ngrok-free.app",
		Topic:   "checkouts/create",
		Format:  "json",
	}
	webhook, err := sc.Webhook.Create(ctx, addWebhook)
	if err != nil {
		panic(err)
	}
	fmt.Println(toJsonStr(webhook, true))

	// remove webhook
	if err = sc.Webhook.Delete(ctx, webhook.Id); err != nil {
		panic(err)
	}
}

func toJsonStr(v interface{}, p bool) string {
	var (
		data []byte
		err  error
	)
	if p {
		data, err = json.MarshalIndent(v, "", "\t")
	} else {
		data, err = json.Marshal(v)
	}
	if err != nil {
		panic(err)
	}
	return string(data)
}
