import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widget/keyboard.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/word_study/word_study_answer.dart';
import 'package:mobile/widget/word_study/problem_card.dart';
import 'package:mobile/provider/problem_provider.dart';

class WordStudyAnswer extends ConsumerWidget {
  const WordStudyAnswer({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final problemNotifier = ref.read(problemProvider.notifier);
    final problemWithStatus = ref.watch(problemProvider);
    final wordStudyStatusNotifier = ref.read(wordStudyStatusProvider.notifier);
    return Container(
      child: Row(
        children: [
          Expanded(
            child: ElevatedButton(
              onPressed: () {
                problemNotifier.next();
                wordStudyStatusNotifier.setQuestionStatus();
              },
              child: Text('Check'),
            ),
          ),
          Expanded(
            child: ElevatedButton(
              onPressed: () {
                problemNotifier.next();
                wordStudyStatusNotifier.setQuestionStatus();
              },
              child: Text('Next'),
            ),
          ),
        ],
      ),
    );
  }
}
