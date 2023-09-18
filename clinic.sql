-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema clinic
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `clinic` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `clinic` ;

-- -----------------------------------------------------
-- Table `clinic`.`dentist`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinic`.`dentist` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `registration_number` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `registration_number_UNIQUE` (`registration_number` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `clinic`.`patient`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinic`.`patient` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `address` VARCHAR(80) NOT NULL,
  `dni` INT NOT NULL,
  `discharge_date` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `dni_UNIQUE` (`dni` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `clinic`.`appointment`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinic`.`appointment` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `patient_id` BIGINT NOT NULL,
  `dentist_id` BIGINT NOT NULL,
  `date` DATETIME NOT NULL,
  `description` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `appointment_dentist_dentist_id_id_idx` (`dentist_id` ASC) VISIBLE,
  INDEX `appointment_patient_patient_id_id` (`patient_id` ASC) VISIBLE,
  CONSTRAINT `appointment_dentist_dentist_id_id`
    FOREIGN KEY (`dentist_id`)
    REFERENCES `clinic`.`dentist` (`id`),
  CONSTRAINT `appointment_patient_patient_id_id`
    FOREIGN KEY (`patient_id`)
    REFERENCES `clinic`.`patient` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

-- Test records for the 'dentist' table
INSERT INTO `dentist` (`first_name`, `last_name`, `registration_number`) VALUES
 ('Dr. Smile', 'McDentist', '12345'),
 ('Dr. Sparkle', 'Tooth Fairy', '67890'),
 ('Dr. Chomp', 'Flossington', '54321');

-- Test records for the 'patient' table
INSERT INTO `patient` (`first_name`, `last_name`, `address`, `dni`, `discharge_date`) VALUES
('Toothless', 'McGums', '123 Cavity Ln', 12345678, '2023-09-15 09:00:00'),
('Candy', 'Cane', '456 Sugar Ave', 98765432, '2023-09-16 10:30:00'),
('Molar', 'Incisor', '789 Brush St', 56789012, '2023-09-17 14:15:00');

-- Test records for the 'appointment' table
INSERT INTO `appointment` (`patient_id`, `dentist_id`, `date`, `description`) VALUES
(1, 1, '2023-09-15 11:30:00', 'Appointment for a dazzling smile'),
(2, 2, '2023-09-16 15:45:00', 'Magical dental cleaning'),
(3, 3, '2023-09-17 09:30:00', 'Chew-style tooth extraction operation');

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
