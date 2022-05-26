CREATE TABLE genres (
   id serial PRIMARY KEY,
   genre_name VARCHAR (200) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);

CREATE TABLE movies (
    id         serial PRIMARY KEY,       
	title       VARCHAR (200) UNIQUE NOT NULL,      
	description VARCHAR (200) UNIQUE NOT NULL,    
	year        int,            
	release_date TIMESTAMP NOT NULL,      
	runtime     int,            
	rating      int,            
	mpaa_rating  VARCHAR (200) UNIQUE NOT NULL,         
	created_at   TIMESTAMP NOT NULL,     
	updated_at   TIMESTAMP NOT NULL,     
	movie_genre  map[int]string 
);

CREATE TABLE users (
   id serial PRIMARY KEY,
   email VARCHAR (200) UNIQUE NOT NULL,
   password VARCHAR (200) UNIQUE NOT NULL,
);
//	flag.StringVar(&cfg.jwt.secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
