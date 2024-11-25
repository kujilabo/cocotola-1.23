create table `audio` (
 `id` int auto_increment
,`lang5` varchar(5) character set ascii not null
,`text` varchar(100) character set ascii not null
,`audio_content` text character set ascii not null
,`audio_length_sec` float not null
,primary key(`id`)
,unique(`lang5`, `text`)
);
