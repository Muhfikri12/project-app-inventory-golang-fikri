-- Table: users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR,
    password VARCHAR,
    status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table: products
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    code VARCHAR,
    stocks INT,
    category_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table: inventories
CREATE TABLE inventories (
    id SERIAL PRIMARY KEY,
    product_id INT,
    row INT,
    part INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table: transactions
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    product_id INT,
    qty INT,
    is_out BOOLEAN default TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Table: categories
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);


INSERT INTO categories (name) VALUES
('Computer Science'),
('Programming'),
('Software Engineering'),
('Mathematics'),
('Artificial Intelligence'),
('Database Systems'),
('Networking'),
('Operating Systems'),
('Web Development'),
('Machine Learning'),
('Data Science'),
('Algorithms'),
('Cyber Security'),
('Cloud Computing'),
('Discrete Mathematics');

INSERT INTO products (name, code, stocks, category_id) VALUES
('Introduction to Algorithms', 'B001', 10, 1),
('Clean Code', 'B002', 5, 2),
('Design Patterns', 'B003', 7, 3),
('The Pragmatic Programmer', 'B004', 8, 4),
('Artificial Intelligence: A Modern Approach', 'B005', 3, 5),
('Database System Concepts', 'B006', 6, 6),
('Computer Networking: A Top-Down Approach', 'B007', 4, 1),
('Operating System Concepts', 'B008', 9, 6),
('Head First Java', 'B009', 10, 2),
('Python Crash Course', 'B010', 15, 2),
('JavaScript: The Good Parts', 'B011', 5, 3),
('Introduction to Machine Learning', 'B012', 7, 5),
('Modern Operating Systems', 'B013', 3, 6),
('Programming in Haskell', 'B014', 4, 4),
('Discrete Mathematics and Its Applications', 'B015', 8, 1);

INSERT INTO inventories (product_id, row, part) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 3),
(4, 1, 4),
(5, 2, 1),
(6, 2, 2),
(7, 2, 3),
(8, 2, 4),
(9, 3, 1),
(10, 3, 2),
(11, 3, 3),
(12, 3, 4),
(13, 4, 1),
(14, 4, 2),
(15, 4, 3);

insert into users (username, password, status) values ("admin", "admin123", true);



