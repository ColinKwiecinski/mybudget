DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Email VARCHAR(255),
    Hash VARCHAR(10000),
    Contact_Num VARCHAR(20)
);

CREATE TABLE Transactions (
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    User_ID INT,
    Transaction_Name VARCHAR(255),
    Memo VARCHAR(255),
    Transaction_Date VARCHAR(255) NOT NULL,
    Amount DECIMAL(8,2),
    Transaction_Type VARCHAR(255),
    CONSTRAINT fk_userid2 FOREIGN KEY (User_ID)
    REFERENCES Users(ID)
);