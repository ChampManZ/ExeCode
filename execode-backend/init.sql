CREATE DATABASE execode;

CREATE TABLE AccountDetail(
  userId SERIAL PRIMARY KEY,
  userName varchar(20) NOT NULL UNIQUE,
  firstName varchar(25) NOT NULL,
  lastName varchar(25) NOT NULL,
  email varchar(80) NOT NULL UNIQUE,
  hashedPassword varchar(200),
  salt varchar(25)
);

CREATE TABLE Class (
  classId SERIAL PRIMARY KEY,
  className varchar(20) NOT NULL,
  classDescription varchar(1000)
);

CREATE TABLE ClassLecturer (
  lecturer int NOT NULL,
  class int NOT NULL,
  PRIMARY KEY (lecturer, class),
  FOREIGN KEY (lecturer) REFERENCES AccountDetail(userId),
  FOREIGN KEY (class) REFERENCES Class(classId)
);

CREATE TABLE Lecture (
  lectureId SERIAL PRIMARY KEY,
  class int NOT NULL,
  lectureName varchar(60) NOT NULL,
  lectureDescription varchar(500),
  createdOn date NOT NULL,
  lastEditedOn date NOT NULL,
  FOREIGN KEY (class) REFERENCES Class(classId)
);

CREATE TABLE LectureContent (
  lectureId int PRIMARY KEY,
  content text NOT NULL,
  FOREIGN KEY (lectureId) REFERENCES LectureNote(lectureId)
);

CREATE TABLE ProblemStatement (
  problemId SERIAL PRIMARY KEY,
  class int NOT NULL,
  problemName varchar(80) NOT NULL,
  FOREIGN KEY (class) REFERENCES Class(classId)
);

CREATE TABLE ProblemStatementContent (
  problemId int PRIMARY KEY,
  content text NOT NULL,
  FOREIGN KEY (problemId) REFERENCES ProblemStatement(problemId)
);

CREATE TABLE TestCase (
  problemId int,
  testCaseId int,
  caseInput text NOT NULL,
  caseOutput text NOT NULL,
  PRIMARY KEY (problemId, testCaseId),
  FOREIGN KEY (problemId) REFERENCES ProblemStatement(problemId)
);
