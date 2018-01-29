
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `meeting_member` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `slack_user_id` VARCHAR(50) NOT NULL,
    `slack_team_id` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `meeting_at` DATETIME NOT NULL,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

CREATE INDEX `slack_user_id` ON `meeting_member` (`slack_user_id`);
CREATE INDEX `slack_team_id` ON `meeting_member` (`slack_team_id`);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX `slack_team_id` ON `meeting_member`;
DROP INDEX `slack_user_id` ON `meeting_member`;
DROP TABLE IF EXISTS `meeting_member`;
