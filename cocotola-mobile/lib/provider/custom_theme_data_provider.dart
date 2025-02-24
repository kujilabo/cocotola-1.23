import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/theme_data.dart';

class ThemeDataRepository extends Notifier<CustomThemeData> {
  @override
  CustomThemeData build() {
    return CustomThemeData(
      problemCard: ThemeProblemCard(
        translationBackgroundColor: Colors.lightBlue,
        translationPlainTextStyle: const TextStyle(
          fontSize: 20,
          color: Colors.black,
        ),
        translationProblemTextStyle: const TextStyle(
          fontSize: 20,
          color: Colors.red,
        ),
        englishBackgroundColor: Colors.green,
        englishProblemDecoration: BoxDecoration(
          color: Colors.grey[200],
          borderRadius: const BorderRadius.vertical(
            top: Radius.circular(4.0),
            bottom: Radius.circular(4.0),
          ),
        ),
        englishAnswerDecoration: const BoxDecoration(
          color: Colors.blue,
          borderRadius: BorderRadius.vertical(
            top: Radius.circular(4.0),
            bottom: Radius.circular(4.0),
          ),
        ),
        englishPlainTextStyle: const TextStyle(
          fontSize: 20,
          color: Colors.black,
        ),
        englishProblemTextStyle: const TextStyle(
          fontSize: 20,
          color: Colors.grey,
        ),
        englishAnswerTextStyle: const TextStyle(
          fontSize: 20,
          color: Colors.white,
        ),
        cusrotTextStyle: const TextStyle(color: Colors.green, fontSize: 20),
      ),
    );
  }
}

final customThemDataProvider =
    NotifierProvider<ThemeDataRepository, CustomThemeData>(
        ThemeDataRepository.new);
