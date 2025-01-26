import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';

class ProblemSetRepository extends Notifier<List<WordProblem>> {
  @override
  List<WordProblem> build() {
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
      )
    ];
  }

  void fetch() {
    state = [];
  }
}

final problemSetProvider =
    NotifierProvider<ProblemSetRepository, List<WordProblem>>(
        ProblemSetRepository.new);
