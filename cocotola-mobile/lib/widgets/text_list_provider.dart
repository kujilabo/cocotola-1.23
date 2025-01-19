import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:riverpod_annotation/riverpod_annotation.dart';

class TextFieldValue {
  final String text;
  final int position;
  final bool completed;
  const TextFieldValue(
      {required this.text, required this.position, required this.completed});
}

class TextFieldValueList {
  final List<TextFieldValue> texts;
  final int index;
  const TextFieldValueList({required this.texts, required this.index});
}

class TextFieldValueListNotifier extends Notifier<TextFieldValueList> {
  @override
  TextFieldValueList build() {
    final texts = List.generate(
        10, (index) => TextFieldValue(text: '', position: 0, completed: false));
    return TextFieldValueList(texts: texts, index: 0);
  }

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

    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: newText,
          position: newPosition,
          completed: currTextField.completed),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(texts: texts, index: state.index);
  }

  void setPosition(int index, int position) {
    print('position: $position');
    final currText = state.texts[index];

    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: currText.text,
          position: position,
          completed: currText.completed),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(texts: texts, index: state.index);
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

    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: newText,
          position: currPosition - 1,
          completed: currTextField.completed),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(texts: texts, index: state.index);
  }

  void setIndex(int index) {
    state = TextFieldValueList(texts: state.texts, index: index);
  }

  void setComplete(int index) {
    final currTextField = state.texts[index];
    final texts = [
      ...state.texts.sublist(0, index),
      TextFieldValue(
          text: currTextField.text,
          position: currTextField.position,
          completed: true),
      ...state.texts.sublist(index + 1),
    ];
    state = TextFieldValueList(texts: texts, index: index);
  }
}

final textFieldValueListProvider =
    NotifierProvider<TextFieldValueListNotifier, TextFieldValueList>(
        TextFieldValueListNotifier.new);
