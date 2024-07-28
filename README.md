# Technical Task
## Problem

A product ships in various pack sizes:

- 250 Items
- 500 Itmes
- 1000 Items
- 2000 Items
- 5000 Items

A customer can order any number of these items, but they will always only be given complete packs.

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

## Solution caveat

The solution proposed has a caveat - it assumes that all pack sizes are multiples of the smallest pack size.
This is true for the exact problem listed above, but if pack sizes were changed so that this was no longer true, the solution no longer works.

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

To update the pack sizes for the API, change the config file and start runing the API (instructions [here](#Run-the-API-locally))