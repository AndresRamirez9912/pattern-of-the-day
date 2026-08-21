CREATE TABLE challenges (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    
    name TEXT NOT NULL,

    description TEXT NOT NULL,

    difficulty TEXT NOT NULL,

    type TEXT NOT NULL,

    target_pattern TEXT NOT NULL,

    status TEXT NOT NULL,

    created_at DATETIME NOT NULL,

    updated_at DATETIME NOT NULL
);
