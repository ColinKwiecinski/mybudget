DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    ID INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
    Name VARCHAR(255),
    Email VARCHAR(255),
    Contact_Num VARCHAR(20)
);

CREATE TABLE Auth (
    ID INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
    Hash VARCHAR(MAX),
    User_ID INT,
    FOREIGN KEY ([User_ID]) REFERENCES Users(ID)
);

-- CREATE TABLE Transaction_Type (
--     ID INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
--     Transaction_Type_Name VARCHAR(255),
--     Description VARCHAR(MAX)
-- );

CREATE TABLE Transactions (
    ID INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
    User_ID INT,
    Transaction_Name VARCHAR(255),
    Memo VARCHAR(255),
    Transaction_Date VARCHAR(255) NOT NULL,
    Amount DECIMAL(8,2),
    -- Transaction_Type_ID INT,
    Transaction_Type VARCHAR(255),
    FOREIGN KEY ([User_ID]) REFERENCES Users(ID),
    -- FOREIGN KEY ([Transaction_Type_ID]) REFERENCES Transaction_Type(ID)
);