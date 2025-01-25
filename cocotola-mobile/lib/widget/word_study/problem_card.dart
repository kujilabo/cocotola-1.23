import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/widget/word_study/plain_text.dart';
import 'package:mobile/widget/word_study/problem_text_field.dart';

class ProblemCard extends ConsumerWidget {
  final WordProblem problem;
  final List<FocusNode> focusNodeList;
  final List<TextEditingController> controllerList;
  final List<bool> completedList;

  const ProblemCard({
    super.key,
    required this.problem,
    required this.focusNodeList,
    required this.controllerList,
    required this.completedList,
  });

  List<Widget> _buildEnglishTexts() {
    List<Widget> englishTexts = [];
    var index = 0;
    var firstProblem = true;
    for (final english in problem.englishList) {
      if (english.isProblem) {
        // print('${english.text} == ${controllerList[index].text}');
        englishTexts.add(ProblemTextField(
          index: index,
          englishText: english.text,
          controller: controllerList[index],
          focusNode: focusNodeList[index],
          completed: completedList[index],
          first: firstProblem,
        ));
        firstProblem = false;
        index++;
      } else {
        englishTexts.add(PlainText(
          text: english.text,
        ));
      }
    }

    return englishTexts;
  }

  List<Widget> _buildTranslationTexts() {
    List<Widget> translationTexts = [];
    for (final translation in problem.translationList) {
      translationTexts.add(PlainText(
        text: translation.text,
      ));
    }
    return translationTexts;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    List<Widget> englishTexts = _buildEnglishTexts();
    List<Widget> translationTexts = _buildTranslationTexts();

    return Card(
      child: Container(
        alignment: Alignment.topLeft,
        // height: 100.0,
        width: double.infinity,
        color: Colors.blueGrey[50],
        // color: Colors.red,
        padding: EdgeInsets.all(15),
        child: Column(
          children: [
            Wrap(children: englishTexts),
            SizedBox(height: 10),
            Wrap(children: translationTexts),
          ],
        ),
      ),
    );
  }
}
