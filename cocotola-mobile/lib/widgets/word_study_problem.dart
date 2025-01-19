import 'package:flutter/material.dart';
import 'package:mobile/widgets/english_text.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/models/problem_word_study.dart';

class WordStudyProblem extends ConsumerWidget {
  late void Function(int) onCompletedWord;
  final ProblemWordStudy problem;
  final List<FocusNode> focusNodeList;
  final List<TextEditingController> controllerList;

  WordStudyProblem({
    super.key,
    required this.problem,
    required this.focusNodeList,
    required this.controllerList,
  }) {
    onCompletedWord = (int index) => {};
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    List<Widget> englishTexts = [];
    var index = 0;
    var length = problem.englishList.length;
    print('length: $length');

    var firstProblem = true;

    for (var i = 0; i < length; i++) {
      print('i: $i');
      final english = problem.englishList[i];
      if (english.isProblem) {
        englishTexts.add(EnglishBlankTextWidget(
          index: index,
          englishText: english.text,
          controller: controllerList[index],
          focusNode: focusNodeList[index],
          onCompleted: onCompletedWord,
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
