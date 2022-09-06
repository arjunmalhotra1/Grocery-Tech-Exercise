# Grocery-Tech-Exercise
 Supermarket API User Stories

This is to mock a software employed by a local grocery store chain. Let's say that the company needs an application which can add and fetch produce from their system. 

| User Stories                | Narrative                                                                                 |
|-----------------------------|-------------------------------------------------------------------------------------------|
| Add new produce          | As an employee, I want to add items so that I can grow the database to include new items            |
| Fetch the produce inventory | As an employee, I want to look up produce, so that I understand what produce is available |
| Fetch an item from the inventory | As an employee, I want to look up a specific item, so that I can get information about a single product |


## How to run the project
	make docker-compose-build
	make docker-compose-up  

## How to run the Unit-tests
  make unit-tests
## How to run the API-tests
  ### Run the project
  make api-tests



## Assumptions
* The produce database is a single, in memory array of data and supports reads and writes
* Supports adding more than one new produce item at a time
* The produce includes name, produce code, and unit price
* The produce name is alphanumeric and case insensitive
* The produce codes are sixteen characters long, with dashes separating each four character group
* The produce codes are alphanumeric and case insensitive
* The produce unit price is a number with up to 2 decimal places
* Error handling (GET nonexistent produce, bad POST payload, etc.)
* The API is RESTful
* The functionality is tested and verifiable without manually exercising the endpoints
* The API supports adding and deleting individual produce. You can also get any produce in the database.


## GET /get-inventory
  Gets all the produce in the database. 
  Returns JSON array of produce. 

## GET /item/{code}
    Gets a produce in the database.
    Returns JSON object of the produce.
## POST /item/
  Add one or more new produce items to the database. Accepts JSON input with the following parameters:
   * Name - required
   * Produce Code - required
   * Unit Price - required

## DELETE /item/{code}
    Deleted a produce in the database if it exists.