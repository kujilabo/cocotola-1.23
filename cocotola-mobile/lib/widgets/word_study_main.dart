import 'package:flutter/src/widgets/framework.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widgets/keyboard.dart';
import 'package:mobile/widgets/text_list_provider.dart';
// import 'package:mobile/widgets/editor_screen.dart';
import 'package:mobile/models/problem_word_study.dart';
import 'package:mobile/widgets/text_list_provider.dart';
import 'package:mobile/widgets/english_text.dart';
import 'package:mobile/widgets/word_study_problem.dart';

class WordStudyMain extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('WordStudyMain build');
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final textFieldValueList = ref.watch(textFieldValueListProvider);

    final problem = ProblemWordStudy(
      englishList: [
        ProblemWordStudyEnglish('I'),
        ProblemWordStudyEnglish('always'),
        ProblemWordStudyEnglish('prefer'),
        ProblemWordStudyEnglish('meeting'),
        ProblemWordStudyEnglish('in'),
        ProblemWordStudyEnglish('person'),
        ProblemWordStudyEnglish('over', isProblem: true),
        ProblemWordStudyEnglish('talking'),
        ProblemWordStudyEnglish('on'),
        ProblemWordStudyEnglish('the', isProblem: true),
        ProblemWordStudyEnglish('phone.'),
      ],
      translationList: [
        ProblemWordStudyTranslation('aaa'),
        ProblemWordStudyTranslation('bbb'),
        ProblemWordStudyTranslation('ccc'),
      ],
    );

    final focusNodeList = List.generate(10, (index) => FocusNode());
    final controllerList =
        List.generate(10, (index) => TextEditingController());
    final completedList =
        textFieldValueList.texts.map((e) => e.completed).toList();

    var wordSturyProblem = WordStudyProblem(
      problem: problem,
      focusNodeList: focusNodeList,
      controllerList: controllerList,
      completedList: completedList,
    );

    focusNodeList.asMap().forEach((index, focusNode) {
      focusNode.addListener(() {
        if (focusNode.hasFocus) {
          print('focusNode ${index} has focus');
          textFieldListNotifier.setIndex(index);
          textFieldListNotifier.setPosition(
              index, controllerList[index].selection.baseOffset);
        } else {
          print('focusNode ${index} doesnt have focus');
        }
      });
    });
    controllerList.asMap().forEach((index, controller) {
      controller.text = textFieldValueList.texts[index].text;
      controller.addListener(() {
        print('over == ${controllerList[index].text}');
        if (completedList[index]) {
          return;
        }
        if ("over" == controllerList[index].text) {
          print("SET COMPLETESSSSS");
          // textFieldListNotifier.setComplete(index);
        }
      });
    });

    final index = textFieldValueList.index;
    controllerList[index].selection = TextSelection.fromPosition(
        TextPosition(offset: textFieldValueList.texts[index].position));
    print(
        'textFieldValueList.texts[${index}].position: ${textFieldValueList.texts[index].position}');

    return Scaffold(
      appBar: AppBar(
        title: Text('WordStudyMain'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            wordSturyProblem,
            Keyboard(
              onPresskey: (String text) {
                textFieldListNotifier.addText(text);
                var index = textFieldValueList.index;
                if ("over" == textFieldValueList.texts[index].text) {
                  textFieldListNotifier.setComplete(index);
                }
              },
              onPressBackspace: () {
                textFieldListNotifier.backspace();
              },
            ),
          ],
        ),
      ),
    );
  }

  void _onPressBackspace() {}
}
