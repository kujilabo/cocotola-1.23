import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/custom_theme_data_provider.dart';

class CustomAnswerTextField extends ConsumerWidget {
  final String text;

  const CustomAnswerTextField({
    super.key,
    required this.text,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final customThemeData = ref.watch(customThemDataProvider);
    final customTheme = customThemeData.problemCard;
    final style = customTheme.englishProblemTextStyle;
    final decoration = customTheme.englishAnswerDecoration;

    return Container(
      padding: EdgeInsets.only(left: 0, right: 10),
      child:
          DecoratedBox(decoration: decoration, child: Text(text, style: style)),
    );
  }
}
