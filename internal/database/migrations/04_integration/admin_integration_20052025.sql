CREATE TABLE IF NOT EXISTS admin_integracao_marcas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL UNIQUE
);
CREATE TABLE IF NOT EXISTS admin_integracoes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL, -- ex: Gemini
    marca_id INTEGER NOT NULL REFERENCES admin_integracao_marcas(id) ON DELETE CASCADE,
    UNIQUE(nome, marca_id)
);
CREATE TABLE IF NOT EXISTS admin_integracao_enterprise (
    id SERIAL PRIMARY KEY,
    enterprise_id INTEGER NOT NULL REFERENCES admin_enterprise(id) ON DELETE CASCADE,
    integracao_id INTEGER NOT NULL REFERENCES admin_integracoes(id) ON DELETE CASCADE,
    UNIQUE(enterprise_id, integracao_id)
);
CREATE TABLE IF NOT EXISTS admin_integracao_user (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES admin_user(id) ON DELETE CASCADE,
    integracao_id INTEGER NOT NULL REFERENCES admin_integracoes(id) ON DELETE CASCADE,
    UNIQUE(user_id, integracao_id)
);
CREATE TABLE IF NOT EXISTS admin_integracao_tokens (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL UNIQUE,
    user_id INTEGER NOT NULL REFERENCES admin_user(id) ON DELETE CASCADE,
    integracao_id INTEGER NOT NULL REFERENCES admin_integracoes(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
