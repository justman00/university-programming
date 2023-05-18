ALTER table bookings ADD COLUMN type TEXT NOT NULL DEFAULT 'any table';

ALTER table bookings ADD COLUMN table_number INTEGER NOT NULL DEFAULT 0;
