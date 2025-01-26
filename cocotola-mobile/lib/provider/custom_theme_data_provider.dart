import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/problem_set_provider.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/model/theme_data.dart';

class ThemeDataRepository extends Notifier<CustomThemeData> {
  @override
  CustomThemeData build() {
    return CustomThemeData(
      problemCard: ThemeProblemCard(
        translationBackgroundColor: Colors.lightBlue,
        translationPlainTextStyle: TextStyle(
          fontSize: 20,
          color: Colors.black,
        ),
        translationProblemTextStyle: TextStyle(
          fontSize: 20,
          color: Colors.red,
        ),
        englishBackgroundColor: Colors.green,
        englishProblemDecoration: BoxDecoration(
          color: Colors.grey[200],
          borderRadius: BorderRadius.vertical(
            top: Radius.circular(4.0),
            bottom: Radius.circular(4.0),
          ),
        ),
        englishAnswerDecoration: BoxDecoration(
          color: Colors.blue,
          borderRadius: BorderRadius.vertical(
            top: Radius.circular(4.0),
            bottom: Radius.circular(4.0),
          ),
        ),
        englishPlainTextStyle: TextStyle(
          fontSize: 20,
          color: Colors.black,
        ),
        englishProblemTextStyle: TextStyle(
          fontSize: 20,
          color: Colors.grey,
        ),
        englishAnswerTextStyle: TextStyle(
          fontSize: 20,
          color: Colors.white,
        ),
        cusrotTextStyle: TextStyle(color: Colors.green, fontSize: 20),
      ),
    );
  }
}

final customThemDataProvider =
    NotifierProvider<ThemeDataRepository, CustomThemeData>(
        ThemeDataRepository.new);
