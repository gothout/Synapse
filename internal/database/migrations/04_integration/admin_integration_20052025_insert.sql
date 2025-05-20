-- Marcas
INSERT INTO admin_integracao_marcas (nome) VALUES 
('Google'), 
('Chatvolt');

-- Integrações específicas
INSERT INTO admin_integracoes (nome, marca_id) VALUES 
('Vision AI', (SELECT id FROM admin_integracao_marcas WHERE nome = 'Google')),
('Agent', (SELECT id FROM admin_integracao_marcas WHERE nome = 'Chatvolt'));
