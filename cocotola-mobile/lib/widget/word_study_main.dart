import 'package:flutter/src/widgets/framework.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widget/keyboard.dart';
import 'package:mobile/provider/text_list_provider.dart';
// import 'package:mobile/widgets/editor_screen.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/provider/text_list_provider.dart';
import 'package:mobile/widget/english_text.dart';
import 'package:mobile/widget/word_study_problem.dart';
import 'package:mobile/gateway/problem_repository.dart';

class WordStudyMain extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('WordStudyMain build');
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final textFieldValueList = ref.watch(textFieldValueListProvider);
    final problemNotifier = ref.read(problemProvider.notifier);
    final problem = ref.watch(problemProvider);

    final numProblems = problem.getNumProblems();
    final focusNodeList = List.generate(numProblems, (index) => FocusNode());
    final controllerList =
        List.generate(numProblems, (index) => TextEditingController());
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

    final index = textFieldValueList.index;
    if (numProblems > 0) {
      print('index: $index');
      print(
          ' textFieldValueList.texts[index].position ${textFieldValueList.texts[index].position}');
      controllerList[index].selection = TextSelection.fromPosition(
          TextPosition(offset: textFieldValueList.texts[index].position));
      print(
          'textFieldValueList.texts.length: ${textFieldValueList.texts.length}');
      print(
          'textFieldValueList.texts[${index}].position: ${textFieldValueList.texts[index].position}');
    }

    return Scaffold(
      appBar: AppBar(
        title: Text('WordStudyMain'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(textFieldValueList.allCompleted.toString()),
            ElevatedButton(
              onPressed: () {
                problemNotifier.fetchProblem(0);
              },
              child: Text('---0---'),
            ),
            ElevatedButton(
              onPressed: () {
                problemNotifier.fetchProblem(1);
              },
              child: Text('---1---'),
            ),
            wordSturyProblem,
            Keyboard(
              onPresskey: (String text) {
                textFieldListNotifier.addText(text);
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
}
