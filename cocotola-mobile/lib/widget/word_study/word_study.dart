import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widget/keyboard.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/word_study/word_study_answer.dart';
import 'package:mobile/widget/word_study/word_study_question.dart';
import 'package:mobile/widget/word_study/problem_card.dart';

class WordStudy extends ConsumerWidget {
  const WordStudy({super.key});

  Widget _buildMain(WordStudyStatus status) {
    switch (status) {
      case WordStudyStatus.question:
        return WordStudyQuestion();
      case WordStudyStatus.answer:
        return WordStudyAnswer();
      default:
        return Text('aaaaa');
    }
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('WordStudy build');
    final wordStudyStatus = ref.watch(wordStudyStatusProvider);

    final main = _buildMain(wordStudyStatus);
    return Scaffold(
      appBar: AppBar(
        title: Text('WordStudy'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('WordStudy'),
            main,
          ],
        ),
      ),
    );
  }
}
