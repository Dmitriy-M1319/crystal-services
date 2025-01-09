-- +goose Up
-- +goose StatementBegin
INSERT INTO products (product_name, company_name, client_price, purchase_price, count_on_warehouse) VALUES
('Filter A', 'Company X', 150.000, 100.000, 50),
('Filter B', 'Company Y', 200.000, 150.000, 30),
('Filter C', 'Company Z', 250.000, 200.000, 20),
('Filter D', 'Company X', 300.000, 250.000, 40),
('Filter E', 'Company Y', 350.000, 300.000, 10),
('Filter F', 'Company Z', 400.000, 350.000, 25),
('Filter G', 'Company X', 450.000, 400.000, 35),
('Filter H', 'Company Y', 500.000, 450.000, 15),
('Filter I', 'Company Z', 550.000, 500.000, 28),
('Filter J', 'Company X', 600.000, 550.000, 32),
('Filter K', 'Company Y', 650.000, 600.000, 18),
('Filter L', 'Company Z', 700.000, 650.000, 22),
('Filter M', 'Company X', 750.000, 700.000, 45),
('Filter N', 'Company Y', 800.000, 750.000, 12),
('Filter O', 'Company Z', 850.000, 800.000, 27);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE products;
-- +goose StatementEnd
