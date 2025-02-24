import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';

import 'package:mobile/sql/sql.dart';
import 'package:mobile/util/logger.dart';

class ProblemSetRepository extends AsyncNotifier<List<WordProblem>> {
  DB? db;
  @override
  Future<List<WordProblem>> build() async {
    db = DB();
    await fetch();
    return [
      WordProblem(
        englishList: [
          WordProblemEnglish('I'),
          WordProblemEnglish('always'),
          WordProblemEnglish('prefer'),
          WordProblemEnglish('meeting'),
          WordProblemEnglish('in'),
          WordProblemEnglish('person'),
          WordProblemEnglish('over', isProblem: true),
          WordProblemEnglish('talking'),
          WordProblemEnglish('on'),
          WordProblemEnglish('on'),
          WordProblemEnglish('the', isProblem: true),
          WordProblemEnglish('phone.'),
        ],
        translationList: [
          WordProblemTranslation('aaa', isProblem: true),
          WordProblemTranslation('bbb'),
          WordProblemTranslation('ccc', isProblem: true),
        ],
      ),
      WordProblem(
        englishList: [
          WordProblemEnglish('I'),
          WordProblemEnglish('died', isProblem: true),
          WordProblemEnglish('of', isProblem: true),
          WordProblemEnglish('cancer.'),
        ],
        translationList: [
          WordProblemTranslation('aaa'),
          WordProblemTranslation('bbb'),
          WordProblemTranslation('ccc'),
        ],
      ),
    ];
  }

  Future<void> fetch() async {
    final database = await db?.db;
    final results = await database?.query('word_sentence_pair');
    logger.i('results:  $results');
    await database?.insert('word_sentence_pair', {
      'document_id': 'abc',
      'workbook_id': 1,
      'src_sentence_number': 100,
      'src_lang2': 'en',
      'src_text': 'hello',
      'src_author': 'xxx',
      'dst_sentence_number': 200,
      'dst_lang2': 'en',
      'dst_text': 'hello',
      'dst_author': 'xxx',
      'created_at': '2021-01-01 00:00:00',
      'updated_at': '2021-01-01 00:00:00',
    });
    // 'https://jsonplaceholder.typicode.com/post';
    // final db = FirebaseFirestore.instance;
    // final snapshot = await db.collection('word_sentence_pair').get();
    // logger.i('snapshot: $snapshot');
    // logger.i('snapshot: $snapshot.docs');
    // for (final doc in snapshot.docs) {
    //   logger.i('doc: $doc');
    //   logger.i('doc: ${doc.data()}');
    // }
    // state = const AsyncValue.data([]);
  }
}

final problemSetProvider =
    AsyncNotifierProvider<ProblemSetRepository, List<WordProblem>>(
  ProblemSetRepository.new,
);
