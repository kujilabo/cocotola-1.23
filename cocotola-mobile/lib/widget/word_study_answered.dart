import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

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
