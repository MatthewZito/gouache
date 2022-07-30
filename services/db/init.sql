-- CREATE USER testuser WITH PASSWORD 'test';

-- CREATE DATABASE test;

CREATE extension IF NOT EXISTS "uuid-ossp";

CREATE TABLE resource_record (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE,
  tags VARCHAR(255)[],
  title VARCHAR(255) NOT NULL
);


GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO testuser;
