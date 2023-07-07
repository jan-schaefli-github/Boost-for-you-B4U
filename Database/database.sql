DROP Database IF EXISTS b4u;

CREATE DATABASE b4u;

USE b4u;

SET sql_mode = '';

-- -----------------------------------------------------
-- Table clan
-- -----------------------------------------------------
DROP TABLE IF EXISTS clan;

CREATE TABLE clan(
    tag         VARCHAR(100) NOT NULL UNIQUE,
    joinDate    DATE DEFAULT (CURRENT_DATE),
    
    PRIMARY KEY(tag)
);

CREATE INDEX clan_tag ON clan(tag);

-- -----------------------------------------------------
-- Table person
-- -----------------------------------------------------
DROP TABLE IF EXISTS person;

CREATE TABLE person(
    tag                 VARCHAR(100) NOT NULL UNIQUE,
    name                VARCHAR(100) NOT NULL,
    role                VARCHAR(100) NOT NULL DEFAULT 'member',
    clanRank            INTEGER NOT NULL DEFAULT 0,
	clanStatus          BOOLEAN NOT NULL DEFAULT 1,
    trophies            INTEGER NOT NULL DEFAULT 0,
    wholeFame           INTEGER NOT NULL DEFAULT 0,
    wholeDecksUsed      INTEGER NOT NULL DEFAULT 0,
    wholeMissedDecks    INTEGER NOT NULL DEFAULT 0,
    wholeRepairPoints   INTEGER NOT NULL DEFAULT 0,
    wholeBoatAttacks    INTEGER NOT NULL DEFAULT 0,
    joinDate            DATE DEFAULT (CURRENT_DATE),
    fk_clan             VARCHAR(100),

    PRIMARY KEY(tag),
    FOREIGN KEY(fk_clan) REFERENCES clan(tag) ON DELETE SET NULL
);

CREATE INDEX person_tag ON person(tag);

CREATE INDEX person_fk_clan ON person(fk_clan);

-- -----------------------------------------------------
-- Table daily_report
-- -----------------------------------------------------
DROP TABLE IF EXISTS daily_report;

CREATE TABLE daily_report(
    id	                INTEGER NOT NULL UNIQUE AUTO_INCREMENT,
    fameToday	        INTEGER NOT NULL DEFAULT 0,
    decksUsedToday      INTEGER NOT NULL DEFAULT 0,
    missedDecksToday    INTEGER NOT NULL DEFAULT 0,
    repairPointsToday   INTEGER NOT NULL DEFAULT 0,
    boatAttacksToday    INTEGER NOT NULL DEFAULT 0,
    dayIdentifier       VARCHAR(100) NOT NULL,
    date	            DATE DEFAULT (CURRENT_DATE),
    fk_person           VARCHAR(100) NOT NULL,

    PRIMARY KEY(id),
    FOREIGN KEY(fk_person) REFERENCES person(tag) ON DELETE CASCADE
);

CREATE INDEX daily_report_id ON daily_report(id);

CREATE INDEX daily_report_fk_person ON daily_report(fk_person);

-- -----------------------------------------------------
-- Table weekly_report
-- -----------------------------------------------------
DROP TABLE IF EXISTS weekly_report;

CREATE TABLE weekly_report(
    id	                    INTEGER NOT NULL UNIQUE AUTO_INCREMENT,
    fameThisWeek	        INTEGER NOT NULL DEFAULT 0,
    decksUsedThisWeek       INTEGER NOT NULL DEFAULT 0,
    missedDecksThisWeek     INTEGER NOT NULL DEFAULT 0,
    repairPointsThisWeek    INTEGER NOT NULL DEFAULT 0,
    boatAttacksThisWeek     INTEGER NOT NULL DEFAULT 0,
    weekIdentifier          VARCHAR(100) NOT NULL,
    date	                DATE DEFAULT (CURRENT_DATE),
    fk_person               VARCHAR(100) NOT NULL,

    PRIMARY KEY(id),
    FOREIGN KEY(fk_person) REFERENCES person(tag) ON DELETE CASCADE
);

CREATE INDEX weekly_report_id ON weekly_report(id);

CREATE INDEX weekly_report_fk_person ON weekly_report(fk_person);