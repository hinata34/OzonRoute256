CREATE TABLE assessors
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    skills TEXT[],
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE managers
(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE vacancies
(
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    skills TEXT[],
    salary MONEY,
   -- manager_id BIGSERIAL NOT NULL, -- у вакансии может быть только один менеджер
    closed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE reactions
(
    vacancy_id BIGSERIAL NOT NULL,
    assessor_id BIGSERIAL NOT NULL,
    cv TEXT NOT NULL,
    description TEXT,
    salary MONEY,
    result TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE(vacancy_id, assessor_id)
);

CREATE TABLE managers_vacancy 
(
    vacancy_id BIGSERIAL NOT NULL,
    manager_id BIGSERIAL NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE(vacancy_id, manager_id)
);

CREATE INDEX reactions_vacancy_id_assessor_id ON reactions(vacancy_id, assessor_id);

CREATE INDEX managers_vacancy_vacancy_id_manager_id ON managers_vacancy(vacancy_id, manager_id);
