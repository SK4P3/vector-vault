# Vector Vault

Simple Vector Database to store and query OpenAI Embedding Vectors.

This project is still work in progress!

## Components
### Frontend
- Using the frontend it is possible to monitor, modify and manage indices

### Vector Index
- A Vector Index stores and manages the data and vectors
- There can be multiple replicas of each index

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
