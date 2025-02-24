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

  Widget _buildBottom(
    BuildContext context,
    ProblemRepository problemNotifier,
    WordStudyStatusNotifier wordStudyStatusNotifier,
    AsyncValue<ProblemWithStatus> problemWithStatus,
    double bottomHeight,
  ) {
    switch (problemWithStatus) {
      case AsyncData(:final value):
        if (value.hasNext()) {
          return SizedBox(
            height: bottomHeight,
            child: Column(
              children: [
                const Spacer(),
                Row(
                  children: [
                    Expanded(
                      child: ElevatedButton(
                        onPressed: () {
                          problemNotifier.next();
                          wordStudyStatusNotifier.setQuestionStatus();
                        },
                        child: const Text('Check'),
                      ),
                    ),
                    Expanded(
                      child: ElevatedButton(
                        onPressed: () {
                          problemNotifier.next();
                          wordStudyStatusNotifier.setQuestionStatus();
                        },
                        child: const Text('Next'),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          );
        } else {
          return SizedBox(
            height: bottomHeight,
            child: Column(
              children: [
                const Spacer(),
                Row(
                  children: [
                    Expanded(
                      child: ElevatedButton(
                        onPressed: () {
                          problemNotifier.next();
                          wordStudyStatusNotifier.setQuestionStatus();
                        },
                        child: const Text('Check'),
                      ),
                    ),
                    Expanded(
                      child: ElevatedButton(
                        onPressed: () {
                          Navigator.of(context).pop();
                        },
                        child: const Text('Finish'),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          );
        }
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
    final screenHeight = MediaQuery.of(context).size.height;
    final problemNotifier = ref.read(problemProvider.notifier);
    final problemWithStatus = ref.watch(problemProvider);
    final wordStudyStatusNotifier = ref.read(wordStudyStatusProvider.notifier);
    final bottomHeight = screenHeight * 0.3;
    final textFieldValueList = ref.watch(textFieldValueListProvider);

    var problemCard = _buidProblemCard(textFieldValueList);
    var bottom = _buildBottom(
      context,
      problemNotifier,
      wordStudyStatusNotifier,
      problemWithStatus,
      bottomHeight,
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
