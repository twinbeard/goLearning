-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_two_factor_9999` (
    `two_factor_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, -- Auto-incrementing primary key
    `user_id` INT UNSIGNED NOT NULL, -- Foreign key referencing the user table
    `two_factor_auth_type` ENUM('SMS', 'EMAIL', 'APP') NOT NULL, -- Type of 2FA method (SMS, Email, or Third-party App such as facebook, google)
    `two_factor_auth_secret` VARCHAR(255) NOT NULL, -- Secret information for 2FA (e.g., OTP key)
    `two_factor_phone` VARCHAR(20) NULL, -- Phone number for SMS 2FA (if applicable)
    `two_factor_email` VARCHAR(255) NULL, -- Email address for Email 2FA (if applicable)
    `two_factor_is_active` BOOLEAN NOT NULL DEFAULT TRUE, -- Activation status of the 2FA method
    `two_factor_created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Time when the 2FA method was created
    `two_factor_updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Time when the 2FA method was last updated

    -- Foreign key constraint
    FOREIGN KEY (`user_id`) REFERENCES pre_go_acc_user_base_9999 (`user_id`) ON DELETE CASCADE,

    -- Indexes for optimization
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_auth_type` (`two_factor_auth_type`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='pre_go_acc_user_two_factor_9999';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_two_factor_9999`;
-- +goose StatementEnd
