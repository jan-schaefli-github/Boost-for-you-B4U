DROP Database IF EXISTS b4u;

CREATE DATABASE b4u;

USE b4u;

-- -----------------------------------------------------
-- Table clan
-- -----------------------------------------------------
DROP TABLE IF EXISTS clan;

CREATE TABLE clan(
    tag         VARCHAR(100) NOT NULL UNIQUE,
    
    PRIMARY KEY(tag)
);

-- -----------------------------------------------------
-- Table person
-- -----------------------------------------------------
DROP TABLE IF EXISTS person;

CREATE TABLE person(
    tag         VARCHAR(100) NOT NULL UNIQUE,
    wholeFame   INTEGER NOT NULL DEFAULT 0,
	clanStatus  BOOLEAN NOT NULL DEFAULT 1,
    joinDate    DATE DEFAULT (CURRENT_DATE),
    fk_clan     VARCHAR(100),

    PRIMARY KEY(tag),
    FOREIGN KEY(fk_clan) REFERENCES clan(tag) ON DELETE SET NULL
);

-- -----------------------------------------------------
-- Table weekly_report
-- -----------------------------------------------------
DROP TABLE IF EXISTS weekly_report;

CREATE TABLE weekly_report(
    id	        INTEGER NOT NULL UNIQUE AUTO_INCREMENT,
    fame	    INTEGER NOT NULL DEFAULT 0,
    missedDecks INTEGER NOT NULL DEFAULT 0,
    absenceDays INTEGER NOT NULL DEFAULT 0,
    date	    DATE DEFAULT (CURRENT_DATE),
    fk_person   VARCHAR(100) NOT NULL,

    PRIMARY KEY(id),
    FOREIGN KEY(fk_person) REFERENCES person(tag) ON DELETE CASCADE
);

-- -----------------------------------------------------
-- Table daily_report
-- -----------------------------------------------------
DROP TABLE IF EXISTS daily_report;

CREATE TABLE daily_report(
    id	            INTEGER NOT NULL UNIQUE AUTO_INCREMENT,
    decksUsedToday  INTEGER NOT NULL DEFAULT 0,
    fame	        INTEGER NOT NULL DEFAULT 0,
    dayIndex        INTEGER NOT NULL,
    date	        DATE DEFAULT (CURRENT_DATE),
    fk_person       VARCHAR(100) NOT NULL,

    PRIMARY KEY(id),
    FOREIGN KEY(fk_person) REFERENCES person(tag) ON DELETE CASCADE
);