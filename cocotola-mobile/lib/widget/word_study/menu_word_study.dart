import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/word_study/word_study.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';

class MenuWordStudy extends ConsumerWidget {
  const MenuWordStudy({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final wordStudyStatusNotifier = ref.watch(wordStudyStatusProvider.notifier);

    // final textFieldValueListProvider = ref.watch(textFieldValueListProvider);
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
                ref.invalidate(textFieldValueListProvider);

                Navigator.of(context).push(MaterialPageRoute(
                  builder: (context) => WordStudy(),
                ));
              },
              child: const Text('Save Expense'),
            ),
          ],
        ),
      ),
    );
  }
}
