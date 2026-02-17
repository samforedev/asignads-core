-- Time Zone
SET timezone = 'UTC';

-- Tabla catalogo para Estatus de Tenants
CREATE TABLE IF NOT EXISTS tenant_status(
    id INT PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    description TEXT
);

-- Tabla de Tenants
CREATE TABLE IF NOT EXISTS tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    subdomain VARCHAR(100) UNIQUE NOT NULL,
    db_dsn TEXT NOT NULL,
    status_id INT NOT NULL DEFAULT 1 REFERENCES tenant_status(id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO tenant_status (id, name, description)
VALUES
(1, 'ACTIVE', 'Tenant operativo'),
(2, 'SUSPENDED', 'Tenant suspendido por falta de pago'),
(3, 'PROVISIONING', 'Tenant en proceso de creación'),
(4, 'DELETED', 'Tenant marcado para eliminación')
ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, description = EXCLUDED.description;


INSERT INTO tenants (name, subdomain, db_dsn)
VALUES
('Agencia Prueba Uno', 'pruebauno', 'host=localhost port=5433 user=asigna_admin password=admin_pass dbname=asigna_db sslmode=disable' ),
('Agencia Prueba Dos', 'pruebados', 'host=localhost port=5434 user=asigna_admin password=admin_pass dbname=asigna_db sslmode=disable');