import 'package:flutter_riverpod/flutter_riverpod.dart';

enum WordStudyStatus { question, answer, end }

class WordStudyStatusNotifier extends Notifier<WordStudyStatus> {
  @override
  WordStudyStatus build() {
    return WordStudyStatus.question;
  }

  void setQuestionStatus() {
    state = WordStudyStatus.question;
  }

  void setAnswerStatus() {
    state = WordStudyStatus.answer;
  }

  void setEndStatus() {
    state = WordStudyStatus.end;
  }
}

final wordStudyStatusProvider =
    NotifierProvider<WordStudyStatusNotifier, WordStudyStatus>(
        WordStudyStatusNotifier.new,);
