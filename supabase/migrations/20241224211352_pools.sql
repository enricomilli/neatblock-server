
-- Create enum for pool status
CREATE TYPE pool_status AS ENUM (
    'active',
    'inactive',
    'pending'
);

CREATE TABLE pools (
    id uuid PRIMARY KEY NOT NULL,
    pool_url text not null,
    status pool_status NOT NULL DEFAULT 'pending',
    user_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    name text not null,
    bought_hashrate numeric,
    highest_achieve_hashrate numeric NOT NULL,
    total_btc_payout numeric NOT NULL,
    total_btc_mined numeric NOT NULL,
    total_btc_sold numeric,
    total_btc_expensed numeric,
    total_usd_gain numeric,
    total_usd_expensed numeric,
    revenue_share numeric,
    uptime_percent numeric NOT NULL,
    initial_investment numeric,
    coc numeric,
    roi numeric,
    btc_price_at_deployment numeric,
    deployment_date timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pool_rewards (
    id uuid PRIMARY KEY,
    pool_id uuid REFERENCES pools(id) ON DELETE CASCADE,
    date timestamp without time zone,
    hashrate numeric NOT NULL,
    btc_reward numeric NOT NULL,
    btc_tx_fee numeric NOT NULL,
    total numeric NOT NULL,
    payout numeric NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

-- Create index on user_id for faster lookups
CREATE INDEX pools_user_id_idx ON pools(user_id);
