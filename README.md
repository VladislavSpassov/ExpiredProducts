# ExpiredProducts

## Project Idea:

The idea of the project is the user to be able to upload groceries with name, quantity and expiration date.
Also the user should be able to upload recipies which have list of products and name.
The main idea is the user to receive suggestions for recipies based on his current products.



## Architechture 

Because the project is for educational purpose, I have choosen MVC architechture. 

The models for the DB are in the models folder, the controllers - in controllers folder and views - in views folder. 

I have 3 models. Product, Recipy and Recipy_Product. The relation between a recipy and recipy product is many2many, because 
one product can be in a more than 1 recipy.



## Technologies

gorm - An ORM for queries on the DB 
gin - http web framework
go doc - for documentation