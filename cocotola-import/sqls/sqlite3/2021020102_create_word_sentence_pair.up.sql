create table `word_sentence_pair` (
 `id` integer primary key
,`document_id` varchar(36)
,`workbook_id` int not null
,`src_sentence_number` int not null
,`src_lang2` varchar(2) not null
,`src_text` varchar(1000) not null
,`src_author` varchar(20) not null
,`dst_sentence_number` int not null
,`dst_lang2` varchar(2) not null
,`dst_text` varchar(1000) not null
,`dst_author` varchar(20) not null
,`created_at` datetime not null
,`updated_at` datetime not null
);

CREATE INDEX document_id_word_sentence_pair_idx ON word_sentence_pair(document_id);
CREATE INDEX workbook_id_word_sentence_pair_idx ON word_sentence_pair(workbook_id);
CREATE INDEX src_lang2_word_sentence_pair_idx ON word_sentence_pair(src_lang2);
CREATE INDEX dst_lang2_word_sentence_pair_idx ON word_sentence_pair(dst_lang2);
