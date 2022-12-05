CREATE SCHEMA `godemodb`;
CREATE TABLE QuizData
(
    ID bigint auto_increment,
    Question varchar(255) NOT NULL,
    Answer varchar(255) NOT NULL,
    PRIMARY KEY (ID)
);

INSERT INTO QuizData(Question, Answer) VALUES ('low-key', '「控え目、地味」「秘密」「ちょっと、なんとなく」'), ('sloppy', '雑な');