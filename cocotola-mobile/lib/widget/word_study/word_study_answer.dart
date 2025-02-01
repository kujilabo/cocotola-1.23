import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/provider/problem_provider.dart';
import 'package:mobile/widget/word_study/problem_card.dart';

class WordStudyAnswer extends ConsumerWidget {
  const WordStudyAnswer({super.key});

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
    final double screenHeight = MediaQuery.of(context).size.height;
    final problemNotifier = ref.read(problemProvider.notifier);
    final problemWithStatus = ref.watch(problemProvider);
    final wordStudyStatusNotifier = ref.read(wordStudyStatusProvider.notifier);
    final textFieldValueList = ref.watch(textFieldValueListProvider);

    // final problem = await problemWithStatus.then((value) => value.currentProblem);
    // final problem = problemWithStatus.then((value) => value.currentProblem);
    // final problem = problemWithStatus.currentProblem;

    // final texts = textFieldValueList.texts.map((e) => e.text).toList();
    // final texts =  textFieldValueList
    //     .then((value) => value.texts.map((e) => e.text).toList());
    // final completedList =
    //     textFieldValueList.texts.map((e) => e.completed).toList();
    // final completedList = textFieldValueList
    //     .then((value) => value.texts.map((e) => e.completed).toList());

    var problemCard = _buidProblemCard(textFieldValueList);

    final bottomHeight = screenHeight * 0.3;

    final bottom = SizedBox(
      height: bottomHeight,
      child: Column(
        children: [
          Spacer(),
          Row(
            children: [
              Expanded(
                child: ElevatedButton(
                  onPressed: () {
                    problemNotifier.next();
                    wordStudyStatusNotifier.setQuestionStatus();
                  },
                  child: Text('Check'),
                ),
              ),
              Expanded(
                child: ElevatedButton(
                  onPressed: () {
                    problemNotifier.next();
                    wordStudyStatusNotifier.setQuestionStatus();
                  },
                  child: Text('Next'),
                ),
              ),
            ],
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
          // Spacer(), // 余
          bottom,
        ],
      ),
    );
  }
}
