import 'package:flutter/material.dart';
import 'package:mobile/widgets/word_study_main.dart';
import 'package:mobile/widgets/word_study_top.dart';
import 'package:mobile/widgets/keyboard.dart';

class WordStudyTop extends StatefulWidget {
  const WordStudyTop({super.key});

  @override
  State<WordStudyTop> createState() => _WordStudyTopState();
}

class _WordStudyTopState extends State<WordStudyTop> {
  @override
  Widget build(BuildContext context) {
    TextEditingController textController = TextEditingController();
    return Scaffold(
      appBar: AppBar(
        title: const Text('Word Study'),
      ),
      body: SafeArea(
        child: Column(
          children: [
            const Center(
              child: Text(
                'Word StudyTOP',
                style: TextStyle(fontSize: 24),
              ),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.of(context)
                    .push(MaterialPageRoute(builder: (context) {
                  return WordStudyMain();
                }));
              },
              child: const Text('Save Expense'),
            ),
            TextField(
              controller: textController,
              keyboardType: TextInputType.none,
            ),
            // const Spacer(),
            Keyboard(
              controller: textController,
            ),
          ],
        ),
      ),
    );
  }
}
