-- 用户
CREATE TABLE IF NOT EXISTS users (
    id              int PRIMARY KEY AUTO_INCREMENT,
    username        varchar(127) NOT NULL,
    password        varchar(127) NOT NULL,
    nick_name       varchar(127) NOT NULL,
    display_name    varchar(127) NOT NULL,
    gender          char NOT NULL DEFAULT '0',
    birthdate       date NOT NULL ,
    email           varchar(127) NULL,
    email_verified  tinyint(1) NOT NULL DEFAULT 0,
    phone           varchar(20) NULL,
    phone_verified  tinyint(1) NOT NULL DEFAULT 0,
    addr            varchar(255) NOT NULL,
    avatar          varchar(255) NOT NULL,
    type            int NOT NULL,
    is_online       tinyint(1) NOT NULL DEFAULT 0,
    is_disabled     tinyint(1) NOT NULL DEFAULT 0,
    created_at      timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_users_username (username),
    INDEX idx_users_email (email),
    INDEX idx_users_phone (phone)
) ENGINE=InnoDB DEFAULT CHARSET= utf8;