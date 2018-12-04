
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `hotel` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(200) NOT NULL,
    `address` VARCHAR(45) NULL,
    `created_at` DATETIME NULL,
    `updated_at` DATETIME NULL,
    PRIMARY KEY (`id`)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `hotel`;
