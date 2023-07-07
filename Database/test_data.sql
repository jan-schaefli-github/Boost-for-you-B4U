INSERT INTO clan (tag) VALUES ('#P9UVQCJV');
INSERT INTO clan (tag) VALUES ('#QPPQ2LQP');



SELECT p.tag, p.name, p.joinDate, p.clanStatus, p.role, p.trophies, p.clanRank, dr.fameToday, dr.decksUsedToday, dr.missedDecksToday, dr.repairPointsToday, dr.boatAttacksToday, dr.dayIdentifier
FROM person p
INNER JOIN daily_report dr ON p.tag = dr.fk_person
WHERE p.fk_clan = '#P9UVQCJV'