CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    passport_serie INT NOT NULL,
    passport_number INT NOT NULL,
    name VARCHAR(50),
    surname VARCHAR(50),
    patronymic VARCHAR(50),
    address TEXT
);

CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    people_id INT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    time_start TIMESTAMP,
    time_end TIMESTAMP,
    CONSTRAINT task_fk0 FOREIGN KEY (people_id) REFERENCES people (id)
);