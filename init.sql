CREATE TABLE IF NOT EXISTS ru_en (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(50) UNIQUE,
                       translation VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS reports (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

/*
INSERT INTO ru_en (title, translation) VALUES
                                           ('Привет', 'Hello'),
                                           ('Мир', 'World'),
                                           ('Книга', 'Book'),
                                           ('Стол', 'Table'),
                                           ('Яблоко', 'Apple'),
                                           ('Солнце', 'Sun'),
                                           ('Вода', 'Water'),
                                           ('Дом', 'House'),
                                           ('Кот', 'Cat'),
                                           ('Собака', 'Dog'),
                                           ('Человек', 'Human'),
                                           ('Школа', 'School'),
                                           ('Машина', 'Car'),
                                           ('Окно', 'Window'),
                                           ('Ручка', 'Pen');
*/