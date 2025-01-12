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
  late TextEditingController controller1;
  TextEditingController c = TextEditingController();
  void _inputText(String text) {
    print('input text');
  }

  @override
  Widget build(BuildContext context) {
    final keyboardKey = GlobalObjectKey<KeyboardState>(context);
    controller0 = TextEditingController();
    var focusNode0 = FocusNode();
    focusNode0.addListener(() {
      if (focusNode0.hasFocus) {
        print('focusNode0 has focus');
        final state = keyboardKey.currentState;
        state!.setController(controller0);
      } else {
        print('focusNode0 doesnt have focus');
      }
    });

    controller1 = TextEditingController();
    var focusNode1 = FocusNode();
    focusNode1.addListener(() {
      if (focusNode1.hasFocus) {
        print('focusNode1 has focus');
        final state = keyboardKey.currentState;
        state!.setController(controller1);
      } else {
        print('focusNode1 doesnt have focus');
      }
    });
    var card = WordStudyProblem(
      englishTexts: [
        EnglishText('I'),
        EnglishText('always'),
        EnglishText('prefer'),
        EnglishText('meeting'),
        EnglishText('in'),
        EnglishText('person'),
        EnglishText('over',
            isProblem: true, controller: controller0, focusNode: focusNode0),
        EnglishText('talking'),
        EnglishText('on'),
        EnglishText('the',
            isProblem: true, controller: controller1, focusNode: focusNode1),
        EnglishText('phone.'),
      ],
      japaneseTexts: ['JAPANESE TITLE 1'],
    );
    // focusNode0.requestFocus();

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
            SizedBox(height: 40),
            Keyboard(key: keyboardKey, controllers: [controller0, controller1]),

            TextField(
                // controller: c,
                ),
            // Keyboard(controllers: []),
            ElevatedButton(
                onPressed: () {
                  print('pressed');
                },
                child: Text('aaa')),
          ],
        ),
      ),
    );
  }
}
