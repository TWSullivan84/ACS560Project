CREATE TABLE IF NOT EXISTS Users (
	userID 			INT NOT NULL AUTO_INCREMENT,
	userName 		VARCHAR(20) NOT NULL UNIQUE,
	userPassword 	VARCHAR(20) NOT NULL,
	PRIMARY KEY(userID)
);

CREATE TABLE IF NOT EXISTS PlayerAchievements (
	userID			INT NOT NULL,
	totalWins		INT NOT NULL,
	totalLosses		INT NOT NULL,
	totalMatches	INT NOT NULL,
	bombsDropped	INT NOT NULL,
	PRIMARY KEY(userID),
	CONSTRAINT fk_PAUserID FOREIGN KEY (userID)  REFERENCES Users(userID)
);

CREATE TABLE IF NOT EXISTS MatchHistory (
	matchID			INT NOT NULL AUTO_INCREMENT,
	userID1			INT NOT NULL,
	userID2			INT NOT NULL,
	winnerID		INT NOT NULL,
	matchLength		TIME NOT NULL,
	PRIMARY KEY(matchID),
	CONSTRAINT fk_MH1_userID FOREIGN KEY (userID1) REFERENCES Users(userID),
	CONSTRAINT fk_MH2_userID FOREIGN KEY (userID2) REFERENCES Users(userID)
);