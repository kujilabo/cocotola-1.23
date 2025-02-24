import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/provider/problem_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

class TextFieldValue {
  final String text;
  final String answer;
  final int position;
  final bool completed;
  const TextFieldValue(
      {required this.text,
      required this.answer,
      required this.position,
      required this.completed});
}

class TextFieldValueList {
  final WordProblem problem;
  final List<TextFieldValue> texts;
  final int index;
  final int numProblems;
  final bool allCompleted;
  const TextFieldValueList({
    required this.problem,
    required this.texts,
    required this.index,
    required this.numProblems,
    required this.allCompleted,
  });
}

class TextFieldValueListNotifier extends AsyncNotifier<TextFieldValueList> {
  @override
  Future<TextFieldValueList> build() async {
    final problemWithStatus = await ref.watch(problemProvider.future);
    final problem = problemWithStatus.currentProblem;
    final texts = <TextFieldValue>[];

    for (var i = 0; i < problem.englishList.length; i++) {
      final english = problem.englishList[i];
      if (english.isProblem) {
        texts.add(TextFieldValue(
            text: '', answer: english.text, position: 0, completed: false));
        print('problem: ${english.text}');
      }
    }
    return TextFieldValueList(
      problem: problem,
      texts: texts,
      index: 0,
      numProblems: problem.getNumProblems(),
      allCompleted: false,
    );
  }

  void addText(String text) {
    final currentState = state.value!;
    final index = currentState.index;
    final currTextField = currentState.texts[index];
    if (currTextField.completed) {
      return;
    }

    String newText;
    int newPosition;

    if (currTextField.text.isEmpty) {
      newText = text;
      newPosition = 1;
    } else {
      final currPosition = currTextField.position;
      final currText = currTextField.text;
      final text0 = currText.substring(0, currPosition);
      final text1 = currText.substring(currPosition, currText.length);
      newText = text0 + text + text1;
      newPosition = currTextField.position + 1;
    }

    var allCompleted = false;
    final completed = newText == currTextField.answer;
    var newIndex = currentState.index;
    if (completed) {
      var numCorrect = 1;
      for (var i = 0; i < currentState.numProblems; i++) {
        newIndex = (i + index + 1) % currentState.numProblems;
        if (currentState.texts[newIndex].completed) {
          numCorrect++;
        } else {
          break;
        }
      }
      if (numCorrect == currentState.numProblems) {
        allCompleted = true;
      }
    }
    print('completed: $completed, ${currTextField.answer}, $newIndex');
    final texts = [
      ...currentState.texts.sublist(0, index),
      TextFieldValue(
          text: newText,
          answer: currTextField.answer,
          position: newPosition,
          completed: completed),
      ...currentState.texts.sublist(index + 1),
    ];
    state = AsyncValue.data(
      TextFieldValueList(
        problem: currentState.problem,
        texts: texts,
        index: newIndex,
        numProblems: currentState.numProblems,
        allCompleted: allCompleted,
      ),
    );
  }

  void setPosition(int index, int position) {
    final currentState = state.value!;
    print('position: $position');
    final currTextField = currentState.texts[index];
    if (currTextField.completed) {
      return;
    }

    final texts = [
      ...currentState.texts.sublist(0, index),
      TextFieldValue(
          text: currTextField.text,
          answer: currTextField.answer,
          position: position,
          completed: currTextField.completed),
      ...currentState.texts.sublist(index + 1),
    ];
    state = AsyncValue.data(
      TextFieldValueList(
        problem: currentState.problem,
        texts: texts,
        index: currentState.index,
        numProblems: currentState.numProblems,
        allCompleted: currentState.allCompleted,
      ),
    );
  }

  void backspace() {
    final currentState = state.value!;
    final index = currentState.index;
    final currTextField = currentState.texts[index];
    if (currTextField.completed) {
      return;
    }
    if (currTextField.text.isEmpty || currTextField.position == 0) {
      return;
    }

    final currPosition = currTextField.position;
    final currText = currTextField.text;
    final text0 = currText.substring(0, currPosition - 1);
    final text1 = currText.substring(currPosition, currText.length);
    final newText = text0 + text1;

    final completed = newText == currTextField.answer;
    final texts = [
      ...currentState.texts.sublist(0, index),
      TextFieldValue(
          text: newText,
          answer: currTextField.answer,
          position: currPosition - 1,
          completed: completed),
      ...currentState.texts.sublist(index + 1),
    ];
    state = AsyncValue.data(
      TextFieldValueList(
        problem: currentState.problem,
        texts: texts,
        index: currentState.index,
        numProblems: currentState.numProblems,
        allCompleted: currentState.allCompleted,
      ),
    );
  }

  void setIndex(int index) {
    final currentState = state.value!;
    state = AsyncValue.data(
      TextFieldValueList(
        problem: currentState.problem,
        texts: currentState.texts,
        index: index,
        numProblems: currentState.numProblems,
        allCompleted: currentState.allCompleted,
      ),
    );
  }
}

final textFieldValueListProvider =
    AsyncNotifierProvider<TextFieldValueListNotifier, TextFieldValueList>(
        TextFieldValueListNotifier.new);
