CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    user_role_id INT NOT NULL,
    name VARCHAR(50) NOT NULL,
    password_hash VARCHAR(60) NOT NULL,
    email VARCHAR(50) NOT NULL,
);

CREATE TABLE IF NOT EXISTS Roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS UserRoles (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES Users(id),
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES Roles(id)
);

CREATE TABLE IF NOT EXISTS Authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
)

CREATE TABLE IF NOT EXISTS Books (
    id SERIAL PRIMARY KEY,
    author_id INT NOT NULL,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    CONSTRAINT fk_author FOREIGN KEY (author_id) REFERENCES Authors(id)
)

