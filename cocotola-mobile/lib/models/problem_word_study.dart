class ProblemWordStudyEnglish {
  final String text;
  final bool isProblem;

  ProblemWordStudyEnglish(this.text, {this.isProblem = false});
}

class ProblemWordStudyTranslation {
  final String text;
  final bool isProblem;
  ProblemWordStudyTranslation(this.text, {this.isProblem = false});
}

class ProblemWordStudy {
  final List<ProblemWordStudyEnglish> englishList;
  final List<ProblemWordStudyTranslation> translationList;

  const ProblemWordStudy({
    required this.englishList,
    required this.translationList,
  });
}
