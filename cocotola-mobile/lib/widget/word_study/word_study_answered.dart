import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/util/logger.dart';

class WordStudyButtons extends ConsumerWidget {
  const WordStudyButtons({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Row(
      children: [
        Expanded(
          child: ElevatedButton(
            onPressed: () {
              logger.i('onPressed');
            },
            child: const Text('Check'),
          ),
        ),
        Expanded(
          child: ElevatedButton(
            onPressed: () {
              logger.i('onPressed');
            },
            child: const Text('Next'),
          ),
        ),
      ],
    );
  }
}
