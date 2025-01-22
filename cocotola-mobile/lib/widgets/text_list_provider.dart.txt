import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:mobile/model/word_problem.dart';
import 'package:mobile/provider/problem_provider.dart';

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
  final List<TextFieldValue> texts;
  final int index;
  final int numProblems;
  final bool allCompleted;
  const TextFieldValueList({
    required this.texts,
    required this.index,
    required this.numProblems,
    required this.allCompleted,
  });
}

class TextFieldValueListNotifier extends Notifier<TextFieldValueList> {
  // final ProblemWordStudy problemRepository;
  // TextFieldValueListNotifier(this.problemRepository);
  @override
  TextFieldValueList build() {
    final problem = ref.watch(problemProvider);

    List<TextFieldValue> texts = [];
    var numProblems = 0;
    for (var i = 0; i < problem.englishList.length; i++) {
      final english = problem.englishList[i];
      if (english.isProblem) {
        texts.add(TextFieldValue(
            text: '', answer: english.text, position: 0, completed: false));
        print('problem: ${english.text}');
        numProblems++;
      }
    }
    final numEmpty = 10 - texts.length;
    for (var i = 0; i < numEmpty; i++) {
      texts.add(
          TextFieldValue(text: '', answer: '', position: 0, completed: false));
    }

    // final texts = List.generate(
    //     10,
    //     (index) => TextFieldValue(
    //         text: '', answer: '', position: 0, completed: false));
    return TextFieldValueList(
      texts: texts,
      index: 0,
      numProblems: numProblems,
      allCompleted: false,
    );
  }

  // void setAnswer(int index, String answer) {}
  void addText(String text) {
    final index = state.index;
    final currTextField = state.texts[index];
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
    var completed = newText == currTextField.answer;
    var newIndex = state.index;
    if (completed) {
      var numCorrect = 1;
      for (var i = 0; i < state.numProblems; i++) {
        newIndex = (i + index + 1) % state.numProblems;
        if (state.texts[newIndex].completed) {
          numCorrect++;
        } else {
          break;
        }
      }
      if (numCorrect == state.numProblems) {
        allCompleted = true;
      }
    }
    print('completed: $completed, ${currTextField.answer}, ${newIndex}');
    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: newText,
          answer: currTextField.answer,
          position: newPosition,
          completed: completed),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(
      texts: texts,
      index: newIndex,
      numProblems: state.numProblems,
      allCompleted: allCompleted,
    );
  }

  void setPosition(int index, int position) {
    print('position: $position');
    final currTextField = state.texts[index];
    if (currTextField.completed) {
      return;
    }

    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: currTextField.text,
          answer: currTextField.answer,
          position: position,
          completed: currTextField.completed),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(
      texts: texts,
      index: state.index,
      numProblems: state.numProblems,
      allCompleted: state.allCompleted,
    );
  }

  void backspace() {
    final index = state.index;
    final currTextField = state.texts[index];
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

    var completed = newText == currTextField.answer;
    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: newText,
          answer: currTextField.answer,
          position: currPosition - 1,
          completed: completed),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(
      texts: texts,
      index: state.index,
      numProblems: state.numProblems,
      allCompleted: state.allCompleted,
    );
  }

  void setIndex(int index) {
    state = TextFieldValueList(
      texts: state.texts,
      index: index,
      numProblems: state.numProblems,
      allCompleted: state.allCompleted,
    );
  }

  // void setComplete(int index) {
  //   final currTextField = state.texts[index];
  //   final texts = [
  //     ...state.texts.sublist(0, index),
  //     TextFieldValue(
  //         text: currTextField.text,
  //         answer: currTextField.answer,
  //         position: currTextField.position,
  //         completed: true),
  //     ...state.texts.sublist(index + 1),
  //   ];
  //   state = TextFieldValueList(
  //       texts: texts, index: index, numProblems: state.numProblems);
  // }
}

final textFieldValueListProvider =
    NotifierProvider<TextFieldValueListNotifier, TextFieldValueList>(
        TextFieldValueListNotifier.new);
// final textFieldValueListProvider =
//     NotifierProvider<TextFieldValueListNotifier, TextFieldValueList>((Ref ref) {
//   final problem = ref.read(problemProvider);
//   return TextFieldValueListNotifier(problem);
// });

// @riverpod
// TextFieldValueList textFieldValueList(Ref ref){
//   final problem =ref.watch(problemProvider);
//   return 
// }