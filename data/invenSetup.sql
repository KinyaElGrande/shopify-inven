DROP TABLE IF EXISTS `warehouses`;
DROP TABLE IF EXISTS `inventories`;

CREATE TABLE warehouses (
  id INT AUTO_INCREMENT,
  city VARCHAR(100),
  country VARCHAR(100),
  code VARCHAR(100),
  manager VARCHAR(100),
  description LONGTEXT,
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