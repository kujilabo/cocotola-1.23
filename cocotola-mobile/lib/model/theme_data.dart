import 'package:flutter/material.dart';

class ThemeProblemCard {

  ThemeProblemCard({
    required this.translationBackgroundColor,
    required this.translationPlainTextStyle,
    required this.translationProblemTextStyle,
    required this.englishBackgroundColor,
    required this.englishProblemDecoration,
    required this.englishAnswerDecoration,
    required this.englishPlainTextStyle,
    required this.englishProblemTextStyle,
    required this.englishAnswerTextStyle,
    required this.cusrotTextStyle,
  });
  final Color translationBackgroundColor;
  final TextStyle translationPlainTextStyle;
  final TextStyle translationProblemTextStyle;

  final Color englishBackgroundColor;
  final BoxDecoration englishProblemDecoration;
  final BoxDecoration englishAnswerDecoration;
  final TextStyle englishPlainTextStyle;
  final TextStyle englishProblemTextStyle;
  final TextStyle englishAnswerTextStyle;
  final TextStyle cusrotTextStyle;
}

class CustomThemeData {

  CustomThemeData({required this.problemCard});
  final ThemeProblemCard problemCard;
}
