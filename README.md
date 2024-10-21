# Tekken Babes Backend (Gin)

Welcome to the **Tekken Babes Gin Backend** project! This backend serves the API for the Tekken Hotness Battle application, allowing users to engage in battles between Tekken characters and track the global leaderboard.

## Table of Contents

- [Tekken Babes Backend (Gin)](#tekken-babes-backend-gin)
  - [Table of Contents](#table-of-contents)
  - [Project Overview](#project-overview)
  - [Technologies Used](#technologies-used)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Database Setup](#database-setup)
  - [Running the Application](#running-the-application)
  - [API Documentation](#api-documentation)
    - [Base URL](#base-url)
    - [Endpoints](#endpoints)
      - [**GET** `/api/battle`](#get-apibattle)
      - [**POST** `/api/battle/vote`](#post-apibattlevote)
      - [**GET** `/api/characters/{id}`](#get-apicharactersid)
      - [**GET** `/api/leaderboard`](#get-apileaderboard)
  - [Project Structure](#project-structure)
  - [Contributing](#contributing)
  - [License](#license)
  - [Contact](#contact)
  - [Acknowledgments](#acknowledgments)

---

## Project Overview

The **Tekken Babes Backend** is a RESTful API built with Go and the Gin web framework. It provides endpoints for:

- Retrieving random characters for battles.
- Submitting votes and tracking win streaks.
- Displaying detailed character profiles.
- Accessing the global leaderboard.

---

## Technologies Used

- **Go (Golang)**: Programming language used for backend development.
- **Gin**: HTTP web framework for building APIs.
- **GORM**: ORM library for database interactions.
- **PostgreSQL**: Relational database for storing character and battle data.
- **godotenv**: For managing environment variables.

---

## Getting Started

### Prerequisites

- **Go**: Version 1.16 or higher.
- **PostgreSQL**: Ensure you have a PostgreSQL database installed and running.
- **Git**: For cloning the repository.

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/tekken-hotness-backend.git
   cd tekken-hotness-backend
   ```

2. **Initialize Go Modules**

   If the `go.mod` file is not present, initialize the module:

   ```bash
   go mod init github.com/yourusername/tekken-hotness-backend
   ```

3. **Install Dependencies**

   ```bash
   go mod tidy
   ```

### Database Setup

1. **Create the Database**

   Create a new PostgreSQL database:

   ```sql
   CREATE DATABASE tekken_babes_db;
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the project root:

   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=tekken_babes_db
   ```

   Replace the placeholders with your actual database credentials.

---

## Running the Application

1. **Run Migrations and Seed Data**

   The application will automatically run migrations and seed initial data when it starts.

2. **Start the Server**

   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

---

## API Documentation

### Base URL

- **Local Development**: `http://localhost:8080/api`

### Endpoints

#### **GET** `/api/battle`

- **Description**: Retrieve two random characters for a battle.
- **Response**:

  ```json
  {
    "characterOne": {
      "id": 1,
      "name": "Character One",
      "image_url": "http://example.com/character_one.jpg",
      "bio": "Bio of character one",
      "total_wins": 10
    },
    "characterTwo": {
      "id": 2,
      "name": "Character Two",
      "image_url": "http://example.com/character_two.jpg",
      "bio": "Bio of character two",
      "total_wins": 8
    }
  }
  ```

#### **POST** `/api/battle/vote`

- **Description**: Submit a vote for the hotter character.
- **Request Body**:

  ```json
  {
    "winnerId": 1,
    "loserId": 2,
    "currentStreak": 3
  }
  ```

- **Response**:

  - If win streak is less than 5:

    ```json
    {
      "message": "Vote recorded",
      "newStreak": 4
    }
    ```

  - If win streak reaches 5:

    ```json
    {
      "message": "Character wins!",
      "characterId": 1
    }
    ```

#### **GET** `/api/characters/{id}`

- **Description**: Get detailed information about a character.
- **Response**:

  ```json
  {
    "id": 1,
    "name": "Character One",
    "image_url": "http://example.com/character_one.jpg",
    "bio": "Bio of character one",
    "total_wins": 15
  }
  ```

#### **GET** `/api/leaderboard`

- **Description**: Retrieve the leaderboard of characters ranked by total wins.
- **Response**:

  ```json
  [
    {
      "id": 1,
      "name": "Character One",
      "total_wins": 15
    },
    {
      "id": 2,
      "name": "Character Two",
      "total_wins": 12
    }
    // ... more characters
  ]
  ```

---

## Project Structure

```
tekken-hotness-backend/
├── controllers/
│   ├── battle_controller.go
│   └── character_controller.go
├── database/
│   ├── database.go
│   └── seeder.go
├── models/
│   ├── battle.go
│   └── character.go
├── routes/
│   └── routes.go
├── .env
├── go.mod
├── go.sum
└── main.go
```

---

## Contributing

We welcome contributions to improve the project!

1. **Fork the Repository**

   ```bash
   git clone https://github.com/yourusername/tekken-hotness-backend.git
   ```

2. **Create a Feature Branch**

   ```bash
   git checkout -b feature/YourFeatureName
   ```

3. **Commit Your Changes**

   ```bash
   git commit -m "Add Your Feature"
   ```

4. **Push to Your Fork**

   ```bash
   git push origin feature/YourFeatureName
   ```

5. **Submit a Pull Request**

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Contact

- **Project Maintainer**: [Nick Cinera](nicolas.cinera@gmail.com)
- ***

## Acknowledgments

- **Gin Web Framework**: For providing an excellent framework for building APIs.
- **GORM**: For simplifying database interactions.
- **Community Contributors**: Thanks to everyone who has contributed to this project.

---
