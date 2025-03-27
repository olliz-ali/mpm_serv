SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'ai_coach_db';

--@block
DROP TABLE muscle_to_train;
DROP TABLE equipment;

DROP TABLE user;