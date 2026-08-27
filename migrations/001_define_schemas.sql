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

    -- Timestamps
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
    difficulty VARCHAR NOT NULL,
    -- type is the subject area the challenge covers (e.g., terraform, design-patterns, data-analytics)
    type TEXT NOT NULL,
    -- target is the specific subject the challenge evaluates within its type
    -- (e.g., a design pattern name, a Terraform concept, a data-analytics technique)
    target TEXT NOT NULL,
    -- user_id is a foreign key that references the id of the user who created the challenge
    user_id INTEGER NOT NULL,

    -- Timestamps
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
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

    -- Timestamps
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
    -- summary is a textual summary of the feedback provided to the user
    summary TEXT NOT NULL,
    -- suggestions is a JSON-encoded list of suggestions for the user to improve their solution.
    -- Stored as TEXT (not JSONB) because the SQLite driver in use can't scan a json()
    -- function result into json.RawMessage — the app marshals/unmarshals it itself.
    suggestions TEXT NOT NULL CHECK(json_valid(suggestions)) DEFAULT '[]',

    -- Timestamps
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- attempts contains the definitions of the attempts that users have made to solve challenges
CREATE TABLE IF NOT EXISTS attempts (
    -- id is the unique identifier for each attempt
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    -- challenge_id is a foreign key that references the id of the challenge entry
    challenge_id INTEGER NOT NULL,
    -- feedback_id is a foreign key that references the id of the feedback entry associated with this attempt
    feedback_id INTEGER,
    -- status indicates whether the attempt was successful, failed, or is still in progress
    status TEXT NOT NULL,

    -- Timestamps
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME,

    FOREIGN KEY (challenge_id) REFERENCES challenges(id) ON DELETE CASCADE,
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
DROP TRIGGER IF EXISTS update_challenges_updated_at;
DROP TRIGGER IF EXISTS update_users_updated_at;

-- Delete tables
DROP TABLE IF EXISTS attempts;
DROP TABLE IF EXISTS feedbacks;
DROP TABLE IF EXISTS clues;
DROP TABLE IF EXISTS challenges;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
