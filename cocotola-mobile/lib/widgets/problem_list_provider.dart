import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/widgets/word_study_problem.dart';
import 'package:mobile/models/problem_word_study.dart';

class ProblemRepository extends Notifier<ProblemWordStudy> {
  @override
  ProblemWordStudy build() {
    return ProblemWordStudy(englishList: [], translationList: []);
  }

  void fetchProblem(int index) {
    state = ProblemWordStudy(
      englishList: [
        ProblemWordStudyEnglish('I'),
        ProblemWordStudyEnglish('always'),
        ProblemWordStudyEnglish('prefer'),
        ProblemWordStudyEnglish('meeting'),
        ProblemWordStudyEnglish('in'),
        ProblemWordStudyEnglish('person'),
        ProblemWordStudyEnglish('over', isProblem: true),
        ProblemWordStudyEnglish('talking'),
        ProblemWordStudyEnglish('on'),
        ProblemWordStudyEnglish('the', isProblem: true),
        ProblemWordStudyEnglish('phone.'),
      ],
      translationList: [
        ProblemWordStudyTranslation('aaa'),
        ProblemWordStudyTranslation('bbb'),
        ProblemWordStudyTranslation('ccc'),
      ],
    );
  }
}

final problemProvider = NotifierProvider<ProblemRepository, ProblemWordStudy>(
    ProblemRepository.new);
