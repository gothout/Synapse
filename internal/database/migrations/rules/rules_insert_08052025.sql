-- 🧩 Regras
INSERT INTO admin_rule (name) VALUES
('admin'),               -- rule_id = 1
('operador_empresas');   -- rule_id = 2

-- 📦 Módulos
INSERT INTO admin_module (name) VALUES
('admin.user'),         -- module_id = 1
('admin.enterprise'),   -- module_id = 2
('admin.rules');        -- module_id = 3

-- 🛡️ Permissões

-- admin.user (module_id = 1)
INSERT INTO admin_permission (module_id, action) VALUES
(1, 'criar'),     -- permission_id = 1
(1, 'listar'),    -- permission_id = 2
(1, 'atualizar'), -- permission_id = 3
(1, 'remover');   -- permission_id = 4

-- admin.enterprise (module_id = 2)
INSERT INTO admin_permission (module_id, action) VALUES
(2, 'criar'),     -- permission_id = 5
(2, 'listar'),    -- permission_id = 6
(2, 'atualizar'), -- permission_id = 7
(2, 'remover');   -- permission_id = 8

-- admin.rules (module_id = 3)
INSERT INTO admin_permission (module_id, action) VALUES
(3, 'criar'),     -- permission_id = 9
(3, 'listar'),    -- permission_id = 10
(3, 'atualizar'), -- permission_id = 11
(3, 'remover');   -- permission_id = 12

-- 🔗 Permissões completas para rule_id = 1 (admin)
INSERT INTO admin_rule_permission (rule_id, permission_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4),
(1, 5), (1, 6), (1, 7), (1, 8),
(1, 9), (1, 10), (1, 11), (1, 12);

-- 🔗 Permissão para operador_empresas: somente admin.enterprise.criar (permission_id = 5)
INSERT INTO admin_rule_permission (rule_id, permission_id) VALUES
(2, 5);
