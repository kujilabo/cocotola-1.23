import 'package:flutter/src/widgets/framework.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widgets/text_list_provider.dart';

class EditorScreen extends ConsumerWidget {
  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final _focusNode0 = FocusNode();
    final _focusNode1 = FocusNode();
    final _controller0 = TextEditingController();
    final _controller1 = TextEditingController();
    final textFieldValueList = ref.watch(textFieldValueListProvider);
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);

    _focusNode0.addListener(() {
      if (_focusNode0.hasFocus) {
        textFieldListNotifier.setIndex(0);
        print('_focusNode0 has focus $_controller0.selection');
        textFieldListNotifier.setPosition(0, _controller0.selection.baseOffset);
      } else {
        print('_focusNode0 doen\'t have focus');
      }
    });
    _focusNode1.addListener(() {
      if (_focusNode1.hasFocus) {
        textFieldListNotifier.setIndex(1);
        print('_focusNode1 has focus $_controller1.selection');
      } else {
        print('_focusNode1 doen\'t have focus');
      }
    });

    print('EditorScreen build');
    _controller0.text = textFieldValueList.texts[0].text;
    _controller1.text = textFieldValueList.texts[1].text;
    final index = textFieldValueList.index;
    if (index == 0) {
      _controller0.selection = TextSelection.fromPosition(
          TextPosition(offset: textFieldValueList.texts[index].position));
    } else if (index == 1) {
      _controller1.selection = TextSelection.fromPosition(
          TextPosition(offset: textFieldValueList.texts[index].position));
    }
    return Column(
      children: [
        TextField(
          autofocus: true,
          focusNode: _focusNode0,
          controller: _controller0,
          readOnly: false,
          style: TextStyle(color: Colors.black),
        ),
        TextField(
          autofocus: true,
          focusNode: _focusNode1,
          controller: _controller1,
          readOnly: false,
          style: TextStyle(color: Colors.black),
        ),
      ],
    );
  }
}
