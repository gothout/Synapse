-- 🚀 Empresa: Wonit Joinville
INSERT INTO admin_enterprise (id, nome, cnpj, created_at, update_at)
VALUES (
    1,
    'Wonit Joinvile',
    '81385593000153',
    '2025-05-16 17:11:27.020264',
    '2025-05-16 20:11:27.38049'
)
ON CONFLICT (id) DO NOTHING;

-- 👤 Usuário admin com senha já criptografada (bcrypt)
INSERT INTO admin_user (
    id,
    nome,
    email,
    senha,
    numero,
    token,
    rule_id,
    enterprise_id,
    created_at,
    updated_at
)
VALUES (
    1,
    'admin',
    'admin@wonit.com.br',
    '$2a$10$6uzsLGQMQdWbTRwv792Y/Oz2Ht6.XCw2T9RCnXM3Q1Fpb6z78XP3a',
    '5547997845532',
    '',
    1,  -- rule_id = 1 (admin)
    1,  -- enterprise_id = 1 (Wonit Joinvile)
    '2025-05-16 20:11:29.630681',
    '2025-05-20 13:11:19.316702'
)
ON CONFLICT (id) DO NOTHING;
