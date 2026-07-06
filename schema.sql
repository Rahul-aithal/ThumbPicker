-- video --
CREATE TABLE IF NOT EXISTS video (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  src TEXT NOT NULL,
  thumbnails_count INT NOT NULL,
  dur TEXT NOT NULL
);

-- thumbnails --
CREATE TABLE IF NOT EXISTS thumbnails (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
  src TEXT NOT NULL,
  time_stamp TEXT NOT NULL,
  idx INT,
  video UUID REFERENCES video (id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL
);

--Alter for adding thumbnail--
ALTER TABLE video ADD thumbnail UUID REFERENCES thumbnails (id) ON DELETE CASCADE ON UPDATE CASCADE;
