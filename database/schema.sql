CREATE DATABASE IF NOT EXISTS chombidb;
USE chombidb;

CREATE TABLE Roles(
    id BINARY(16) PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE Users(
    id BINARY(16) PRIMARY KEY,
    firstName VARCHAR(60) NOT NULL,
    lastName VARCHAR(60) NOT NULL,
    email VARCHAR(250) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    role_id BINARY(16) NOT NULL,

    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
        REFERENCES Roles(id)
);

CREATE TABLE Vehicles(
    id BINARY(16) PRIMARY KEY,
    licensePlate VARCHAR(20) UNIQUE NOT NULL,
    unitNumber INT NOT NULL,
    shift INT NOT NULL,
    isWorking TINYINT(1) DEFAULT 0,
    image_url VARCHAR(255),
    number_of_passenggers INT,
    model VARCHAR(100) NOT NULL,
    driver_name VARCHAR(120) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    user_id BINARY(16) NOT NULL,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES Users(id)
        ON DELETE CASCADE
);

CREATE TABLE vehicle_shift_history(
    id BINARY(16) PRIMARY KEY,
    vehicle_id BINARY(16) NOT NULL,
    shift_order INT NOT NULL,
    snapshot_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_vehicle_history
        FOREIGN KEY (vehicle_id)
        REFERENCES Vehicles(id)
        ON DELETE CASCADE,

    UNIQUE(vehicle_id, snapshot_date)
);
