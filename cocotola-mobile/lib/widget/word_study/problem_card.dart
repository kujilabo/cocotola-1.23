import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/provider/custom_theme_data_provider.dart';
import 'package:mobile/widget/word_study/plain_text.dart';
import 'package:mobile/widget/word_study/custom_problem_text_field.dart';

class ProblemCard extends ConsumerWidget {
  final WordProblem problem;
  final List<String> texts;
  final List<bool> completedList;

  const ProblemCard({
    required this.problem, required this.texts, required this.completedList, super.key,
  });

  List<Widget> _buildEnglishTexts(
    TextStyle plainTextStyle,
    TextStyle answerTextStyle,
  ) {
    List<Widget> widgets = [];
    var index = 0;
    for (final english in problem.englishList) {
      if (english.isProblem) {
        if (completedList[index]) {
          widgets.add(PlainText(text: english.text, style: answerTextStyle));
        } else {
          widgets.add(CustomProblemTextField(
            index: index,
            text: texts[index],
          ));
        }
        index++;
      } else {
        widgets.add(PlainText(text: english.text, style: plainTextStyle));
      }
    }

    return widgets;
  }

  List<Widget> _buildTranslationTexts(
    TextStyle plainTextStyle,
    TextStyle problemTextStyle,
  ) {
    List<Widget> widgets = [];
    for (final translation in problem.translationList) {
      final style = translation.isProblem ? problemTextStyle : plainTextStyle;
      widgets.add(PlainText(text: translation.text, style: style));
    }
    return widgets;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final customThemeData = ref.watch(customThemDataProvider);
    final theme = customThemeData.problemCard;
    List<Widget> englishTexts = _buildEnglishTexts(
      customThemeData.problemCard.englishPlainTextStyle,
      customThemeData.problemCard.englishAnswerTextStyle,
    );
    List<Widget> translationTexts = _buildTranslationTexts(
      customThemeData.problemCard.translationPlainTextStyle,
      customThemeData.problemCard.translationProblemTextStyle,
    );

    return Card(
      color: Colors.transparent,
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          DecoratedBox(
            decoration: BoxDecoration(
              color: theme.englishBackgroundColor,
              borderRadius: BorderRadius.vertical(top: Radius.circular(16.0)),
            ),
            // color: Colors.green,
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: SizedBox(
                height: 150,
                width: double.infinity,
                child: Wrap(children: translationTexts),
              ),
            ),
          ),
          DecoratedBox(
            decoration: BoxDecoration(
              color: theme.translationBackgroundColor,
              borderRadius:
                  BorderRadius.vertical(bottom: Radius.circular(16.0)),
            ),
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: SizedBox(
                width: double.infinity,
                height: 150,
                child: Wrap(children: englishTexts),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
