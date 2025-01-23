import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/widget/word_study/english_text.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/provider/text_list_provider.dart';

class WordStudyAnswer extends ConsumerWidget {
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
              child: Text('Check'),
            ),
          ),
          Expanded(
            child: ElevatedButton(
              onPressed: () {
                print('onPressed');
              },
              child: Text('Next'),
            ),
          ),
        ],
      ),
    );
  }
}
