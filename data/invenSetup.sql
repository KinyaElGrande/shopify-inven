DROP TABLE IF EXISTS `warehouses`;
DROP TABLE IF EXISTS `inventories`;

CREATE TABLE warehouses (
  id INT AUTO_INCREMENT,
  location VARCHAR(100),
  PRIMARY KEY (id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE inventories (
	id INT AUTO_INCREMENT,
	name VARCHAR(100),
	sku VARCHAR(100),
	category VARCHAR(100),
	price INT,
	quantity INT,
	manufacturer VARCHAR(100),
	description LONGTEXT,
	warehouse_id INT,
	PRIMARY KEY (id),
	FOREIGN KEY (warehouse_id) REFERENCES warehouses (id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);