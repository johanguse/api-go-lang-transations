CREATE DATABASE IF NOT EXISTS `api-go-lang-transactions` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `api-go-lang-transactions`;

CREATE TABLE transactions (
    ID          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    Title       VARCHAR(255) NOT NULL,
    Description TEXT NOT NULL,
    Status      ENUM('processing', 'processed', 'created') NOT NULL,
    Amount      DECIMAL(10,2) NOT NULL CHECK (Amount >= 0),
    Date        TIMESTAMP NOT NULL,
    FromUser    VARCHAR(100) NOT NULL,
    ToUser      VARCHAR(100) NOT NULL,
    CreatedAt   TIMESTAMP NOT NULL,
    UpdatedAt   TIMESTAMP NOT NULL
);