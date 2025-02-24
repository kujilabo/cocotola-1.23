import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/word_study/word_study_answer.dart';
import 'package:mobile/widget/word_study/word_study_question.dart';

class WordStudy extends ConsumerWidget {
  const WordStudy({super.key});

  Widget _buildMain(WordStudyStatus status) {
    switch (status) {
      case WordStudyStatus.question:
        return const WordStudyQuestion();
      case WordStudyStatus.answer:
        return const WordStudyAnswer();
      default:
        return const Text('aaaaa');
    }
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('WordStudy build');
    final wordStudyStatus = ref.watch(wordStudyStatusProvider);

    final header = const Text('aaaaaaaaaaaaaaa');
    final main = _buildMain(wordStudyStatus);
    return Scaffold(
      appBar: AppBar(
        title: const Text('WordStudy'),
      ),
      body: Container(
        child: Column(
          children: [
            header,
            Expanded(
              child: main, // Expandedウィジェットでmainを包む
            ),
          ],
        ),
      ),
    );
  }
}
