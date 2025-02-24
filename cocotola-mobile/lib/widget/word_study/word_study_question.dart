import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/keyboard.dart';
import 'package:mobile/widget/word_study/problem_card.dart';

class WordStudyQuestion extends ConsumerWidget {
  const WordStudyQuestion({super.key});

  Widget _buidProblemCard(AsyncValue<TextFieldValueList> textFieldValueList) {
    switch (textFieldValueList) {
      case AsyncData(:final value):
        final problem = value.problem;
        final texts = value.texts.map((e) => e.text).toList();
        final completedList = value.texts.map((e) => e.completed).toList();
        return ProblemCard(
          problem: problem,
          texts: texts,
          completedList: completedList,
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
    print('WordStudyQuestion build');
    final screenWidth = MediaQuery.of(context).size.width;
    print('Screen width: $screenWidth');
    final screenHeight = MediaQuery.of(context).size.height;
    print('Screen height: $screenHeight');

    final wordStudyStatusNotifier = ref.read(wordStudyStatusProvider.notifier);
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);
    final textFieldValueList = ref.watch(textFieldValueListProvider);

    ref.listen(textFieldValueListProvider, (prev, next) {
      switch (next) {
        case AsyncData(:final value):
          if (value.allCompleted) {
            wordStudyStatusNotifier.setAnswerStatus();
          }
          break;
        case AsyncError(:final error):
          break;
        case AsyncLoading():
          break;
        default:
          break;
      }
    });

    final problemCard = _buidProblemCard(textFieldValueList);

    final bottomHeight = screenHeight * 0.3;
    final bottom = SizedBox(
      height: bottomHeight,
      child: Column(
        children: [
          const Spacer(),
          Keyboard(
            onPresskey: textFieldListNotifier.addText,
            onPressBackspace: textFieldListNotifier.backspace,
          ),
        ],
      ),
    );

    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Expanded(
            child: Center(
              child: Padding(
                padding: const EdgeInsets.all(8.0),
                child: problemCard,
              ),
            ),
          ),
          bottom,
        ],
      ),
    );
  }
}
