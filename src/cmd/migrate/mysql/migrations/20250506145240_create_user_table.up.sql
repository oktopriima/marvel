CREATE TABLE `users`
    (
        `id`                int(10) unsigned NOT NULL AUTO_INCREMENT,
        `name`              varchar(100)              DEFAULT NULL,
        `email`             varchar(50)      NOT NULL,
        `email_verified_at` timestamp        NULL     DEFAULT NULL,
        `password`          varchar(255)     NOT NULL DEFAULT '',
        `remember_token`    varchar(255)              DEFAULT '',
        `created_at`        timestamp        NOT NULL DEFAULT current_timestamp(),
        `updated_at`        timestamp        NOT NULL DEFAULT current_timestamp(),
        `deleted_at`        timestamp        NULL     DEFAULT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
      AUTO_INCREMENT = 1
      DEFAULT CHARSET = utf8;