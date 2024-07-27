**The problem**

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


**To run**

At the top level, run:
`go run .` 

This will start the API listening on localhost:8000.

To use it, send a PUT request to 'http://localhost:8000/' with the body in the following format:

`{"Order":<number>}`

eg. `{"Order":251}`

The API will return a map of the pack sizes required, as follows:

`[250:0 500:1 1000:0 2000:0 5000:0]`