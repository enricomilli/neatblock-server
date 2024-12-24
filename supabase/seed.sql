-- First, insert into auth.users
INSERT INTO auth.users (id, email, raw_user_meta_data)
VALUES
  ('d0d4ad7e-6c89-4a44-8e58-2b0d01170c2f', 'john.martinez@example.com',
    '{"full_name": "John Martinez", "avatar_url": "https://avatars.githubusercontent.com/u/1234567"}'::jsonb),
  ('f5f7c1d8-8e91-4b6a-9e67-3c4d5b6a7c8d', 'sarah.chen@example.com',
    '{"full_name": "Sarah Chen", "avatar_url": "https://avatars.githubusercontent.com/u/2345678"}'::jsonb),
  ('a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'michael.weber@example.com',
    '{"full_name": "Michael Weber", "avatar_url": "https://avatars.githubusercontent.com/u/3456789"}'::jsonb);

-- Then update public.users with additional information
UPDATE public.users
SET
  billing_address = '{"street": "123 Tech Lane", "city": "San Francisco", "state": "CA", "postal_code": "94105", "country": "US"}'::jsonb,
  payment_method = '{"card_type": "visa", "last4": "4242"}'::jsonb
WHERE id = 'd0d4ad7e-6c89-4a44-8e58-2b0d01170c2f';

UPDATE public.users
SET
  billing_address = '{"street": "456 Innovation Dr", "city": "Austin", "state": "TX", "postal_code": "78701", "country": "US"}'::jsonb,
  payment_method = '{"card_type": "mastercard", "last4": "8888"}'::jsonb
WHERE id = 'f5f7c1d8-8e91-4b6a-9e67-3c4d5b6a7c8d';

UPDATE public.users
SET
  billing_address = '{"street": "789 Crypto Ave", "city": "Miami", "state": "FL", "postal_code": "33101", "country": "US"}'::jsonb,
  payment_method = '{"card_type": "amex", "last4": "1234"}'::jsonb
WHERE id = 'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d';

-- Insert fake pools data
INSERT INTO pools (
    id, pool_url, status, user_id, name, bought_hashrate, highest_achieve_hashrate,
    total_btc_payout, total_btc_mined, total_btc_sold, total_btc_expensed,
    total_usd_gain, total_usd_expensed, revenue_share, uptime_percent,
    initial_investment, coc, roi, btc_price_at_deployment, deployment_date
)
VALUES
    ('b1d2e3f4-5a6b-4a44-8e58-2b0d01170c2f', 'https://pool1.mining.com', 'active',
    'd0d4ad7e-6c89-4a44-8e58-2b0d01170c2f', 'Alpha Pool', 150.5, 165.2,
    2.5, 3.0, 1.2, 0.5, 75000, 15000, 0.85, 98.5,
    50000, 0.12, 0.25, 45000, '2023-01-15'),

    ('c2d3e4f5-6b7c-4a44-8e58-3c4d5b6a7c8d', 'https://pool2.mining.com', 'active',
    'f5f7c1d8-8e91-4b6a-9e67-3c4d5b6a7c8d', 'Beta Pool', 200.0, 220.5,
    3.8, 4.2, 2.1, 0.8, 115000, 25000, 0.82, 97.8,
    75000, 0.15, 0.32, 42000, '2023-02-20'),

    ('d3e4f5f6-7c8d-4a44-8e58-4d5e6f7a8b9c', 'https://pool3.mining.com', 'pending',
    'a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d', 'Gamma Pool', 175.5, 185.0,
    1.8, 2.2, 1.0, 0.3, 55000, 12000, 0.88, 99.1,
    40000, 0.10, 0.28, 48000, '2023-03-10');

-- Insert fake pool_rewards data
INSERT INTO pool_rewards (
    id, pool_id, date, hashrate, btc_reward, btc_tx_fee,
    total, payout, created_at, updated_at
)
VALUES
    ('e4f5f6f7-8d9e-4a44-8e58-5e6f7a8b9c0d',
    'b1d2e3f4-5a6b-4a44-8e58-2b0d01170c2f',
    '2023-04-01', 162.5, 0.15, 0.002, 0.152, 0.145,
    '2023-04-01 12:00:00', '2023-04-01 12:00:00'),

    ('f5f6f7f8-9e0f-4a44-8e58-6f7a8b9c0d1e',
    'b1d2e3f4-5a6b-4a44-8e58-2b0d01170c2f',
    '2023-04-02', 164.8, 0.18, 0.003, 0.183, 0.175,
    '2023-04-02 12:00:00', '2023-04-02 12:00:00'),

    ('a6b7c8d9-0f1a-4a44-8e58-7a8b9c0d1e2f',
    'c2d3e4f5-6b7c-4a44-8e58-3c4d5b6a7c8d',
    '2023-04-01', 218.2, 0.22, 0.004, 0.224, 0.215,
    '2023-04-01 12:00:00', '2023-04-01 12:00:00'),

    ('b7c8d9e0-1a2b-4a44-8e58-8b9c0d1e2f3a',
    'c2d3e4f5-6b7c-4a44-8e58-3c4d5b6a7c8d',
    '2023-04-02', 220.1, 0.25, 0.003, 0.253, 0.242,
    '2023-04-02 12:00:00', '2023-04-02 12:00:00'),

    ('c8d9e0f1-2b3c-4a44-8e58-9c0d1e2f3a4b',
    'd3e4f5f6-7c8d-4a44-8e58-4d5e6f7a8b9c',
    '2023-04-01', 183.5, 0.12, 0.002, 0.122, 0.115,
    '2023-04-01 12:00:00', '2023-04-01 12:00:00');
