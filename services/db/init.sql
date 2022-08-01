-- Add support for UUID.
CREATE extension IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS resource_record (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE,
  tags VARCHAR(255)[],
  title VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_record (
  username VARCHAR(32) PRIMARY KEY,
  id UUID DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  passhash VARCHAR(255) NOT NULL
);
