CREATE TABLE IF NOT EXISTS integracoes_configuracoes (
    id SERIAL PRIMARY KEY,
    integracao_id INTEGER NOT NULL REFERENCES admin_integracoes(id) ON DELETE CASCADE,
    configuracoes JSONB NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    atualizado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT unique_integracao_config UNIQUE (integracao_id)
);
