-- Time Zone
SET timezone = 'UTC';

-- Tabla catalogo para Estados de Tenants
CREATE TABLE IF NOT EXISTS tenant_status (
    id INT PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    description TEXT
);

-- Tabla de Tenants
CREATE TABLE IF NOT EXISTS tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    subdomain VARCHAR(100) NOT NULL UNIQUE,
    status_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Tabla de contacto de Tenant
CREATE TABLE IF NOT EXISTS tenant_contact_info (
    tenant_id UUID NOT NULL,
    nit VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    principal_address TEXT,
    secondary_address TEXT,
    city VARCHAR(100),
    departament VARCHAR(100)
);

-- Tabla de informacion de base de datos
CREATE TABLE IF NOT EXISTS tenant_database_info (
    tenant_id UUID NOT NULL,
    host VARCHAR(255),
    port INT,
    db_name VARCHAR(100),
    db_user VARCHAR(100),
    db_password TEXT,
    ssl_mode VARCHAR(20) DEFAULT 'disable',
    settings JSONB DEFAULT '{}'
);

-- Relaciones
ALTER TABLE tenants ADD CONSTRAINT fk_tenants_status
FOREIGN KEY (status_id) REFERENCES tenant_status(id)
ON DELETE RESTRICT;

ALTER TABLE tenant_contact_info ADD CONSTRAINT fk_tenant
FOREIGN KEY (tenant_id) REFERENCES tenants(id)
ON DELETE RESTRICT;

ALTER TABLE tenant_database_info ADD CONSTRAINT fk_tenant
FOREIGN KEY (tenant_id) REFERENCES tenants(id)
ON DELETE RESTRICT;
