class WordProblemEnglish {
  final String text;
  final bool isProblem;

  WordProblemEnglish(this.text, {this.isProblem = false});
}

class WordProblemTranslation {
  final String text;
  final bool isProblem;
  WordProblemTranslation(this.text, {this.isProblem = false});
}

class WordProblem {
  final List<WordProblemEnglish> englishList;
  final List<WordProblemTranslation> translationList;

  const WordProblem({
    required this.englishList,
    required this.translationList,
  });

  int getNumProblems() {
    return englishList.where((element) => element.isProblem).length;
  }
}
