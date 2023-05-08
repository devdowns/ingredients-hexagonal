-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table for recipe ingredients
CREATE TABLE ingredients (
                             id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                             name VARCHAR(255) NOT NULL,
                             description VARCHAR(255) NOT NULL
);

INSERT INTO ingredients (name, description) VALUES ('Flour', 'White all-purpose flour');
INSERT INTO ingredients (name, description) VALUES ('Sugar', 'Granulated white sugar');
INSERT INTO ingredients (name, description) VALUES ('Butter', 'Unsalted butter');
INSERT INTO ingredients (name, description) VALUES ('Eggs', 'Large eggs');
INSERT INTO ingredients (name, description) VALUES ('Milk', 'Whole milk');
INSERT INTO ingredients (name, description) VALUES ('Vanilla Extract', 'Pure vanilla extract');
INSERT INTO ingredients (name, description) VALUES ('Baking Powder', 'Leavening agent for baking');
INSERT INTO ingredients (name, description) VALUES ('Salt', 'Table salt');
INSERT INTO ingredients (name, description) VALUES ('Chocolate Chips', 'Semi-sweet chocolate chips');
