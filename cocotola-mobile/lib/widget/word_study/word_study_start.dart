import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class WordStudyStart extends ConsumerWidget {
  const WordStudyStart({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
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
            // ElevatedButton(
            //   onPressed: () {
            //     Navigator.of(context)
            //         .push(MaterialPageRoute(builder: (context) {
            //       return WordStudyMain();
            //     }));
            //   },
            //   child: const Text('Save Expense'),
            // ),
          ],
        ),
      ),
    );
  }
}
