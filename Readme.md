# Vector Vault

Simple Vector Database to store and query OpenAI Embedding Vectors.

This project is more me trying out golang than anything else. Please just use Postgres or a dedicated vector DB for storage :O

## Components

### Frontend

- Using the frontend it is possible to monitor, modify and manage indices

### NGINX

- Used as a link between the management frontend and the vector index instances
- Load Balancing
- API Key authentication

### Vector Index

- A Vector Index stores and manages the data and vectors
- The vector index can be queried using vectors or strings
- Using the Frontend entries can be added, updated and removed
- There can be multiple replicas of each index

#### Swagger API Documentation
![swagger](https://raw.githubusercontent.com/SK4P3/vector-vault/master/docs/images/swagger.png)

## Usage

### Frontend

In `./frontend` run:

- `npm install`
- `npm start`

### Vector Index

Create a `.env` file in src with following content:

```
OPENAPI_KEY=YOUR_KEY
STORE_LOCATION=./mockStore/
```

## Screenshots

### Dashboard
![dashboard](https://raw.githubusercontent.com/SK4P3/vector-vault/master/docs/images/dashboard.png)
### Index Overview
![dashboard](https://raw.githubusercontent.com/SK4P3/vector-vault/master/docs/images/indices.png)
### Index Entries
![dashboard](https://raw.githubusercontent.com/SK4P3/vector-vault/master/docs/images/entries.png)
### Add Entry
![dashboard](https://raw.githubusercontent.com/SK4P3/vector-vault/master/docs/images/addEntry.png)
