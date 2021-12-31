create database world;
use world;

create table City(
	Id int primary key auto_increment,
    Name varchar(50),
    CountryCode varchar(50),
    District varchar(50),
    Population integer
    
)Engine=InnoDB;

insert into city(Name,CountryCode,District,Population)
values
("São Paulo", "Brasil", "Sudoeste",27458258),
("Rio de Janeito", "Brasil", "Sudoeste",14587859),
("Goiás", "Brasil", "Centro-Oeste",8458258);

select * from city;