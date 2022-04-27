CREATE TYPE dna_type as ENUM ('Human', 'Mutant');

CREATE TABLE dna (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "dna" VARCHAR NOT NULL,
    "dna_type" dna_type NOT NULL
);

CREATE INDEX idx_dna_type
ON dna(dna_type);