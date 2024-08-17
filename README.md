# Gcart - Go Backend App for Online Store Management

Gcart is a Go-based backend application designed for managing an online store. It leverages Gin for RESTful APIs, MongoDB as the primary database, Gnosql as the secondary database, and Gque as the message queue protocol.

## Table of Contents
- [Features](#features)
- [Entities](#entities)
- [Architecture](#architecture)
- [Databases](#databases)
- [Message Queue](#message-queue)
- [Workers](#workers)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features
- Manage entities such as city, user, product, order, item, payment, refund-payment, and events.
- RESTful API using Gin.
- MongoDB as the primary database.
- [Gnosql](https://github.com/nanda03dev/gnosql.git) as the secondary in-memory NoSQL database.
- [Gque](https://github.com/nanda03dev/gque.git) for message queue handling.

## Entities
The application manages the following entities:
- City
- User
- Product
- Order
- Item
- Payment
- Refund Payment
- Events

## Architecture
The application is structured with controllers, services, and repositories for each entity to ensure a clean and maintainable codebase.

## Databases
- **MongoDB**: Primary database for persistent storage.
- **Gnosql**: Secondary in-memory NoSQL database for quick data access.

## Message Queue
- **Gque**: Custom message queue protocol used for handling asynchronous operations.

## Workers
### 1. Entity CRUD Worker
- Consume messages from the CRUD channel.
- Construct events with a `checkProcess` flag to handle entity-specific timeouts.
- Create Gnosql documents and save them to GnoSQL database.
- Store events in the MongoDB event collection and publish to Gque.

### 2. Entity Timeout Worker
- Run a cron job every 10 seconds.
- Fetch events from the MongoDB event collection.
- Check for entity timeouts, which vary for each entity. If timed out, update the status to "timeout" in MongoDB.
- Create CRUD events with updated values and push them to the CRUD channel.

### 3. Payment Refund Worker
- Consume messages from the payment refund channel.
- Fetch order and payment details from MongoDB.
- If the order is confirmed, calculate excess refunds and create events for the refund amounts.
- If the order is timed out, calculate the paid amount to refund and create events for the refund amounts.
- Store payment refunds in the MongoDB refund payment collection and create CRUD events for payment refunds.

## Installation
To get started with Gcart, follow these steps:

1. Clone the repository:
```sh
   git clone https://github.com/nanda03dev/gcart.git
   cd gcart
```
2. Run the following command to run gcart:

```bash
go run cmd/main.go
```
