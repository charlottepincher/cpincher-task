# Technical Task
## Problem

A product ships in various pack sizes:

- 250 Items
- 500 Items
- 1000 Items
- 2000 Items
- 5000 Items

A customer can order any number of these items, but they will always only be given complete packs.

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

## Solution

The solution works by first finding the greatest common divisor for all the pack sizes. It then finds the next biggest number to the ordered number that is divisible by the greatest common divisor, therefore eliminating running more computations than necessary.
It tries to find a combination of packs to reach that number, if it is not possible then it adds the gcd to the current ordered number and once again tries to find a combination of packs to reach that number. This continues until it finds the first number that it can create packs for. If multiple combinations are found, it selects the combination with the fewest number of packs in. This ensures that only the minimum items are sent out and then that the minimum packs for the minimum items are sent.

## Website

https://cpincher-gymshark-vejnvl2nca-lz.a.run.app/

This is hosted using google cloud.

Input a number in the submission box and it will return the required packs.

## Run the API locally

At the top level, run:

`go run .` 

This will start the API listening on localhost:8000.

To use it, send a PUT request to 'http://localhost:8000/apitest' with the body in the following format:

`{"Order":<number>}`

eg. `{"Order":251}`

The API will return the pack sizes required, as follows:

```
5000: 0
2000: 0
1000: 0
500: 1
250: 0
```

## Config

The pack sizes are retrieved from the "config.json" file which must have the format:

`{"Packs": [250, 500, 1000, 2000, 5000]}`

To update the pack sizes for the website, change the config file and then redeploy the website.

To update the pack sizes for the API, change the config file and start running the API (instructions [here](#Run-the-API-locally))