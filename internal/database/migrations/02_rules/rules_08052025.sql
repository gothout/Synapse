-- 🧩 Regras (Papéis)
CREATE TABLE IF NOT EXISTS admin_rule (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE  -- ex: admin, operador_empresas
);

-- 📦 Módulos do sistema
CREATE TABLE IF NOT EXISTS admin_module (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE  -- ex: admin.user, admin.enterprise, admin.rules
);

-- 🛡️ Permissões por módulo
CREATE TABLE IF NOT EXISTS admin_permission (
    id SERIAL PRIMARY KEY,
    module_id INTEGER NOT NULL REFERENCES admin_module(id) ON DELETE CASCADE,
    action VARCHAR(50) NOT NULL,  -- ex: criar, remover, listar
    UNIQUE (module_id, action)
);

-- 🔗 Permissões associadas às regras
CREATE TABLE IF NOT EXISTS admin_rule_permission (
    rule_id INTEGER NOT NULL REFERENCES admin_rule(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES admin_permission(id) ON DELETE CASCADE,
    PRIMARY KEY (rule_id, permission_id)
);