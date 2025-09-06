-- Insert initial data
-- Create administrator (password will be hashed in code)
INSERT INTO users (id, login, name, email, password)
VALUES (1, 'admin', 'Administrator', 'admin@localhost', 'admin')
ON CONFLICT (login) DO NOTHING;

-- Create test tags
INSERT INTO tags (id, name, slug, author_id)
VALUES (1, 'Go', 'go', 1),
       (2, 'GORM', 'gorm', 1),
       (3, 'Web Development', 'web-development', 1),
       (4, 'Tutorial', 'tutorial', 1),
       (5, 'Database', 'database', 1)
ON CONFLICT DO NOTHING;

-- Create test post
INSERT INTO posts (title, slug, description, content, author_id)
VALUES ('Welcome to OxygenBlog', 'welcome-to-oxygenblog', 'First post in the blog',
        'This is the first post in our blog created using Go and GORM.', 1)
ON CONFLICT DO NOTHING;

-- Link post to tags
INSERT INTO posts_tags (post_id, tag_id)
VALUES (1, 1), -- Go
       (1, 2), -- GORM
       (1, 3), -- Web Development
       (1, 4)  -- Tutorial
ON CONFLICT DO NOTHING;