CREATE TABLE resource (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
  created_at timestamp with time zone DEFAULT now() NOT NULL,
  updated_at timestamp with time zone,
  tags VARCHAR(255)[],
  title VARCHAR(255) NOT NULL,
)
