-- +goose Up
-- +goose StatementBegin

-- users contains the definitions of the users who can attempt to solve challenges
CREATE TABLE IF NOT EXISTS users (
    -- id is the unique identifier for each user
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- username is the unique username of the user
    username TEXT NOT NULL UNIQUE,
    -- email is the unique email address of the user
    email TEXT NOT NULL UNIQUE,

    -- Timestaps 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- challenges contains the definitions of the challenges that users 
-- can attempt to solve
CREATE TABLE IF NOT EXISTS challenges (
    -- id is the unique identifier for each challenge
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- name is the name of the challenge
    name TEXT NOT NULL,
    -- description is a detailed description of the challenge
    description TEXT NOT NULL,
    -- difficulty is the difficulty level of the challenge (e.g., easy, medium, hard)
    difficulty INTEGER NOT NULL,
    -- type is the type of the pattern that the challenge is based on (e.g., creational, structural, behavioral)
    type TEXT NOT NULL,
    -- target_pattern is the specific design pattern that the challenge is targeting
    target_pattern TEXT NOT NULL,

    -- Timestaps 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- user_challenges contains the definitions of the challenges that users have attempted to solve
CREATE TABLE IF NOT EXISTS user_challenges (
    -- id is the unique identifier for each user challenge entry
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- user_id is a foreign key that references the id of the user attempting the challenge
    user_id INTEGER NOT NULL,
    -- challenge_id is a foreign key that references the id of the challenge being attempted
    challenge_id INTEGER NOT NULL,
    -- status indicates whether the user has completed, is in progress, or has not started the challenge
    status TEXT NOT NULL,

    -- Timestaps 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON DELETE CASCADE
);

-- clues contains the definitions of the clues that are associated with each challenge
CREATE TABLE IF NOT EXISTS clues (
    -- id is the unique identifier for each clue
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- challenge_id is a foreign key that references the id of the challenge this clue belongs to
    challenge_id INTEGER NOT NULL,
    -- description is the description of the clue
    description TEXT NOT NULL,
    -- sequence_order indicates the sequence in which clues should be presented to the user
    sequence_order INTEGER NOT NULL,

    -- Timestaps 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON DELETE CASCADE
);

-- feedbacks contains the feedback provided by users for each challenge
CREATE TABLE IF NOT EXISTS feedbacks (
    -- id is the unique identifier for each feedback entry
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- score is a numerical score provided by the AI model to score the user's solution
    score INTEGER NOT NULL,
    -- rating is a numerical rating provided by the user (e.g., 1-5)
    rating INTEGER NOT NULL,
    -- summary is a textual summary of the feedback provided to the user
    summary TEXT NOT NULL,
    -- suggestions is a JSONB field that contains a list of suggestions for the user to improve their solution
    suggestions JSONB NOT NULL CHECK(json_valid(suggestions)) DEFAULT '[]', 

    -- Timestaps 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- attemps contains the definitions of the attempts that users have made to solve challenges
CREATE TABLE IF NOT EXISTS attempts (
    -- id is the unique identifier for each attempt
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- user_challenge_id is a foreign key that references the id of the user challenge entry
    user_challenge_id INTEGER NOT NULL,
    -- feedback_id is a foreign key that references the id of the feedback entry associated with this attempt
    feedback_id INTEGER NULL,
    -- status indicates whether the attempt was successful, failed, or is still in progress
    status TEXT NOT NULL,

    -- Timestaps 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_challenge_id) REFERENCES user_challenges(id) ON DELETE CASCADE,
    FOREIGN KEY (feedback_id) REFERENCES feedbacks(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose StatementBegin

-- update_users_updated_at contains the trigger definition to update the updated_at automatically when a user is updated
CREATE TRIGGER IF NOT EXISTS update_users_updated_at
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
    UPDATE users 
    SET updated_at = CURRENT_TIMESTAMP 
    WHERE id = OLD.id;
END;

-- update_challenges_updated_at contains the trigger definition to update the updated_at automatically when a challenge is updated
CREATE TRIGGER IF NOT EXISTS update_challenges_updated_at
AFTER UPDATE ON challenges
FOR EACH ROW
BEGIN
    UPDATE challenges 
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- update_user_challenges_updated_at contains the trigger definition to update the updated_at automatically when a user challenge is updated
CREATE TRIGGER IF NOT EXISTS update_user_challenges_updated_at
AFTER UPDATE ON user_challenges
FOR EACH ROW
BEGIN
    UPDATE user_challenges
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- update_clues_updated_at contains the trigger definition to update the updated_at automatically when a clue is updated
CREATE TRIGGER IF NOT EXISTS update_clues_updated_at
AFTER UPDATE ON clues
FOR EACH ROW
BEGIN
    UPDATE clues
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- update_feedback_updated_at contains the trigger definition to update the updated_at automatically when a feedback entry is updated
CREATE TRIGGER IF NOT EXISTS update_feedback_updated_at
AFTER UPDATE ON feedbacks
FOR EACH ROW
BEGIN
    UPDATE feedbacks
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

-- update_attempts_updated_at contains the trigger definition to update the updated_at automatically when an attempt entry is updated
CREATE TRIGGER IF NOT EXISTS update_attempts_updated_at
AFTER UPDATE ON attempts
FOR EACH ROW
BEGIN
    UPDATE attempts
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;    
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Delete Triggers
DROP TRIGGER IF EXISTS update_attempts_updated_at;
DROP TRIGGER IF EXISTS update_feedback_updated_at;
DROP TRIGGER IF EXISTS update_clues_updated_at;
DROP TRIGGER IF EXISTS update_user_challenges_updated_at;
DROP TRIGGER IF EXISTS update_challenges_updated_at;
DROP TRIGGER IF EXISTS update_users_updated_at;

-- Deleta tables
DROP TABLE IF EXISTS attempts;
DROP TABLE IF EXISTS feedbacks;
DROP TABLE IF EXISTS clues;
DROP TABLE IF EXISTS user_challenges;
DROP TABLE IF EXISTS challenges;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
