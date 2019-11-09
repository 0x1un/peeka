create table if not exists atten_list (id integer primary key,planid int not null,checktype text not null, approveid integer,userid text not null,classid integer not null,classsettingid integer not null,planchecktime text not null,groupid integer not null,createdat text not null);

create table if not exists dep_users (id integer primary key,name text not null,userid text not null,createdat text not null);
