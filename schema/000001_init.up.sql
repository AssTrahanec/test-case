CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       username TEXT UNIQUE NOT NULL,
                       password_hash TEXT NOT NULL,
                       balance INTEGER DEFAULT 0,
                       referrer_id UUID REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE tasks (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       name TEXT NOT NULL,
                       description TEXT NOT NULL,
                       reward_points INTEGER NOT NULL
);

CREATE TABLE user_tasks (
                            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                            task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                            completed_at TIMESTAMP NOT NULL
);
INSERT INTO tasks (name, description, reward_points)
VALUES
    ('Join Telegram', 'Join our official Telegram channel to earn points.', 50),
    ('Follow on Twitter', 'Follow us on Twitter to stay updated and earn points.', 30),
    ('Complete Profile', 'Fill in your profile details completely.', 20),
    ('Refer a Friend', 'Refer a friend to join the platform.', 100),
    ('Submit Feedback', 'Submit your feedback about our platform.', 10),
    ('Watch Tutorial', 'Watch our tutorial video to learn how to use the platform.', 25),
    ('Daily Login', 'Login to your account daily to earn points.', 5),
    ('Participate in Survey', 'Participate in a survey to share your thoughts and earn points.', 40),
    ('Write a Review', 'Write a review about our platform.', 15),
    ('Invite to Webinar', 'Invite someone to attend our upcoming webinar.', 70);
INSERT INTO users(username, password_hash)
VALUES
    ('user12', '61736761736761683132346a766a736661ea3a56c6a1f0272ec675c598699add1d43e4cf12'),
    ('user1', '61736761736761683132346a766a736661ea3a56c6a1f0272ec675c598699add1d43e4cf12');