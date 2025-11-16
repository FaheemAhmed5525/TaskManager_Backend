
-- tasks table
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Index for performance
CREATE INDEX IF NOT EXISTS idx_task_created_at ON tasks(created_at);
CREATE INDEX IF NOT EXISTS idx_task_completed ON tasks(completed);

-- Insert sample data
INSERT INTO tasks (title, completed) VALUES
('Day 1: Concurrency Check', TRUE),
('Day 2: Design Task API with in-memory setup', TRUE),
('Day 3: Add database to task API, handle errors', FALSE),
('Day 4: Authentication and JWT tokens', FALSE),
('Day 5: Write unit tests', FALSE)
ON CONFLICT DO NOTHING;

