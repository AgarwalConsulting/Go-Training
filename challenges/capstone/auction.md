# Auction API

The participants will create micro-services for an auction application, either `gRPC` or `ReSTful`...

## Features

* Submit an item for an auction. One item needs to be submitted at a time, though multiple items could be submitted as well.

* List the live items for an auction. Some items can go live at a specified time in the future, and some can have a finite lifetime after which they are no longer active. Both the conditions are independent - it is possible for an item to go live immediately after successful submission and have a finite lifetime, and it is possible for an item to go live in the future and have an infinite lifetime.

* Real time stream of live items for an auction along with their current bid price.

* The services should be designed using [Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) or using any frameworks/libraries: go-kit, gin, buffalo

* The services must never lose an offer or a bid even in the event that the entire system needs to be restarted.

## API Specification

These are the high level API specs, alterations and additions to the spec to meet the features is acceptable.

There are 3 public services:

* Offers - for the items up for sale
* Bids - for making bids for an offer
* Sold Offers - for the successful bids

### Offers

This API exposes the following endpoints.

* `POST /offers`

  * Creates a new item for offer
  * Validates each offer
  * Returns ‘202 Created’ on success
  * Returns ‘422 Unprocessable Entity’ on validation failure
  * POST body specification:

```yaml
offer:
  bid_price:
    Desc: Starting or latest bid price
    Optional: No
    Type: float

  go_live:
    Desc: Date time when the offer will go live, default is immediate
    Optional: Yes
    Type: datetime

  lifetime:
    Desc: Seconds after which the item will be sold at the current bid
    Optional: Yes
    Type: int

  photo_url:
    Desc: Image URL
    Optional: Yes
    Type: str

  title:
    Desc: Product title
    Optional: No
    Type: str
```

* `GET /offers?page=0&size=10&sort=key`

  * Gets every item on offer along with their current bid price.
  * Pagination controlled by optional page and size parameters.
  * Sorting is controlled by sort key.
  * Default page is 0 (first page), size is 10 and sort key is go_live.
  * Real time Stream of Offers:
    * Choice of protocol: websocket or long-lived TCP or streaming rpc
    * Exposes `GET /offers`
    * New offer is streamed back immediately

### Bids

The Bids API exposes endpoints for making bids against current offers.

* `POST /bids`

  * Creates a new bid
  * On successful bid, returns the bid_id
  * POST body specification:

```yaml
  bid:
   bid_price:
     Desc: Bid price
     Optional: No
     Type: float

   client_id:
     Desc: Identifier for client who made the bid
     Optional: No
     Type: str

   offer_id:
     Desc: Offer Identifier
     Optional: No
     Type: str
```

* `PUT /bids/:bid_id`

  * Accepts a bid on an offer
  * On success, removes the offer from Offers listing


### Sold Offers

The sold offers exposes the API for successful bids.

* `GET /sold-offers`

  * Index of all offers that has accepted bids.
