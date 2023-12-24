# Rental Property Payment Management System

Go + Vite

### Paypal Webhooks testing please use cloudflared tunnels: 
``` cloudflared tunnel --url http://localhost:80 ```

### Paypal Webhooks confirmation:

An ENV variable named: PAYPAL_WEBHOOK_ID must correspond to the webhook generated for the application.

The webhook type MUST be 'Checkout order approved' as listed on the paypal dashboard

Requires absolute domain of the application so:

``` domain changes => webhook changes => webhook id changes``` 