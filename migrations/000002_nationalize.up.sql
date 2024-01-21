CREATE TABLE IF NOT EXISTS nationalize(
    id uuid PRIMARY KEY,
    country_id varchar,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE,
    probability FLOAT
)




