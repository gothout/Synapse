-- 🧩 Regras
INSERT INTO admin_rule (name) VALUES
('admin'),               -- rule_id = 1
('operator');            -- rule_id = 2

-- 📦 Módulos
INSERT INTO admin_module (name) VALUES
('admin.user'),          -- module_id = 1
('admin.enterprise'),    -- module_id = 2
('admin.rules'),         -- module_id = 3
('admin.integration');   -- module_id = 4

-- 🛡️ Permissões

-- admin.user (module_id = 1)
INSERT INTO admin_permission (module_id, action) VALUES
(1, 'create'),     -- permission_id = 1
(1, 'read'),       -- permission_id = 2
(1, 'update'),     -- permission_id = 3
(1, 'remove');     -- permission_id = 4

-- admin.enterprise (module_id = 2)
INSERT INTO admin_permission (module_id, action) VALUES
(2, 'create'),     -- permission_id = 5
(2, 'read'),       -- permission_id = 6
(2, 'update'),     -- permission_id = 7
(2, 'remove');     -- permission_id = 8

-- admin.rules (module_id = 3)
INSERT INTO admin_permission (module_id, action) VALUES
(3, 'create'),     -- permission_id = 9
(3, 'read'),       -- permission_id = 10
(3, 'update'),     -- permission_id = 11
(3, 'remove');     -- permission_id = 12

-- admin.integration (module_id = 4)
INSERT INTO admin_permission (module_id, action) VALUES
(4, 'create'),     -- permission_id = 13
(4, 'read'),       -- permission_id = 14
(4, 'update'),     -- permission_id = 15
(4, 'remove');     -- permission_id = 16

-- 🔗 Permissões completas para rule_id = 1 (admin)
INSERT INTO admin_rule_permission (rule_id, permission_id) VALUES
(1, 1), (1, 2), (1, 3), (1, 4),   -- admin.user
(1, 5), (1, 6), (1, 7), (1, 8),   -- admin.enterprise
(1, 9), (1, 10), (1, 11), (1, 12),-- admin.rules
(1, 13), (1, 14), (1, 15), (1, 16);-- admin.integration

-- 🔗 Permissão para operator: somente admin.enterprise.create (permission_id = 5)
INSERT INTO admin_rule_permission (rule_id, permission_id) VALUES
(2, 5);
