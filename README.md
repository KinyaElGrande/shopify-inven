## Mumbi Shopify Inventory Tracker
Shopify-inven is a basic web application that implements CRUD Functionality for the following :

- Create inventory items
- Edit inventory item
- Delete an inventory item
- View a list of them

and ability to create warehouses/locations and assign an inventory to a specific location.

## Prerequisites
- Go 1.15 or higher
- MySQL 5.7 or higher
## How to Run The Project 
Clone the repository to your local machine then install all the application dependencies by running the following command on the application's root directory.
```bash
go get .
```
Then create a database and create both `inventories` and `warehouses` tables as per the `data/invenSetup.sql` then configure your database credentials inside `data/conn.go`.

Finally run the server: 
```bash
go run server/server.go
```
Now navigate to https://localhost:8000 you can see your application running with the index page showing a table list of inventories.

