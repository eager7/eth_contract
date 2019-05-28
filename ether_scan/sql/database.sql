DROP DATABASE IF EXISTS eth_contract;
CREATE DATABASE IF NOT EXISTS eth_contract;

CREATE TABLE IF NOT EXISTS eth_contract.t_contract_info (
    `id`                BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `address`           CHAR(40)        NOT NULL DEFAULT ''     COMMENT '合约地址',
    `creator`           CHAR(64)        NOT NULL DEFAULT ''     COMMENT '合约创建者',
    `timestamp`         INT UNSIGNED    NOT NULL DEFAULT 0      COMMENT '合约创建时间',
    `block_number`      BIGINT UNSIGNED NOT NULL DEFAULT 0      COMMENT '创建合约的区块高度',
    `transaction`       CHAR(64)        NOT NULL DEFAULT ''     COMMENT '创建合约的交易哈希',
    `hash`              CHAR(64)        NOT NULL DEFAULT ''     COMMENT '合约代码的哈希',
    PRIMARY KEY (`id`),
    UNIQUE KEY `address`(`address`),
    KEY `creator`(`creator`(10)) USING BTREE,
    KEY `number`(`block_number`) USING BTREE,
    KEY `hash`(`hash`(10)) USING BTREE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS eth_contract.t_code_info (
    `id`                BIGINT UNSIGNED             NOT NULL AUTO_INCREMENT     COMMENT '自增主键',
    `creator`           CHAR(40)                    NOT NULL DEFAULT ''         COMMENT '创建者地址',
    `address`           CHAR(40)                    NOT NULL DEFAULT ''         COMMENT '合约地址',
    `hash`              CHAR(64)                    NOT NULL DEFAULT ''         COMMENT '合约哈希',
    `code`              Text                        NOT NULL DEFAULT 'code'     COMMENT '合约代码',
    `bin`               Text                        NOT NULL DEFAULT 'bin'      COMMENT '合约二进制',
    `abi`               Text                        NOT NULL DEFAULT 'abi'      COMMENT '合约abi',
    PRIMARY KEY (`id`),
    FULLTEXT abi (`abi`),
    UNIQUE KEY (`hash`(10))
)ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;