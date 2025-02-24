class EnglishWord {
  const EnglishWord({required this.text, required this.isProblem});
  final String text;
  final bool isProblem;
}

class TranslationWord {
  const TranslationWord({required this.text, required this.isProblem});
  final String text;
  final bool isProblem;
}

class EnglishWordProblem {

  const EnglishWordProblem(
      {required this.translationWords, required this.englishWords,});
  final List<TranslationWord> translationWords;
  final List<EnglishWord> englishWords;
}

class EnglishWordProblemSet {

  const EnglishWordProblemSet({required this.englishWordProblems});
  final List<EnglishWordProblem> englishWordProblems;
}
