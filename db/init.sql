-- Таблицы с UUID вместо SERIAL
CREATE TABLE users (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUID вместо SERIAL
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    role TEXT CHECK (role IN ('client', 'moderator', 'pvz_employee')) NOT NULL
);

CREATE TABLE pvz (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUID вместо SERIAL
    registration_date TIMESTAMP DEFAULT NOW(),
    city TEXT CHECK (city IN ('Москва', 'Санкт-Петербург', 'Казань')) NOT NULL,
    created_by UUID REFERENCES users(uuid) ON DELETE SET NULL  -- Используем UUID
);

CREATE TABLE pvz_employees (
    user_id UUID REFERENCES users(uuid) ON DELETE CASCADE,  -- UUID
    pvz_id UUID REFERENCES pvz(uuid) ON DELETE CASCADE,  -- UUID
    PRIMARY KEY (user_id, pvz_id)
);

CREATE TABLE receptions (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUID вместо SERIAL
    execution_date TIMESTAMP DEFAULT NOW(),
    pvz_id UUID REFERENCES pvz(uuid) ON DELETE CASCADE,  -- UUID
    status TEXT CHECK (status IN ('in_progress', 'closed')) NOT NULL
);

CREATE TABLE products (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- UUID вместо SERIAL
    reception_id UUID REFERENCES receptions(uuid) ON DELETE CASCADE,  -- UUID
    reception_date TIMESTAMP DEFAULT NOW(),
    type TEXT CHECK (type IN ('электроника', 'одежда', 'обувь')) NOT NULL
);
