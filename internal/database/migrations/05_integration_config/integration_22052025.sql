CREATE TABLE IF NOT EXISTS integracoes_configuracoes (
    id SERIAL PRIMARY KEY,
    enterprise_id INTEGER NOT NULL REFERENCES admin_enterprise(id) ON DELETE CASCADE,
    integracao_id INTEGER NOT NULL REFERENCES admin_integracoes(id) ON DELETE CASCADE,
    configuracoes JSONB NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT unique_integracao_config UNIQUE (configuracoes)
);
