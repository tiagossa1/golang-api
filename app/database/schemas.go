package database

const createSchema = `
CREATE TABLE IF NOT EXISTS People (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    FirstName TEXT NOT NULL,
    Surname TEXT NOT NULL
);

`
