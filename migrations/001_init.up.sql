CREATE TABLE IF NOT EXISTS user_profiles (
                                             user_id BIGINT PRIMARY KEY,
                                             name TEXT NOT NULL,
                                             birth_date TEXT,
                                             zodiac_sign TEXT,
                                             telegram_username TEXT
);

CREATE TABLE IF NOT EXISTS mbti_results (
    user_id BIGINT PRIMARY KEY REFERENCES user_profiles(user_id) ON DELETE CASCADE,
    mbti TEXT NOT NULL
    );
