CREATE INDEX idx_movies_title ON movies USING gin(to_tsvector('english', title));
