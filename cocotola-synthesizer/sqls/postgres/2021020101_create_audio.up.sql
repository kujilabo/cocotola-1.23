create table audio (
 id serial not null
,lang5 varchar(5) not null
,text varchar(100) not null
,audio_content text not null
,audio_length_sec float not null
,primary key(id)
,unique(lang5, text)
);
