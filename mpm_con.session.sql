--@block 

CREATE TABLE IF NOT EXISTS Users (
    user_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email varchar(255)  UNIQUE,
    password_hash VARCHAR(255) ,
    salt VARCHAR(255) ,
    UserName VARCHAR(255) ,
    bio TEXT,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
--@block 
CREATE INDEX email_index ON Users(email);
CREATE INDEX UserName_index ON Users(UserName);
--@block
SHOW INDEX FROM Users;

-- @block
CREATE TABLE IF NOT EXISTS Followers(
    follower_id BIGINT,
    followed_id BIGINT,
    FOREIGN KEY (follower_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (followed_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    PRIMARY KEY (followed_id, follower_id)
)
--@block
SHOW INDEX FROM Followers;


--@block
CREATE TABLE IF NOT EXISTS Meals( 
    meal_id BIGINT AUTO_INCREMENT PRIMARY KEY ,
    MealName VARCHAR(255),
    owner_id BIGINT ,
    FOREIGN KEY (owner_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    carbs FLOAT ,
    protein FLOAT ,
    fat FLOAT ,
    kcal FLOAT 
);
CREATE TABLE IF NOT EXISTS Ingedients(
    ingredient_id BIGINT AUTO_INCREMENT PRIMARY KEY ,
    off_prod_id BIGINT,
    quantity FLOAT,
    carbs FLOAT ,
    protein FLOAT ,
    fat FLOAT ,
    kcal FLOAT 
);
CREATE TABLE IF NOT EXISTS MicroNutrients(
    meal_owner_id BIGINT  PRIMARY KEY ,
    FOREIGN KEY (meal_owner_id) REFERENCES Meals(meal_id) ON DELETE CASCADE,  
    iron FLOAT ,
    zinc FLOAT ,
    magnesium FLOAT ,
    calcium FLOAT ,
    B12 FLOAT ,
    folate FLOAT ,
    C  FLOAT ,
    copper FLOAT ,
    phosphorus FLOAT ,
    potassium FLOAT ,
    sodium FLOAT ,
    A FLOAT ,
    B6 FLOAT ,
    D FLOAT ,
    niacin FLOAT ,
    riboflavin FLOAT ,
    sulfur FLOAT ,
    thimined FLOAT ,
    cloride FLOAT ,
    iodine FLOAT ,
    manganese FLOAT ,
    selenium FLOAT ,
    biotin FLOAT ,
    fluoride FLOAT 
);
CREATE TABLE IF NOT EXISTS Carbohydrates(
    owner_id BIGINT ,
    FOREIGN KEY (owner_id) REFERENCES Meals(meal_id) ON DELETE CASCADE, 
    PRIMARY KEY (owner_id),
    Glucose FLOAT,
    Sucrose FLOAT,
    ribose FLOAT,
    amylose FLOAT,
    amylopectin FLOAT,
    maltose FLOAT,
    galactose FLOAT,
    fructose FLOAT,
    lactose FLOAT
);
CREATE TABLE IF NOT EXISTS Protein(
    owner_id BIGINT ,
    FOREIGN KEY (owner_id) REFERENCES Meals(meal_id) ON DELETE CASCADE, 
    PRIMARY KEY (owner_id),
    alanine FLOAT,
    arginine FLOAT,
    aspartic FLOAT,
    asparagine FLOAT,
    cysteine FLOAT,
    glutamate FLOAT,
    glutamine FLOAT,
    glycine FLOAT,
    histidine FLOAT,
    isoleucine FLOAT,
    leucine FLOAT,
    lysine FLOAT,
    methionine FLOAT,
    phenylalanine FLOAT,
    proline FLOAT,
    serine FLOAT,
    threonine FLOAT,
    tryptophan FLOAT,
    tyrosine FLOAT,
    valine FLOAT
);
CREATE TABLE IF NOT EXISTS Fat(
    owner_id BIGINT ,
    FOREIGN KEY (owner_id) REFERENCES Meals(meal_id) ON DELETE CASCADE, 
    PRIMARY KEY (owner_id),
    acetic FLOAT,
    butyric FLOAT,
    valeric FLOAT,
    caproic FLOAT,
    caprylic FLOAT,
    capric FLOAT,
    lauric FLOAT,
    myristic FLOAT,
    pentadecanoic FLOAT,
    palmitic FLOAT,
    margaric FLOAT,
    stearic FLOAT,
    arachidic FLOAT,
    behenic FLOAT,
    lignoceric FLOAT,
    cerotic FLOAT,
    myristoleic FLOAT,
    oleic FLOAT,
    eicosenoic FLOAT,
    erucic FLOAT,
    Nervonic FLOAT,
    linaleic FLOAT,
    alinolenic FLOAT,
    stearidonic FLOAT,
    glinolenic FLOAT,
    arachidonic FLOAT,
    eicosatetraenoic FLOAT,
    timnodinic FLOAT,
    clupanodonic FLOAT,
    cervonic FLOAT    
);

CREATE TABLE IF NOT EXISTS Fiber(
    owner_id BIGINT ,
    FOREIGN KEY (owner_id) REFERENCES Meals(meal_id) ON DELETE CASCADE, 
    PRIMARY KEY (owner_id),
    celulose FLOAT,
    methyl FLOAT,
    hemicellulose FLOAT,
    lignans FLOAT,
    plant_waxes FLOAT,
    resistant_starches FLOAT,
    bglucans FLOAT,
    pectins FLOAT,
    natural_gums FLOAT,
    inulin FLOAT,
    oligasaccherides FLOAT
);

--@block
CREATE TABLE IF NOT EXISTS Trainingsplan(
    Trainingsplan_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    owner_id BIGINT,
    FOREIGN KEY (owner_id) REFERENCES Users(user_id)
);
--@block
CREATE TABLE IF NOT EXISTS Workouts(
    Workout_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    owner_id BIGINT,
    FOREIGN KEY (owner_id) REFERENCES Trainingsplan(Trainingsplan_id)
);
--@block
CREATE TABLE IF NOT EXISTS Exercise(
    Exercise_id BIGINT AUTO_INCREMENT PRIMARY Key,
    owner_id BIGINT,
    FOREIGN KEY (owner_id) REFERENCES Workouts(Workout_id)
);
--@block
DROP TABLE Carbohydrates;
DROP TABLE Fat;
Drop TABLE Protein;
Drop TABLE MicroNutrients;
DROP TABLE Meals;
