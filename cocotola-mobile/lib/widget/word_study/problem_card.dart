import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/widget/word_study/plain_text.dart';
import 'package:mobile/widget/word_study/problem_text_field.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'dart:async';

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

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final textFieldList = ref.watch(textFieldValueListProvider);
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final color = completed ? Colors.red : Colors.black;

    final hasFocus = textFieldList.index;
    // print('build EnglishText');
    return GestureDetector(
      onTap: () {
        print('ontap');
        textFieldListNotifier.setIndex(index);
        // textFieldListNotifier.setPosition(
        //     index, controllerList[index].selection.baseOffset);
      },
      child: SizedBox(
        width: 100,
        child: Container(
          decoration: BoxDecoration(color: Colors.white),
          padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
          child: Row(
            children: [
              Text(
                text,
                style: TextStyle(color: color),
              ),
              Opacity(
                opacity: hasFocus == index ? 1.0 : 0.0,
                child: BlinkText(
                  text: Text(
                    '|',
                    style: TextStyle(color: Colors.green),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
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

  List<Widget> _buildEnglishTexts() {
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
