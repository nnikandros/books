CREATE TABLE books(  
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(100) NOT NULL,
    author  TEXT NOT NULL,
    publication_date DATE,
    rating TEXT
);