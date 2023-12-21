CREATE TABLE shows (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    rate VARCHAR(10),
    year VARCHAR(4),
    description TEXT,
    genres VARCHAR(255),
    image_url VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ DEFAULT current_timestamp
);
