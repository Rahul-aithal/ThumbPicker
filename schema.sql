
CREATE TABLE Video  (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    src TEXT NOT NULL,
    thumbnails_count INT NOT NULL,
    dur TEXT NOT NULL,
    thumbnail UUID REFERENCES Thumbnails (id)
);

CREATE TABLE Thumbnails(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    src TEXT NOT NULL,
    video UUID REFERENCES Videos (id) NOT NULL,
    timestamp TEXT  NOT NULL,
    idx INT
);
