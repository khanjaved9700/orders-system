# orders-system
Basic Orders System with Kafka, Redis, and Docker

This project implements a basic order management system that allows creating and processing customer orders. It is designed as a microservice-style architecture using modern tools:

Order Service (Go): Handles REST APIs for creating and retrieving orders. Orders are stored in a relational database (e.g., PostgreSQL/MySQL) for persistence.

Kafka: Used as a message broker to publish order events (e.g., OrderCreated, OrderCancelled) so that other services can react asynchronously. This ensures loose coupling and scalability.

Redis: Used for caching frequently accessed data (e.g., order details, inventory availability) to improve performance and reduce database load. It can also be used for rate-limiting or session management.

Docker & Docker Compose: All services (Order API, Kafka broker, Zookeeper, Redis, and database) run in containers for easy setup, isolation, and portability. Docker Compose is used to orchestrate the multi-container environment.



High-Level Flow

A client calls the Order API to create a new order.
The order is stored in the database and an OrderCreated event is published to Kafka.
Other services (e.g., Payment Service, Notification Service – optional extensions) can subscribe to Kafka topics and process events.
Order details are cached in Redis to allow fast reads.
Developers can run the whole system locally using Docker Compose without installing dependencies manually.
