DROP TABLE IF EXISTS invoice, brokie, user_info, property_info;

CREATE TABLE user_info (
    user_type VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) PRIMARY KEY,
    phone VARCHAR(255),
    hash_password VARCHAR(255),
    one_time_code VARCHAR(255) NOT NULL, 
)

CREATE TABLE property_info (
    property_id SERIAL PRIMARY KEY,
    property_address VARCHAR(255),
)

CREATE TABLE brokie (
    payment_day INT,
    rent_rate INT,
    FOREIGN KEY (property_id) references property_info(property_id),
    active VARCHAR(1) NOT NULL, -- Y or N
    FOREIGN KEY (email) references user_info(email)
)

CREATE TABLE invoice (
    due_date DATE,
    amount INT,
    payment_status VARCHAR(255)
    payment_type VARCHAR(255)
    payment_id SERIAL PRIMARY KEY,
    FOREIGN KEY (email) references user_info(email)

)
