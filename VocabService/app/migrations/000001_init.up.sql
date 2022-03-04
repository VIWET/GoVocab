CREATE TABLE lists (
    id SERIAL,
    UserID INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE words (
    id SERIAL,
    test VARCHAR(50),
    PRIMARY KEY (id)
);

CREATE TABLE words_lists_relation (
    list_id INT NOT NULL,
    word_id INT NOT NULL,
  	FOREIGN KEY (list_id) REFERENCES lists(id), 
    FOREIGN KEY (word_id) REFERENCES words(id),
    UNIQUE (list_id, word_id)
);

CREATE TABLE synonyms (
    word_id INT NOT NULL,
    synonym_id INT NOT NULL,
    FOREIGN KEY (word_id) REFERENCES words(id), 
    FOREIGN KEY (synonym_id) REFERENCES words(id),
    UNIQUE (word_id, synonym_id)
);

CREATE TABLE meanings (
    id SERIAL,
    word_id INT NOT NULL,
    type_of_speech VARCHAR(10) NOT NULL,
    description VARCHAR(100) NOT NULL,
    translation VARCHAR(100) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (word_id) REFERENCES words(id)
);

CREATE TABLE use_cases (
    id SERIAL,
    meaning_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (meaning_id) REFERENCES meanings(id)
);