INSERT INTO clan (tag) VALUES ('#P9UVQCJV');
INSERT INTO clan (tag) VALUES ('#QPPQ2LQP');

INSERT INTO person (tag, name, wholeFame, clanStatus, fk_clan) VALUES ('#2Y9VQVJ8', 'THE DISCONECTOR', 0, 1, '#P9UVQCJV');

SELECT person.tag, person.name, person.clanStatus, weekly_report.fame, weekly_report.missedDecks, daily_report.decksUsedToday, person.fk_clan, daily_report.date
FROM weekly_report
INNER JOIN person ON weekly_report.fk_person = person.tag
INNER JOIN daily_report ON daily_report.fk_person = person.tag;