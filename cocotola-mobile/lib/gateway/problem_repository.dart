import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';

class ProblemRepository extends Notifier<WordProblem> {
  @override
  WordProblem build() {
    return WordProblem(englishList: [], translationList: []);
  }

  void fetchProblem(int index) {
    if (index == 0) {
      state = WordProblem(
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
          WordProblemEnglish('the', isProblem: true),
          WordProblemEnglish('phone.'),
        ],
        translationList: [
          WordProblemTranslation('aaa'),
          WordProblemTranslation('bbb'),
          WordProblemTranslation('ccc'),
        ],
      );
    } else {
      state = WordProblem(
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
      );
    }
  }
}

final problemProvider =
    NotifierProvider<ProblemRepository, WordProblem>(ProblemRepository.new);
