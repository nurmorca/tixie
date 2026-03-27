CREATE TABLE IF NOT EXISTS EVENTS (
    EV_ID BIGSERIAL PRIMARY KEY,
    EV_NAME VARCHAR(255) NOT NULL,
    EV_HOST VARCHAR(255) NOT NULL,
    EV_DESCRIPTION TEXT,
    EV_VENUE VARCHAR(255) NOT NULL,
    EV_DATETIME TIMESTAMP NOT NULL,
    EV_CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS VENUE (
    VE_ID BIGSERIAL PRIMARY KEY,
    VE_NAME VARCHAR(255) NOT NULL UNIQUE,
    VE_CAPACITY INT,
    VE_CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS SEAT (
    se_id BIGSERIAL PRIMARY KEY,
    se_venue_id INT REFERENCES VENUE(VE_ID),
    se_seat_row VARCHAR(10),    
    se_seat_number INT,         
    se_seat_type VARCHAR(50),
    se_created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS EVENT_SEAT (
    ES_ID BIGSERIAL PRIMARY KEY,
    ES_EVENT_ID INT REFERENCES events(ev_id),
    ES_SEAT_ID INT REFERENCES seat(se_id),
    ES_PRICE DECIMAL(10, 2) NOT NULL,
    ES_STATUS VARCHAR(20) DEFAULT 'available'
);

CREATE TABLE IF NOT EXISTS TICKET (
    TI_ID BIGSERIAL PRIMARY KEY,
    TI_EVENT_ID BIGINT NOT NULL REFERENCES events(EV_ID) ON DELETE CASCADE,
    TI_SEAT_ID BIGINT NOT NULL REFERENCES event_seat(ES_ID),
    TI_USER_ID BIGINT NOT NULL,
    TI_CREATED_AT TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(TI_EVENT_ID, TI_SEAT_ID)
);

ALTER TABLE EVENT_SEAT ADD CONSTRAINT ES_STATUS CHECK (ES_STATUS IN ('available', 'locked', 'sold'));
CREATE INDEX idx_tickets_event_id ON ticket(TI_EVENT_ID);
CREATE INDEX idx_events_date ON events(EV_DATETIME);

-- pre-inserted venue and seat data
INSERT INTO events (ev_host, ev_name, ev_venue, ev_description, ev_created_at, ev_datetime)
VALUES (
    'daft punk',
    'istanbul music fest',
    'istanbul arena',
    'a night of live music featuring local and international artists.',
    NOW(),
    '2026-06-15 20:00:00'
);
INSERT INTO venue(VE_NAME, VE_CAPACITY) VALUES ('istanbul arena', 44);
INSERT INTO seat (se_venue_id, se_seat_row, se_seat_number, se_seat_type) 
SELECT 
    1,
    row_letter,
    seat_num,
    CASE WHEN row_letter IN ('A', 'B') THEN 'VIP' ELSE 'STANDARD' END
FROM 
    unnest(ARRAY['A','B','C','D']) AS row_letter,
    generate_series(1, 11) AS seat_num;

INSERT INTO event_seat (ES_EVENT_ID, ES_SEAT_ID, ES_STATUS, ES_PRICE)
SELECT 
    1,
    se_id,
    'available',
    CASE WHEN se_seat_type = 'VIP' THEN 1500.00 ELSE 750.00 END
FROM seat
WHERE se_venue_id = 1;