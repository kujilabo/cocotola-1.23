import 'dart:async';
import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/custom_theme_data_provider.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';

class TimerState {

  const TimerState({required this.flag, required this.timer});
  final bool flag;
  final Timer timer;
}

class TimerRepository extends Notifier<TimerState> {
  @override
  TimerState build() {
    ref.onDispose(() => state.timer.cancel());
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

  const BlinkText({required this.text, super.key});
  final Text text;

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

class CustomProblemTextField extends ConsumerWidget {

  const CustomProblemTextField({
    required this.index, required this.text, super.key,
  });
  final int index;
  final String text;

  double _calcWidth(String text, TextStyle style) {
    final textPainter = TextPainter(
      text: TextSpan(text: text, style: style),
      maxLines: 1,
      textDirection: TextDirection.ltr,
    )..layout();
    return textPainter.size.width;
  }

  Widget _buildFocus(AsyncValue<bool> hasFocus, TextStyle style) {
    switch (hasFocus) {
      case AsyncData(:final value):
        return Opacity(
          opacity: value ? 1.0 : 0.0,
          child: BlinkText(text: Text('|', style: style)),
        );
      case AsyncError(:final error):
        return Text('Error: $error');
      case AsyncLoading():
        return const CircularProgressIndicator();
      default:
        return const CircularProgressIndicator();
    }
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final customThemeData = ref.watch(customThemDataProvider);
    final customTheme = customThemeData.problemCard;
    final textFieldList = ref.watch(textFieldValueListProvider);
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final style = customTheme.englishProblemTextStyle;
    final decoration = customTheme.englishProblemDecoration;
    final hasFocus = textFieldList.whenData((v) => index == v.index);

    final textWidth = _calcWidth(text, style) * 1.06;
    final curosrWidth = _calcWidth('|', customTheme.cusrotTextStyle);
    print('textWidth: $textWidth');
    final width = max(50.0, textWidth + curosrWidth + 12);

    return GestureDetector(
      onTap: () {
        print('ontap');
        textFieldListNotifier.setIndex(index);
        // textFieldListNotifier.setPosition(
        //     index, controllerList[index].selection.baseOffset);
      },
      child: Container(
        width: width,
        padding: const EdgeInsets.only(right: 10),
        child: DecoratedBox(
          decoration: decoration,
          child: Row(
            children: [
              Text(text, style: style),
              _buildFocus(hasFocus, customTheme.cusrotTextStyle),
            ],
          ),
        ),
      ),
    );
  }
}
