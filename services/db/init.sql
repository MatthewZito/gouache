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
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
  username VARCHAR(32) NOT NULL,
  passhash VARCHAR(16) NOT NULL
);
