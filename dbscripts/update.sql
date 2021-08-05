ALTER TABLE `classicmodels`.`employees` 
ADD COLUMN `password` VARCHAR(255) NOT NULL DEFAULT 'password' AFTER `jobTitle`;