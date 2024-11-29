-- Insert categories
INSERT INTO categories (name, description) VALUES
('Food', 'Groceries, restaurants, etc.'),
('Transportation', 'Fuel, public transport, etc.'),
('Entertainment', 'Movies, hobbies, etc.'),
('Utilities', 'Electricity, water, internet, etc.');

-- Insert users
INSERT INTO users (name, email, password, currency) VALUES
('John Doe', 'john@example.com', 'password123', 'USD'),
('Jane Smith', 'jane@example.com', 'password123', 'USD');

-- Insert budgets
INSERT INTO budgets (user_id, amount, month, year) VALUES
(1, 500.00, 1, 2024),
(2, 400.00, 1, 2024);
