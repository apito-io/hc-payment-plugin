package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-payment-plugin] Starting Payment plugin...")
	plugin := sdk.Init("hc-payment-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("PaymentStatus", "Payment plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getPaymentStatus",
		sdk.ComplexObjectField("Get payment plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-payment-plugin] Plugin ready")
	plugin.Serve()
}
