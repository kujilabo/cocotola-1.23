import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/provider/problem_provider.dart';
import 'package:mobile/widget/keyboard.dart';
import 'package:mobile/widget/word_study/problem_card.dart';

class WordStudyQuestion extends ConsumerWidget {
  const WordStudyQuestion({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('WordStudyQuestion build');
    final double screenWidth = MediaQuery.of(context).size.width;
    print('Screen width: $screenWidth');
    final double screenHeight = MediaQuery.of(context).size.height;
    print('Screen height: $screenHeight');

    final wordStudyStatusNotifier = ref.read(wordStudyStatusProvider.notifier);

    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final textFieldValueList = ref.watch(textFieldValueListProvider);

    final problemWithStatus = ref.watch(problemProvider);

    final problem = problemWithStatus.currentProblem;

    ref.listen(textFieldValueListProvider, (prev, next) {
      if (next.allCompleted) {
        wordStudyStatusNotifier.setAnswerStatus();
      }
    });

    // final numProblems = problem.getNumProblems();
    // final focusNodeList = List.generate(numProblems, (index) => FocusNode());
    // final controllerList =
    //     List.generate(numProblems, (index) => TextEditingController());
    final completedList =
        textFieldValueList.texts.map((e) => e.completed).toList();
    final texts = textFieldValueList.texts.map((e) => e.text).toList();

    var problemCard = ProblemCard(
      problem: problem,
      texts: texts,
      completedList: completedList,
    );

    // focusNodeList.asMap().forEach((index, focusNode) {
    //   focusNode.addListener(() {
    //     if (focusNode.hasFocus) {
    //       print('focusNode $index has focus');
    //       textFieldListNotifier.setIndex(index);
    //       textFieldListNotifier.setPosition(
    //           index, controllerList[index].selection.baseOffset);
    //     } else {
    //       print('focusNode $index doesnt have focus');
    //     }
    //   });
    // });

    // controllerList.asMap().forEach((index, controller) {
    //   controller.text = textFieldValueList.texts[index].text;
    // });

    // final index = textFieldValueList.index;
    // if (numProblems > 0) {
    //   print('index: $index');
    //   print(
    //       ' textFieldValueList.texts[index].position ${textFieldValueList.texts[index].position}');
    //   controllerList[index].selection = TextSelection.fromPosition(
    //       TextPosition(offset: textFieldValueList.texts[index].position));
    //   print(
    //       'textFieldValueList.texts.length: ${textFieldValueList.texts.length}');
    //   print(
    //       'textFieldValueList.texts[$index].position: ${textFieldValueList.texts[index].position}');
    // }

    final bottomHeight = screenHeight * 0.3;
    final bottom = SizedBox(
      height: bottomHeight,
      child: Column(
        children: [
          Spacer(),
          Keyboard(
            onPresskey: (String text) => textFieldListNotifier.addText(text),
            onPressBackspace: () => textFieldListNotifier.backspace(),
          ),
        ],
      ),
    );

    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Expanded(
            child: Center(
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: problemCard,
              ),
            ),
          ),
          bottom,
        ],
      ),
    );
  }
}
