import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/widgets/english_text.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/models/problem_word_study.dart';
import 'package:mobile/widgets/text_list_provider.dart';

class WordStudyProblem extends ConsumerWidget {
  late void Function(int) onCompletedWord;
  final ProblemWordStudy problem;
  final List<FocusNode> focusNodeList;
  final List<TextEditingController> controllerList;
  final List<bool> completedList;

  WordStudyProblem({
    super.key,
    required this.problem,
    required this.focusNodeList,
    required this.controllerList,
    required this.completedList,
  }) {
    onCompletedWord = (int index) => {};
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    List<Widget> englishTexts = [];

    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    {
      var index = 0;
      for (var i = 0; i < problem.englishList.length; i++) {
        if (problem.englishList[i].isProblem) {
          // print("SET COMPLETE");
          // textFieldListNotifier.setComplete(index);
          index++;
        }
      }
    }
    // controllerList.asMap().forEach((index, controller) {
    //   // controller.addListener(() {
    //   if (controller.text == problem.englishList[index].text) {
    //     // onCompleted(index);
    //     // setState(() {
    //     //   readOnly = true;
    //     //   color = Colors.red;
    //   }
    //   print("${controller.text} : ${problem.englishList[index].text}");
    //   // });
    // });
    var index = 0;
    var length = problem.englishList.length;
    print('length: $length');

    var firstProblem = true;

    for (var i = 0; i < length; i++) {
      print('i: $i');
      final english = problem.englishList[i];
      if (english.isProblem) {
        // final completed = english.text == controllerList[index].text;
        print('${english.text} == ${controllerList[index].text}');
        englishTexts.add(EnglishBlankTextWidget(
          index: index,
          englishText: english.text,
          controller: controllerList[index],
          focusNode: focusNodeList[index],
          completed: completedList[index],
          onCompleted: () {
            // print('${english.text} == ${controllerList[index].text}');
            // if (english.text == controllerList[index].text) {
            //   print("SET COMPLETESSSSS");
            //   textFieldListNotifier.setComplete(index);
            // }
          },
          first: firstProblem,
        ));
        firstProblem = false;
        index++;
      } else {
        englishTexts.add(EnglishPlainTextWidget(
          englishText: english.text,
        ));
      }
    }
    return Card(
      child: Container(
        alignment: Alignment.topLeft,
        // height: 100.0,
        width: double.infinity,
        // color: Colors.red,
        padding: EdgeInsets.all(15),
        child: Column(
          children: [
            Wrap(children: englishTexts),
            SizedBox(height: 10),
            // Wrap(children: englishTexts),
          ],
        ),
      ),
    );
  }
}
