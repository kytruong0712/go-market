CREATE TABLE IF NOT EXISTS public.uploaded_files
(
    id               BIGINT PRIMARY KEY,
    file_name        TEXT                     NOT NULL CHECK (file_name <> ''::TEXT),
    file_size_kb     TEXT                     NOT NULL CHECK (file_size_kb <> ''::TEXT),
    file_path        TEXT                     NOT NULL CHECK (file_path <> ''::TEXT),
    file_type        TEXT                     NOT NULL CHECK (file_type <> ''::TEXT),
    status           TEXT                     NOT NULL CHECK (status <> ''::TEXT),
    resolution       TEXT                              CHECK (resolution <> ''::TEXT),
    created_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS public.categories
(
    id                 BIGINT PRIMARY KEY,
    parent_id          BIGINT                            REFERENCES public.categories (id),
    category_name      TEXT                     NOT NULL CHECK (category_name <> ''::TEXT),
    category_code      TEXT                     NOT NULL CHECK (category_code <> ''::TEXT),
    description        TEXT                     NOT NULL CHECK (description <> ''::TEXT),
    status             TEXT                     NOT NULL CHECK (status <> ''::TEXT),
    is_navgitation     BOOLEAN,
    is_filtering       BOOLEAN,
    created_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS public.category_images
(
    id                   BIGINT PRIMARY KEY,
    category_id          BIGINT                   NOT NULL REFERENCES public.categories (id),
    uploaded_file_id     BIGINT                   NOT NULL REFERENCES public.uploaded_files (id),
    status               TEXT                     NOT NULL CHECK (status <> ''::TEXT),
    image_type           TEXT                              CHECK (image_type <> ''::TEXT),
    created_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
