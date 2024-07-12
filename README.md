
# Esynx CMS

## Author and Designer

**Author**: Mohammad Julfikar  
**Email**: [julfikar@eztech.com.my](mailto:julfikar@eztech.com.my)

## Introduction

This project follows the Hexagonal Architecture (also known as Ports and Adapters) combined with Domain-Driven Design (DDD) principles. The aim is to maintain a clean separation between the core business logic and the external systems, such as databases and user interfaces, ensuring the system is flexible, maintainable, and testable.

## Project Structure

The project is structured into several key directories, each with specific responsibilities. This modular design allows for easy maintenance and scalability.

### Directory Breakdown

#### `contracts/`

- **Purpose**: Contains interfaces that define how the core domain interacts with external systems.
- **Example Files**:
    - `cms_creditnote.go`
    - `cms_customer.go`
    - `database.go`

#### `entities/`

- **Purpose**: Holds the core domain entities, representing the business objects and logic.
- **Example Files**:
    - `cms_creditnote.go`
    - `cms_customer.go`
    - `cms_invoice.go`

#### `models/`

- **Purpose**: Contains additional business logic related to the entities, such as value objects or domain services.
- **Example Files**:
    - `cms_creditnote.go`
    - `cms_customer.go`
    - `cms_invoice.go`

#### `repositories/`

- **Purpose**: Implements the data access logic, providing the necessary operations to interact with the database.
- **Subdirectories**:
    - `firebase/`
    - `mysql/`
        - `agent/`
        - `creditnote/`
        - `customer/`
        - `debitnote/`
        - `invoice/`
        - `module/`
        - `stock/`

#### Root Files

- **`esynx.go`**: The main application entry point, responsible for initializing components and starting the service.
- **`esynx_test.go`**: Contains unit tests for the core functionalities.
- **`go.mod` and `go.sum`**: Dependency management files for Go modules.

## Hexagonal Architecture Overview

Hexagonal Architecture, also known as Ports and Adapters, is a design pattern that aims to create a flexible and maintainable system by decoupling the core business logic from external systems. This approach emphasizes the separation of concerns, making the system more modular and easier to test.

### Key Concepts

1. **Domain (Core)**:
    - **Entities**: Core business objects that represent the main concepts and rules of the domain.
    - **Value Objects**: Objects that represent a descriptive aspect of the domain with no conceptual identity.
    - **Domain Services**: Services that encapsulate domain logic not naturally fitting within entities or value objects.

2. **Ports**:
    - **Interfaces** that define how the core interacts with external systems.
    - Can be divided into **incoming ports** (how the application interacts with the core) and **outgoing ports** (how the core interacts with external systems).

3. **Adapters**:
    - Implementations of the ports that connect the core to the external systems.
    - Examples include database adapters, user interfaces, and external APIs.

### Diagram

Below is a visual representation of Hexagonal Architecture:

```
       +------------------+
       |    Adapters      |
       | (Infrastructure) |
       +--------+---------+
                |
+---------------+----------------+
|               |                |
|           Ports (Interfaces)   |
|               |                |
|      +--------+---------+      |
|      | Domain (Business)|      |
|      +------------------+      |
|      |     Entities     |      |
|      +------------------+      |
+--------------------------------+
```

### Benefits of Hexagonal Architecture

1. **Decoupling**: The core domain logic is decoupled from external systems, making it easier to manage and modify each part independently.
2. **Testability**: The separation of concerns allows for easier testing of individual components. The core business logic can be tested without involving external systems.
3. **Flexibility**: Adapters can be swapped or modified without affecting the core logic, allowing for easier integration with different systems or technologies.
4. **Maintainability**: The modular structure promotes cleaner code and easier maintenance.

### Implementation in the Project

In this project, the Hexagonal Architecture is implemented as follows:

- **Domain Layer**: Located in the `entities/` directory, containing the core business entities and logic.
- **Ports**: Defined in the `contracts/` directory, specifying interfaces for interacting with the core domain.
- **Adapters**: Implemented in the `repositories/` directory, containing the actual code for interacting with databases or other external systems.

## Domain-Driven Design (DDD) Principles

While following Hexagonal Architecture, this project also incorporates DDD principles. DDD focuses on the complexity of the business domain and places emphasis on creating a rich, expressive domain model.

### Key Concepts of DDD

1. **Entities**: Objects that have a distinct identity and lifecycle.
2. **Value Objects**: Objects that describe some characteristic or attribute but do not have a distinct identity.
3. **Aggregates**: A cluster of domain objects that are treated as a single unit.
4. **Repositories**: Provide methods to access aggregates.
5. **Domain Services**: Operations that do not naturally fit within an entity or value object.
6. **Factories**: Responsible for creating instances of aggregates and entities.

### Applying DDD in the Project

- **Entities**: Defined in the `entities/` directory, representing core business objects.
- **Value Objects**: Included within the entities as needed.
- **Aggregates**: Grouped entities that form a single unit.
- **Repositories**: Implemented in the `repositories/` directory, providing access to aggregates.
- **Domain Services**: Encapsulated in the `models/` directory if they do not fit within entities.
- **Factories**: May be included in the `models/` directory to handle complex creation logic.

## Development Guide

### Getting Started

1. **Clone the repository**:
   ```bash
   git clone <repository_url>
   cd esynx-cms
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

### Where to Write Code

- **Business Logic and Entities**: Add or modify files in the `entities/` directory.
- **Interfaces (Ports)**: Define new interfaces in the `contracts/` directory.
- **Implementations (Adapters)**: Implement the interfaces in the `repositories/` directory.
    - For database interactions, you can add files under `repositories/mysql/` or `repositories/firebase/`.
- **Main Application Logic**: Modify the `esynx.go` file to initialize and configure the application.

### Example Workflow

1. **Define an Interface (Port)**:
    - Create or modify an interface in the `contracts/` directory that defines the expected behavior.

2. **Implement the Interface (Adapter)**:
    - Create a corresponding implementation in the `repositories/` directory.

3. **Create or Modify Entities**:
    - Define new entities or modify existing ones in the `entities/` directory to reflect the business logic.

4. **Add Business Logic**:
    - Implement additional business rules and logic in the `models/` directory.

### Testing

- **Unit Tests**: Located in files ending with `_test.go`. Ensure all critical logic has corresponding tests to maintain code quality and reliability.

## Contact

For any questions or support, please contact Mohammad Julfikar at [julfikar@eztech.com.my](mailto:julfikar@eztech.com.my).

