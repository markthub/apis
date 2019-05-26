CREATE TABLE `users`
(
  `id` int PRIMARY KEY,
  `username` varchar(255) UNIQUE NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `customers`
(
  `id` int PRIMARY KEY,
  `user_id` int UNIQUE NOT NULL,
  `name` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `phone` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `zipcode` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `orders`
(
  `id` int PRIMARY KEY,
  `customer_id` int NOT NULL,
  `status_code` varchar(255) NOT NULL,
  `date_placed` datetime NOT NULL,
  `order_details` varchar(255),
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `invoices`
(
  `number` varchar(255) PRIMARY KEY,
  `order_id` int NOT NULL,
  `status_code` varchar(255) NOT NULL,
  `placed` datetime NOT NULL,
  `details` varchar(255),
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `payments`
(
  `id` int PRIMARY KEY,
  `invoice_number` varchar(255) UNIQUE NOT NULL,
  `placed` datetime NOT NULL,
  `amount` double NOT NULL
);

CREATE TABLE `ref_invoice_status_codes`
(
  `status_code` varchar(255) PRIMARY KEY,
  `details` varchar(255)
);

CREATE TABLE `shipments`
(
  `id` int PRIMARY KEY,
  `order_id` int NOT NULL,
  `invoice_number` varchar(255) NOT NULL,
  `tracking_number` varchar(255) NOT NULL,
  `placed` datetime NOT NULL,
  `details` varchar(255),
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `shipment_items`
(
  `shipment_id` int NOT NULL,
  `order_item_id` int NOT NULL
);

CREATE TABLE `stores`
(
  `id` int PRIMARY KEY,
  `user_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `zipcode` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `products`
(
  `id` int PRIMARY KEY,
  `store_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` varchar(255) NOT NULL,
  `description` varchar(255),
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `order_items`
(
  `id` int PRIMARY KEY,
  `product_id` int NOT NULL,
  `order_id` int NOT NULL,
  `status_code` varchar(255) NOT NULL,
  `quantity` int NOT NULL,
  `price` double NOT NULL,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime
);

CREATE TABLE `ref_order_item_status_code`
(
  `status_code` varchar(255) PRIMARY KEY,
  `description` varchar(255)
);

ALTER TABLE `customers` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `stores` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `payments` ADD FOREIGN KEY (`invoice_number`) REFERENCES `invoices` (`number`);

ALTER TABLE `invoices` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `invoices` ADD FOREIGN KEY (`status_code`) REFERENCES `ref_invoice_status_codes` (`status_code`);

ALTER TABLE `shipments` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `shipments` ADD FOREIGN KEY (`invoice_number`) REFERENCES `invoices` (`number`);

ALTER TABLE `shipment_items` ADD FOREIGN KEY (`shipment_id`) REFERENCES `shipments` (`id`);

ALTER TABLE `shipment_items` ADD FOREIGN KEY (`order_item_id`) REFERENCES `order_items` (`id`);

ALTER TABLE `orders` ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);

ALTER TABLE `products` ADD FOREIGN KEY (`store_id`) REFERENCES `stores` (`id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

ALTER TABLE `order_items` ADD FOREIGN KEY (`status_code`) REFERENCES `ref_order_item_status_code` (`status_code`);
