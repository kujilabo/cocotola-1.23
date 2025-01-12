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
  late TextEditingController currCtrl;
  late TextSelection _selection;
  // late TextEditingController controller0;
  // late TextEditingController controller1;
  // late FocusNode focusNode0;
  // late FocusNode focusNode1;
  late List<TextEditingController> ctrls;
  late List<FocusNode> focusNodes;

  @override
  void dispose() {
    // _titleController.dispose();
    // _amountController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
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

    // print('build main');
    // controller0 = TextEditingController();
    // focusNode0 = FocusNode();
    // focusNode0.addListener(() {
    //   if (focusNode0.hasFocus) {
    //     print('focusNode0 has focus');
    //     currCtrl.removeListener(_onSelectionChanged);
    //     currCtrl = controller0;
    //     currCtrl.addListener(_onSelectionChanged);
    //   } else {
    //     print('focusNode0 doesnt have focus');
    //   }
    // });

    // controller1 = TextEditingController();
    // focusNode1 = FocusNode();
    // focusNode1.addListener(() {
    //   if (focusNode1.hasFocus) {
    //     print('focusNode1 has focus');
    //     currCtrl.removeListener(_onSelectionChanged);
    //     currCtrl = controller1;
    //     currCtrl.addListener(_onSelectionChanged);
    //   } else {
    //     print('focusNode1 doesnt have focus');
    //   }
    // });
    // currCtrl = controller0;
    currCtrl = ctrls[0];

    // currCtrl.selection =
    //     TextSelection.fromPosition(const TextPosition(offset: 0));

    _selection = currCtrl.selection;

    var card = WordStudyProblem(
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

        // if (index == 0) {
        //   print('nextFocus');
        //   // focusNode0.nextFocus();
        //   focusNode1.requestFocus();
        // }
      },
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
            SizedBox(height: 40),
            Keyboard(
              // key: keyboardKey,
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
