DROP TABLE IF EXISTS Accounts CASCADE;
CREATE TABLE IF NOT EXISTS Accounts(
	login 			VARCHAR(64)		PRIMARY KEY,
	role 			VARCHAR(15) 	NOT NULL,
	salt 			VARCHAR(64)		NOT NULL,
	hashed_password	VARCHAR(64)		NOT NULL
);

DROP TABLE IF EXISTS Ingredients CASCADE;
CREATE TABLE IF NOT EXISTS Ingredients(
	title 			VARCHAR(64)		PRIMARY KEY
);

DROP TABLE IF EXISTS Categories CASCADE;
CREATE TABLE IF NOT EXISTS Categories (
	title 			VARCHAR(64)		PRIMARY KEY
);

DROP TABLE IF EXISTS Recipes CASCADE;
CREATE TABLE IF NOT EXISTS Recipes (
	id				SERIAL 			PRIMARY KEY,
	author			VARCHAR(64)		REFERENCES Accounts(login) NOT NULL,
	title 			VARCHAR(128)	NOT NULL,
	created_at		TIMESTAMP		NOT NULL DEFAULT CURRENT_TIMESTAMP,
	description		TEXT			NOT NULL,
	duration		INT,
	portion_num		INT
);

DROP TABLE IF EXISTS Steps CASCADE;
CREATE TABLE IF NOT EXISTS Steps (
	recipe			INT 			REFERENCES Recipes(id) NOT NULL,
	num				SERIAL			NOT NULL,
	description		TEXT			NOT NULL,
	title 			VARCHAR(64),
	PRIMARY KEY (recipe, num)
);

DROP TABLE IF EXISTS Grades CASCADE;
CREATE TABLE IF NOT EXISTS Grades (
	grades_recipe	INT 			REFERENCES Recipes(id) NOT NULL,
	account_login	VARCHAR(64)		REFERENCES Accounts(login) NOT NULL,
	PRIMARY KEY (recipe_id, account_login)
);

DROP TABLE IF EXISTS Recipe_Category CASCADE;
CREATE TABLE IF NOT EXISTS Recipe_Category (
	recipe_id		INT 			REFERENCES Recipes(id) NOT NULL,
	category_title	VARCHAR(64)		REFERENCES Categories(title) NOT NULL,
	PRIMARY KEY (recipe_id, category_title)
);

DROP TABLE IF EXISTS Recipe_Ingredient CASCADE;
CREATE TABLE IF NOT EXISTS Recipe_Ingredient (
	recipe			INT 			REFERENCES Recipes(id) NOT NULL,
	item			VARCHAR(64)		REFERENCES Ingredients(title) NOT NULL,
	amount			VARCHAR(64),
	PRIMARY KEY (recipe, item)
);
