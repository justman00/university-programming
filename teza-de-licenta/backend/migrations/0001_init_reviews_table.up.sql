CREATE TABLE reviews (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    rating INTEGER,
    source TEXT NOT NULL,
    review TEXT NOT NULL,
    analysis JSONB NOT NULL,
    original_payload JSONB NOT NULL,
    review_created_at TIMESTAMPTZ NOT NULL,
    review_updated_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
