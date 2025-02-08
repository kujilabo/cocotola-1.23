create table `tatoeba_link` (
 `src` int not null
,`dst` int not null
,unique(`src`, `dst`)
,foreign key(`src`) references `tatoeba_sentence`(`sentence_number`)
,foreign key(`dst`) references `tatoeba_sentence`(`sentence_number`)
);
