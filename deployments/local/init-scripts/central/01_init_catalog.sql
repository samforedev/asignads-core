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
    nit VARCHAR(20) UNIQUE,

    contact_name VARCHAR(100) NOT NULL,
    contact_email VARCHAR(150) NOT NULL,
    contact_phone VARCHAR(20) NOT NULL,

    address TEXT,
    city VARCHAR(100),
    department VARCHAR(100),

    db_host VARCHAR(255),
    db_port INT,
    db_name VARCHAR(100),
    db_user VARCHAR(100),
    db_password TEXT, -- ENCRIPTADO
    db_ssl_mode VARCHAR(20) DEFAULT 'disable',

    settings JSONB DEFAULT '{}',
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

INSERT INTO tenants (
    name,
    subdomain, 
    nit, 
    contact_name, 
    contact_email, 
    contact_phone, 
    address, 
    city, 
    department, 
    db_host, 
    db_port, 
    db_name, 
    db_user, 
    db_password, 
    db_ssl_mode, 
    settings, 
    status_id
) VALUES
(
    'Agencia Prueba Uno', 
    'pruebauno', 
    '123456789', 
    'Metacho Samuel', 
    'samuelf.dev@outlook.com', 
    '+573013317096', 
    'Calle 1 # 2-3', 
    'Bogota', 
    'Bogota', 
    'localhost', 
    5433, 
    'asigna_db', 
    'asigna_admin', 
    'admin_pass', 
    'disable', 
    '{}', 
    1
);
