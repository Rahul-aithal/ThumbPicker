
CREATE TABLE video   (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    src TEXT NOT NULL,
    thumbnails_count INT NOT NULL,
    dur TEXT NOT NULL,
    thumbnail UUID REFERENCES thumbnails (id)
);

create table thumbnails (
    id uuid primary key default gen_random_uuid(),
    src text not null,
    video uuid references video (id) not null,
    timestamp text  not null,
    idx int
);
