-- create user
INSERT INTO users(name, username, password, address, role, telp) VALUES ('user', 'user', '$2y$10$4y/qwZDvEO6VUGSE2s1Xx.LgbDe67wKTHUzIIvkPfZJ85owcn6Uby', 'bumi', 'USER', '081231231');
-- create admin
INSERT INTO users(name, username, password, address, role, telp) VALUES ('admin', 'admin', '$2y$10$1FPBWgESfbSNwl/B1RHsw./niphWxbxNCCx9eF8r5FMLXB8GWWs4.', 'bekasi', 'ADMIN', '081231231');

-- create balance
INSERT INTO balance(amount, user_id) VALUES (0, 'c11ac713-19be-4d73-874e-e748ea6d2738');