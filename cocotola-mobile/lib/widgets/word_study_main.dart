import 'package:flutter/material.dart';
import 'package:mobile/widgets/keyboard.dart';
import 'package:mobile/widgets/english_text.dart';
import 'package:mobile/widgets/word_study_problem.dart';

class WordStudyMain extends StatefulWidget {
  const WordStudyMain({super.key});

  @override
  State<WordStudyMain> createState() => _WordStudyMainState();
}

class _WordStudyMainState extends State<WordStudyMain> {
  late TextEditingController controller0;
  @override
  Widget build(BuildContext context) {
    var controller0 = TextEditingController();
    var card = WordStudyProblem(
      englishTexts: [
        EnglishText('I'),
        EnglishText('always'),
        EnglishText('prefer'),
        EnglishText('meeting'),
        EnglishText('in'),
        EnglishText('person'),
        // EnglishText('over', isProblem: true, controller: controller0),
        EnglishText('talking'),
        EnglishText('on'),
        EnglishText('the'),
        EnglishText('phone.'),
      ],
      japaneseTexts: ['JAPANESE TITLE 1'],
    );

    return Scaffold(
      appBar: AppBar(
        title: const Text('Word Study'),
      ),
      body: SafeArea(
        child: Column(
          children: [
            Container(
              // height: 100.0,
              width: double.infinity,
              // color: Colors.red,
              child: Padding(
                padding: EdgeInsets.all(15),
                child: card,
              ),
            ),
            // Keyboard(controllers: [controller0]),
            // Keyboard(controllers: []),
          ],
        ),
      ),
    );
  }
}
