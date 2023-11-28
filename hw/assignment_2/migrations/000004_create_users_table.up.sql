CREATE TABLE IF NOT EXISTS users (
                                     id bigserial PRIMARY KEY,
                                     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
                                     name text NOT NULL,
                                     email citext UNIQUE NOT NULL,
                                     password_hash bytea NOT NULL,
                                     activated bool NOT NULL,
                                     balance numeric NOT NULL DEFAULT 0, -- Add this line for the account balance
                                     version integer NOT NULL DEFAULT 1
);