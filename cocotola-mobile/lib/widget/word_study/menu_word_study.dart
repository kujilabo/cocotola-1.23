import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/word_study/word_study.dart';

class MenuWordStudy extends ConsumerWidget {
  const MenuWordStudy({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final wordStudyStatusNotifier = ref.watch(wordStudyStatusProvider.notifier);
    return Scaffold(
      appBar: AppBar(
        title: const Text('Word Study'),
      ),
      body: SafeArea(
        child: Column(
          children: [
            TextField(
              controller: TextEditingController(),
            ),
            const Center(
              child: Text(
                'Word StudyTOP',
                style: TextStyle(fontSize: 24),
              ),
            ),
            ElevatedButton(
              onPressed: () {
                wordStudyStatusNotifier.setQuestionStatus();
                Navigator.of(context)
                    .push(MaterialPageRoute(builder: (context) {
                  return WordStudy();
                }));
              },
              child: const Text('Save Expense'),
            ),
          ],
        ),
      ),
    );
  }
}
