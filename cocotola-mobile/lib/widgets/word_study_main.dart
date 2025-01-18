import 'package:flutter/material.dart';
import 'package:mobile/widgets/keyboard.dart';
import 'package:mobile/widgets/english_text.dart';
import 'package:mobile/widgets/word_study_problem.dart';
import 'package:mobile/models/english_word_problem.dart';

class WordStudyMain extends StatefulWidget {
  List<EnglishWordProblem> problems;
  WordStudyMain({Key? key}) : this.init(key: key);
  WordStudyMain.init({
    super.key,
    this.problems = const [
      EnglishWordProblem(
        translationWords: [
          TranslationWord(text: '私は', isProblem: false),
          TranslationWord(text: 'いつも', isProblem: false),
          TranslationWord(text: '電話で話すより', isProblem: true),
          TranslationWord(text: '会って話すことを好む。', isProblem: false),
        ],
        englishWords: [
          EnglishWord(text: 'I', isProblem: false),
          EnglishWord(text: 'always', isProblem: false),
          EnglishWord(text: 'prefer', isProblem: false),
          EnglishWord(text: 'meeting', isProblem: false),
          EnglishWord(text: 'in', isProblem: false),
          EnglishWord(text: 'person', isProblem: false),
          EnglishWord(text: 'over', isProblem: true),
          EnglishWord(text: 'talking', isProblem: false),
          EnglishWord(text: 'on', isProblem: false),
          EnglishWord(text: 'the', isProblem: true),
          EnglishWord(text: 'phone.', isProblem: false),
        ],
      ),
    ],
  });

  @override
  State<WordStudyMain> createState() => _WordStudyMainState();
}

class _WordStudyMainState extends State<WordStudyMain> {
  late TextEditingController currCtrl;
  late TextSelection _selection;
  late List<TextEditingController> ctrls;
  late List<FocusNode> focusNodes;

  @override
  void dispose() {
    currCtrl.addListener(_onSelectionChanged);
    for (var i = 0; i < 10; i++) {
      ctrls[i].dispose();
      focusNodes[i].dispose();
    }
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    print('build main');
    ctrls = [];
    focusNodes = [];
    for (var i = 0; i < 10; i++) {
      ctrls.add(TextEditingController());
      focusNodes.add(FocusNode());
      focusNodes[i].addListener(() {
        if (focusNodes[i].hasFocus) {
          print('focusNode0 has focus');
          currCtrl.removeListener(_onSelectionChanged);
          currCtrl = ctrls[i];
          currCtrl.addListener(_onSelectionChanged);
        } else {
          print('focusNode0 doesnt have focus');
        }
      });
    }

    currCtrl = ctrls[0];

    _selection = currCtrl.selection;

    var card = buildProblem(widget.problems[0]);
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
            Keyboard(
              onPresskey: _onPressKey,
              onPressBackspace: _onPressBackspace,
            ),
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

  Widget buildProblem(EnglishWordProblem problem) {
    return WordStudyProblem(
      englishTexts: [
        EnglishText('I'),
        EnglishText('always'),
        EnglishText('prefer'),
        EnglishText('meeting'),
        EnglishText('in'),
        EnglishText('person'),
        EnglishText('over',
            isProblem: true,
            controller: ctrls[0],
            focusNode: focusNodes[0],
            first: true),
        EnglishText('talking'),
        EnglishText('on'),
        EnglishText('the',
            isProblem: true, controller: ctrls[1], focusNode: focusNodes[1]),
        EnglishText('phone.'),
      ],
      japaneseTexts: ['JAPANESE TITLE 1'],
      onCompletedWord: (int index) {
        print('completed word $index');
        var nextIndex = (index + 1) % 10;
        focusNodes[nextIndex].requestFocus();
      },
    );
  }

  void _onPressKey(String text) {
    final value = currCtrl.text;
    if (currCtrl.text.isEmpty) {
      currCtrl.text = value + text;
      currCtrl.selection =
          TextSelection.fromPosition(const TextPosition(offset: 1));
      return;
    }

    final position = _selection.base.offset;
    print('position: $position');
    final suffix = value.substring(position, value.length);
    currCtrl.text = value.substring(0, position) + text + suffix;
    currCtrl.selection =
        TextSelection.fromPosition(TextPosition(offset: position + 1));
  }

  void _onPressBackspace() {
    final value = currCtrl.text;
    final position = _selection.base.offset;

    if (value.isEmpty || position == 0) {
      return;
    }

    var suffix = value.substring(position, value.length);
    currCtrl.text = value.substring(0, position - 1) + suffix;
    currCtrl.selection =
        TextSelection.fromPosition(TextPosition(offset: position - 1));
  }

  void _onSelectionChanged() {
    _selection = currCtrl.selection;
    print('Cursor position: ${_selection.base.offset}');
  }
}
