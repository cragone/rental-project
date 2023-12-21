DROP TABLE IF EXISTS invoice, brokie, user_info, property_info, session;

CREATE TABLE user_info (
    user_type VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) PRIMARY KEY,
    phone VARCHAR(255),
    hash_password VARCHAR(255),
    one_time_code VARCHAR(255) NOT NULL
);

CREATE TABLE session (
    email VARCHAR(255) REFERENCES user_info(email),
    session_token VARCHAR(255) PRIMARY KEY 
);

CREATE TABLE property_info (
    property_id SERIAL PRIMARY KEY,
    property_address VARCHAR(255)
);

CREATE TABLE brokie (
    b_id SERIAL PRIMARY KEY,
    payment_day INT,
    rent_rate INT,
    property_id INT REFERENCES property_info(property_id),
    active VARCHAR(1) NOT NULL, -- Y or N
    email VARCHAR(255) REFERENCES user_info(email)
);

CREATE TABLE invoice (
    due_date DATE,
    amount INT,
    payment_status VARCHAR(255),
    paypal_id VARCHAR(255),
    payment_type VARCHAR(255),
    payment_id VARCHAR(255) PRIMARY KEY,
    tennant_id INT REFERENCES brokie(b_id) 
);