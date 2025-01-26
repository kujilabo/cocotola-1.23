import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/widget/word_study/plain_text.dart';
import 'package:mobile/widget/word_study/problem_text_field.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'dart:async';
import 'package:mobile/provider/custom_theme_data_provider.dart';

class TimerState {
  final bool flag;
  final Timer? timer;

  const TimerState({required this.flag, required this.timer});
}

class TimerRepository extends Notifier<TimerState> {
  @override
  TimerState build() {
    ref.onDispose(() {
      state.timer?.cancel();
    });
    // final now = DateTime.now();
    final interval = const Duration(milliseconds: 800);
    return TimerState(flag: true, timer: Timer.periodic(interval, _onTick));
  }

  void _onTick(Timer timer) {
    state = TimerState(flag: !state.flag, timer: timer);
  }
}

final timerProvider =
    NotifierProvider<TimerRepository, TimerState>(TimerRepository.new);

class BlinkText extends ConsumerWidget {
  final Text text;

  const BlinkText({super.key, required this.text});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final timerState = ref.watch(timerProvider);
    final even = timerState.flag;
    final opacity = even ? 0.0 : 1.0;

    return Opacity(
      opacity: opacity,
      child: text,
    );
  }
}

class CustomTextField extends ConsumerWidget {
  final int index;
  final String text;
  // final TextEditingController controller;
  // final FocusNode focusNode;
  // final bool first;
  final bool completed;

  const CustomTextField({
    super.key,
    required this.index,
    required this.text,
    // required this.controller,
    // required this.focusNode,
    // this.first = false,
    this.completed = false,
  });

  double _calcWidth(String text, TextStyle style) {
    final textPainter = TextPainter(
      text: TextSpan(text: text, style: style),
      maxLines: 1,
      textDirection: TextDirection.ltr,
    )..layout(minWidth: 0);
    // textPainter.layout();
    return textPainter.size.width;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final customThemeData = ref.watch(customThemDataProvider);
    final customTheme = customThemeData.problemCard;
    final textFieldList = ref.watch(textFieldValueListProvider);
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final style = completed
        ? customTheme.englishAnswerTextStyle
        : customTheme.englishProblemTextStyle;
    final decoration = completed
        ? customTheme.englishAnswerDecoration
        : customTheme.englishProblemDecoration;

    final hasFocus = index == textFieldList.index;
    // print('build EnglishText');

    final textWidth = _calcWidth(text, style) * 1.06;
    final curosrWidth = _calcWidth('|', customTheme.cusrotTextStyle);
    print('textWidth: $textWidth');
    final width =
        completed ? textWidth + 12 : max(50.0, textWidth + curosrWidth + 12);
    final widget = completed
        ? Container(
            padding: EdgeInsets.only(left: 0, right: 10),
            child: DecoratedBox(
                decoration: decoration, child: Text(text, style: style)),
          )
        : Container(
            width: width,
            padding: EdgeInsets.only(left: 0, right: 10),
            child: DecoratedBox(
              decoration: decoration,
              child: Row(
                children: [
                  Text(text, style: style),
                  Opacity(
                    opacity: hasFocus && !completed ? 1.0 : 0.0,
                    child: BlinkText(
                      text: Text(
                        '|',
                        style: customThemeData.problemCard.cusrotTextStyle,
                      ),
                    ),
                  ),
                ],
              ),
            ),
          );
    return GestureDetector(
      onTap: () {
        print('ontap');
        textFieldListNotifier.setIndex(index);
        // textFieldListNotifier.setPosition(
        //     index, controllerList[index].selection.baseOffset);
      },
      child: widget,
    );
  }
}

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

  List<Widget> _buildEnglishTexts(TextStyle plainTextStyle) {
    List<Widget> englishTexts = [];
    var index = 0;
    var firstProblem = true;
    for (final english in problem.englishList) {
      if (english.isProblem) {
        englishTexts.add(
          CustomTextField(
            index: index,
            text: controllerList[index].text,
            completed: completedList[index],
          ),
        );
        // print('${english.text} == ${controllerList[index].text}');
        // englishTexts.add(ProblemTextField(
        //   index: index,
        //   englishText: english.text,
        //   controller: controllerList[index],
        //   focusNode: focusNodeList[index],
        //   completed: completedList[index],
        //   first: firstProblem,
        // ));
        firstProblem = false;
        index++;
      } else {
        englishTexts.add(PlainText(
          text: english.text,
          style: plainTextStyle,
        ));
      }
    }

    return englishTexts;
  }

  List<Widget> _buildTranslationTexts(
    TextStyle plainTextStyle,
    TextStyle problemTextStyle,
  ) {
    List<Widget> translationTexts = [];
    for (final translation in problem.translationList) {
      final style = translation.isProblem ? problemTextStyle : plainTextStyle;
      translationTexts.add(PlainText(text: translation.text, style: style));
    }
    return translationTexts;
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final customThemeData = ref.watch(customThemDataProvider);

    List<Widget> englishTexts = _buildEnglishTexts(
      customThemeData.problemCard.englishPlainTextStyle,
    );
    List<Widget> translationTexts = _buildTranslationTexts(
      customThemeData.problemCard.translationPlainTextStyle,
      customThemeData.problemCard.translationProblemTextStyle,
    );

    return Card(
      // color: Colors.red,
      color: Colors.transparent,
      child: Column(
        mainAxisSize: MainAxisSize.min,
        // mainAxisSize: MainAxisSize.min, //
        children: [
          DecoratedBox(
            decoration: BoxDecoration(
              color: Colors.green,
              borderRadius: BorderRadius.vertical(
                  top: Radius.circular(16.0)), // 上部の角を丸く設定
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
              color: Colors.blue,
              borderRadius: BorderRadius.vertical(
                  bottom: Radius.circular(16.0)), // 上部の角を丸く設定
            ),
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: SizedBox(
                width: double.infinity,
                height: 150,
                // color: Colors.green,
                child: Wrap(children: englishTexts),
              ),
            ),
          ),
        ],
      ),
      // ),
      // child: Container(
      //   alignment: Alignment.topLeft,
      //   // height: 100.0,
      //   width: double.infinity,
      //   color: Colors.blueGrey[50],
      //   // color: Colors.red,
      //   // padding: EdgeInsets.all(15),
      //   child: Column(
      //     children: [
      //       Container(color: Colors.white, child: ),
      //       SizedBox(height: 10),
      //       Wrap(children: translationTexts),
      //     ],
      //   ),
      // ),
    );

    // return Card(

    //   child: Container(
    //     alignment: Alignment.topLeft,
    //     // height: 100.0,
    //     width: double.infinity,
    //     color: Colors.blueGrey[50],
    //     // color: Colors.red,
    //     // padding: EdgeInsets.all(15),
    //     child: Column(
    //       children: [
    //         Container(color: Colors.white, child: Wrap(children: englishTexts)),
    //         SizedBox(height: 10),
    //         Wrap(children: translationTexts),
    //       ],
    //     ),
    //   ),
    // );
  }
}
