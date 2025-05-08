CREATE TABLE IF NOT EXISTS admin_user (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    senha TEXT NOT NULL,
    numero VARCHAR(20),
    token TEXT,
    rule_id INTEGER REFERENCES admin_rule(id) ON DELETE SET NULL,
    enterprise_id INTEGER REFERENCES admin_enterprise(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);