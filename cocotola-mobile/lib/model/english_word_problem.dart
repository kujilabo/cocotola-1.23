class EnglishWord {
  final String text;
  final bool isProblem;
  const EnglishWord({required this.text, required this.isProblem});
}

class TranslationWord {
  final String text;
  final bool isProblem;
  const TranslationWord({required this.text, required this.isProblem});
}

class EnglishWordProblem {
  final List<TranslationWord> translationWords;
  final List<EnglishWord> englishWords;

  const EnglishWordProblem(
      {required this.translationWords, required this.englishWords});
}

class EnglishWordProblemSet {
  final List<EnglishWordProblem> englishWordProblems;

  const EnglishWordProblemSet({required this.englishWordProblems});
}
