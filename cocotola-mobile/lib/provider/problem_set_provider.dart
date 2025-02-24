import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';

class ProblemSetRepository extends AsyncNotifier<List<WordProblem>> {
  @override
  Future<List<WordProblem>> build() async {
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
    'https://jsonplaceholder.typicode.com/post';
    state = const AsyncValue.data([]);
  }
}

final problemSetProvider =
    AsyncNotifierProvider<ProblemSetRepository, List<WordProblem>>(
        ProblemSetRepository.new,);
