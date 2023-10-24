-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;

CREATE DATABASE `sqex23`;
USE `sqex23`;

-- ---
-- Table 'workers'
-- A table which holds currently active workers and their status updated by the minute.
-- ---

DROP TABLE IF EXISTS `workers`;

CREATE TABLE `workers` (
                           `id` INTEGER AUTO_INCREMENT,
                           `ip` VARCHAR(255) NULL DEFAULT NULL,
                           `status` INTEGER(5) NOT NULL DEFAULT -1,
                           `last_updated` DATETIME NOT NULL,
                           `created` DATETIME NOT NULL DEFAULT NOW(),
                           PRIMARY KEY (`id`)
) COMMENT 'A table which holds currently active workers and their status';

-- ---
-- Table 'jobs'
-- A table which holds a list of jobs
-- ---

DROP TABLE IF EXISTS `jobs`;

CREATE TABLE `jobs` (
                        `id` INTEGER AUTO_INCREMENT,
                        `worker_id` INTEGER NULL DEFAULT NULL,
                        `data_id` INTEGER NULL DEFAULT NULL,
                        `calculation` VARCHAR(255) NULL DEFAULT NULL,
                        PRIMARY KEY (`id`)
) COMMENT 'A table which holds a list of jobs';

-- ---
-- Table 'data'
-- Table which holds all the data to be processed.
-- ---

DROP TABLE IF EXISTS `data`;

CREATE TABLE `data` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `input_a` float NOT NULL,
                        `input_b` float NOT NULL,
                        `output` float DEFAULT NULL,
                        PRIMARY KEY (`id`)
) COMMENT 'Table which holds all the data to be processed.';

-- ---
-- Foreign Keys
-- ---

ALTER TABLE `jobs` ADD FOREIGN KEY (worker_id) REFERENCES `workers` (`id`);
ALTER TABLE `jobs` ADD FOREIGN KEY (data_id) REFERENCES `data` (`id`);

-- ---
-- Table Properties
-- ---

-- ALTER TABLE `workers` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `jobs` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
-- ALTER TABLE `data` ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
