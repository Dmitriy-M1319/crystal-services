-- +goose Up
-- +goose StatementBegin
create table products(
	id serial primary key,
	product_name varchar(255),
	company_name varchar(255),
	client_price decimal(10,3),
	purchase_price decimal(10,3),
	count_on_warehouse integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
