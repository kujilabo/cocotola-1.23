import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/util/logger.dart';
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
      case WordStudyStatus.end:
        throw UnimplementedError();
    }
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    logger.i('WordStudy build');
    final wordStudyStatus = ref.watch(wordStudyStatusProvider);

    const header = Text('aaaaaaaaaaaaaaa');
    final main = _buildMain(wordStudyStatus);
    return Scaffold(
      appBar: AppBar(
        title: const Text('WordStudy'),
      ),
      body: Column(
        children: [
          header,
          Expanded(
            child: main, // Expandedウィジェットでmainを包む
          ),
        ],
      ),
    );
  }
}
