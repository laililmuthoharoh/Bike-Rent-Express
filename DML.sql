-- insert user
INSERT INTO users(name, username, password, address, role, telp) VALUES ('user', 'user', '$2y$10$4y/qwZDvEO6VUGSE2s1Xx.LgbDe67wKTHUzIIvkPfZJ85owcn6Uby', 'bumi', 'USER', '081231231');
-- insert admin
INSERT INTO users(name, username, password, address, role, telp) VALUES ('admin', 'admin', '$2y$10$1FPBWgESfbSNwl/B1RHsw./niphWxbxNCCx9eF8r5FMLXB8GWWs4.', 'bekasi', 'ADMIN', '081231231');

-- insert balance
INSERT INTO balance(amount, user_id) VALUES (0, 'c11ac713-19be-4d73-874e-e748ea6d2738');

-- insert motor_vechile
INSERT INTO motor_vechile(name, type, price, plat, production_year, status)
VALUES('Honda ct 125', 'KOPLING', 150000, 'BB2423KG', '2019', 'AVAILABLE');

--insert employee
INSERT INTO employee(name, telp, username, password) 
VALUES('Didi', '0812321342', 'didi123', '$2y$10$j7kqZIf7upB2XZ6KYjZYsehivlSIoPQNPDWIEvHae/ftgCxv2IIP2');
