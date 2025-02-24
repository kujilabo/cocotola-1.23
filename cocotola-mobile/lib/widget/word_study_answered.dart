import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/widget/word_study/problem_text_field.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';

class WordStudyButtons extends ConsumerWidget {
  const WordStudyButtons({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Container(
      child: Row(
        children: [
          Expanded(
            child: ElevatedButton(
              onPressed: () {
                print('onPressed');
              },
              child: const Text('Check'),
            ),
          ),
          Expanded(
            child: ElevatedButton(
              onPressed: () {
                print('onPressed');
              },
              child: const Text('Next'),
            ),
          ),
        ],
      ),
    );
  }
}
