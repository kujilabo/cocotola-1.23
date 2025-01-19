import 'package:flutter/src/widgets/framework.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widgets/keyboard.dart';
import 'package:mobile/widgets/text_list_provider.dart';
import 'package:mobile/widgets/editor_screen.dart';

class Editor extends ConsumerWidget {
  final _focusNode = FocusNode();
  final _controller = TextEditingController();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('Editor build');
    final textFieldListNotifier = ref.read(textFieldValueListProvider.notifier);

    return Scaffold(
      appBar: AppBar(
        title: Text('Editor'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            EditorScreen(),
            // TextField(
            //   autofocus: true,
            //   focusNode: _focusNode,
            //   controller: _controller,
            //   readOnly: false,
            //   style: TextStyle(color: Colors.black),
            // ),
            Keyboard(
              onPresskey: (String text) {
                textFieldListNotifier.addText(text);
              },
              onPressBackspace: () {
                textFieldListNotifier.backspace();
              },
            ),
          ],
        ),
      ),
    );
  }

  void _onPressBackspace() {}
}
