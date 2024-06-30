CREATE INDEX idx_shows_title ON shows USING gin(to_tsvector('english', title));
