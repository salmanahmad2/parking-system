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
