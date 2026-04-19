# Restaurant API - High-Performance Go Backend

A robust and scalable RESTful API built with **Go (Golang)**, designed with a focus on clean architecture, optimized performance, and modern DevOps practices. This project demonstrates the implementation of a backend system without a heavy framework, utilizing a curated selection of industry-standard libraries.

## 🚀 Project Overview

The **Restaurant API** is a specialized backend service for managing restaurant operations, focusing on high concurrency, secure authentication, and efficient data caching. By avoiding monolithic frameworks, the project maintains a lightweight footprint and offers granular control over every component of the system.

## 🏗️ System Architecture: Layered Pattern

This project strictly adheres to a **Layered Architecture** to ensure a clear separation of concerns, making the codebase highly maintainable and testable.

### The Request Flow:

`Client Request` ➡️ `Presentation Layer (Controller)` ➡️ `Business Logic Layer (Service)` ➡️ `Data Access Layer (Repository)` ➡️ `Database`

- **Presentation Layer (Controller):** Handles incoming HTTP requests, input validation, and delivers the appropriate HTTP response.
- **Business Logic Layer (Service):** Orchestrates core application rules and business logic, serving as the bridge between user requests and data persistence.
- **Data Access Layer (Repository):** Manages all interactions with the database (MariaDB) using GORM, abstracting the data storage logic.
- **Models:** Represent the structural schema of the database tables, serving as a unified contract across all layers.

## 🔐 Security & Authentication

The system implements a sophisticated security layer combining **JWT (JSON Web Tokens)** and **OAuth 2.0** principles:

- **RSA 256 Encryption:** Access tokens are secured using the RSA 256 asymmetric algorithm, providing a higher level of security compared to symmetric signatures.
- **Two-Factor Authorization Flow:**
    1.  **API Key Requirement:** Initial login requests require a valid API Key.
    2.  **Credential Verification:** Upon successful validation of username and password, the server issues a JWT.
    3.  **Secure Access:** The issued JWT grants authorized access to protected resources across the API.

## ⚡ Performance Optimization & Caching

To enhance performance, this project utilizes **Redis** for strategic data caching:

- **Strategic Caching:** Frequently accessed but rarely changed data, such as `Cuisines`, are cached in Redis to minimize database load.
- **Cache Consistency:** Implements a proactive cache invalidation strategy. The system automatically flushes specific cache keys during `Create`, `Update`, or `Delete` operations to ensure data integrity.

## 🛠️ Tech Stack & Key Libraries

Built with a curated set of high-quality libraries to maintain a lightweight and efficient environment:

| Library                                                           | Purpose                                             |
| :---------------------------------------------------------------- | :-------------------------------------------------- |
| **[GORM](https://github.com/go-gorm/gorm)**                       | Database ORM and connectivity for MariaDB           |
| **[HttpRouter](https://github.com/julienschmidt/httprouter)**     | High-performance, lightweight HTTP request router   |
| **[Viper](https://github.com/spf13/viper)**                       | Comprehensive environment configuration management  |
| **[Go-Redis](https://github.com/redis/go-redis)**                 | Advanced Redis client for high-speed caching        |
| **[Logrus](https://github.com/sirupsen/logrus)**                  | Structured and pluggable logging for error tracking |
| **[JWT-Go (v5)](https://github.com/golang-jwt/jwt)**              | Secure token-based authentication (RSA 256)         |
| **[Validator (v10)](https://github.com/go-playground/validator)** | Robust request body and data validation             |
| **[Prometheus](https://github.com/prometheus/client_golang)**     | Integrated monitoring and metrics collection        |

## 🐳 Infrastructure & CI/CD

### Containerization (Multi-Stage Docker)

The application is containerized using a **Multi-Stage Build** approach to ensure the production image is minimal, secure, and lightweight:

1.  **Build Stage:** Compiles the Go source code into a static binary.
2.  **Final Stage:** Utilizes a lightweight **Alpine Linux** image to host only the binary and necessary environment files, drastically reducing the attack surface and image size.

### Automated Deployment (GitHub Actions)

A full **CI/CD pipeline** is implemented via GitHub Actions for seamless deployment to a Linux VPS:

- **Build & Push:** Automatically builds the Docker image and pushes it to Docker Hub.
- **Automated Deployment:** Executes a Remote SSH command to the VPS to pull the latest image and restart the container, ensuring zero-downtime deployment principles.
