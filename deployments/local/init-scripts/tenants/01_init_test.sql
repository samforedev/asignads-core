SET timezone = 'UTC';

-- Tabla de Pruebas
CREATE TABLE IF NOT EXISTS test (
    id SERIAL PRIMARY KEY,
    tenant_name VARCHAR(100) NOT NULL,
    message TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);