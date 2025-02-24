class WordProblemEnglish {

  WordProblemEnglish(this.text, {this.isProblem = false});
  final String text;
  final bool isProblem;
}

class WordProblemTranslation {
  WordProblemTranslation(this.text, {this.isProblem = false});
  final String text;
  final bool isProblem;
}

class WordProblem {

  const WordProblem({
    required this.englishList,
    required this.translationList,
  });
  final List<WordProblemEnglish> englishList;
  final List<WordProblemTranslation> translationList;

  int getNumProblems() {
    return englishList.where((element) => element.isProblem).length;
  }
}
