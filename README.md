# Parking Lot System

## Overview

This project implements a parking lot system in Golang using the Gin framework. It provides functionalities to manage parking lots, slots, and book slots.

## Features

- Parking lots CRUD
- Slots of a parking lot CRUD
- Book slots for parking CRUD
- Stats of Book slots

## Installation

### Prerequisites

- Go installed on your machine
- PostgreSQL database

### Steps

1. **Clone the repository:**

```bash
git clone https://github.com/salmanahmad2/parking-system.git
cd parking-system
```

2. **Install dependencies:**

    ```bash
    go mod tidy
    ```
3. **Create Postgres User & DB**
 ```bash
psql -U postgres -c "CREATE USER carparking PASSWORD 'carparking123' SUPERUSER"
psql -U postgres -c "CREATE DATABASE carparking OWNER carparking"
```
4. **Run Migration:**

   Run the following bash script to run all the migrations from the root dir for the initial setup.

```bash
bash ./db/migrations/all_up.sh
```

5. **Run Project**
```bash
make run
```
## Parking Lot System API Endpoints

### User Endpoints

- **POST /user**: Create a new user
- **PUT /user/:id**: Update user details by ID
- **GET /user/:id**: Get user details by ID
- **DELETE /user/:id**: Delete user by ID
- **GET /user**: Get all users

### Parking Lot Endpoints

- **POST /lot**: Create a new parking lot
- **PUT /lot/:id**: Update parking lot details by ID
- **GET /lot/:id**: Get parking lot details by ID
- **DELETE /lot/:id**: Delete parking lot by ID
- **GET /lot**: Get all parking lots

### Slot Endpoints

- **POST /slot**: Create a new parking slot
- **PATCH /slot/:id**: Update parking slot status by ID
- **GET /slot/:id**: Get parking slot details by ID
- **DELETE /slot/:id**: Delete parking slot by ID
- **GET /slot**: Get all parking slots

### Book Slot Endpoints

- **POST /book**: Book a parking slot
- **PATCH /book/:id**: Update booked parking slot details by ID
- **GET /book/:id**: Get booked parking slot details by ID
- **DELETE /book/:id**: Delete booked parking slot by ID
- **GET /book**: Get all booked parking slots
- **GET /book/stats**: Get statistics for booked parking slots


