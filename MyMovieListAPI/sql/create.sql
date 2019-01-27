CREATE TYPE loginTypes AS ENUM ('fb','google');

CREATE TABLE IF NOT EXISTS Users (
  id varchar (128),
  username varchar (20) NOT NULL,
  dateCreated timestamp NOT NULL DEFAULT NOW(),
  lastLogin timestamp NOT NULL DEFAULT NOW(),
  loginType loginTypes NOT NULL,
  public boolean NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS Lists (
  id serial,
  owner varchar (128) NOT NULL,
  public boolean NOT NULL DEFAULT FALSE,
  items json NOT NULL DEFAULT '[]',
  PRIMARY KEY (id),
  FOREIGN KEY (owner) REFERENCES Users (id)
);

