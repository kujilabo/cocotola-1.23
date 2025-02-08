create table `tatoeba_link` (
 `src` int not null
,`dst` int not null
,unique(`src`, `dst`)
,foreign key(`src`) references `tatoeba_sentence`(`sentence_number`) on delete cascade
,foreign key(`dst`) references `tatoeba_sentence`(`sentence_number`) on delete cascade
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
