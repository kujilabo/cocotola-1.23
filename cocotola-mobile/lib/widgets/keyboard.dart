import 'package:flutter/material.dart';

class Keyboard extends StatefulWidget {
  final List<TextEditingController> controllers;

  const Keyboard({super.key, required this.controllers});

  @override
  State<Keyboard> createState() => _KeyboardState();
}

class _KeyboardState extends State<Keyboard> {
  late List<TextEditingController> _controllers;
  late TextSelection _selection;

  @override
  void initState() {
    super.initState();
    // for (var controller in widget.controllers) {
    //   controller.addListener(_onSelectionChanged);
    // }
    // _controller = widget.controller.addListener(_onSelectionChanged);
    // _selection = _controller.selection;
  }

  @override
  void dispose() {
    // for (var controller in _controllers) {
    //   controller.removeListener(_onSelectionChanged);
    // }

    super.dispose();
  }

  void _onSelectionChanged() {
    // setState(() {
    //   if (_controller != null) {
    //     // update selection on change (updating position too)
    //     _selection = _controller!.selection;
    //     print('Cursor position: ${_selection!.base.offset}'); // print position
    //   }
    // });
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.grey,
      child: Column(
        spacing: 5,
        children: [
          Row(
            spacing: 5,
            children: [
              _buildButton('1'),
              _buildButton('2'),
              _buildButton('3'),
            ],
          ),
          Row(
            children: [
              _buildButton('1'),
              _buildButton('1'),
              _buildButton('1'),
            ],
          ),
          Row(
            children: [
              _buildButton('1'),
              _buildButton('1'),
              _buildButton('⌫', onPressed: _backspace),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildButton(String text, {VoidCallback? onPressed}) {
    return Expanded(
      child: TextButton(
        onPressed: onPressed ?? () => _input(text),
        // style: TextButton.styleFrom(
        //   foregroundColor: Colors.white,
        //   backgroundColor: Colors.blue,
        // ),
        child: Text(text),
      ),
    );
  }

// 3
  void _input(String text) {
    // var position = _selection.base.offset; // gets position of cursor
    // var value = _controller.text; // text in our textfield

    // if (value.isNotEmpty) {
    //   // 1) suffix: the string
    //   var suffix = value.substring(position, value.length);
    //   // from the position of the cursor to the end of the text in the controller

    //   // 2) value.substring gets
    //   value = value.substring(0, position) + text + suffix;
    //   // a new string from start of the string in our textfield, appends the new input to our
    //   // new string and appends the suffix to it.

    //   // 3) set our controller text to the gotten value
    //   _controller.text = value;
    //   // 4) update selection
    //   // to update our position.
    //   _controller.selection =
    //       TextSelection.fromPosition(TextPosition(offset: position + 1));
    // } else {
    //   // 5) appends controller text and new input
    //   value = _controller.text + text;
    //   // and assigns to value
    //   // 6) set our controller text to the gotten value
    //   _controller.text = value;
    //   // 7) since this is the first input
    //   // set position of cursor to 1, so the cursor is placed at the end
    //   _controller.selection =
    //       TextSelection.fromPosition(const TextPosition(offset: 1));
    // }
  }

  void _backspace() {
    // var position = _selection.base.offset; // cursor position
    // final value = _controller.text; // string in out textfield

    // // 1) only erase when string in textfield is not empty and when position is not zero (at the start)
    // if (value.isNotEmpty && position != 0) {
    //   var suffix = value.substring(
    //       position, value.length); // 2) get string after cursor position
    //   _controller.text = value.substring(0, position - 1) +
    //       suffix; // 3) get string before the cursor and append to
    //   // suffix after removing the last char before the cursor
    //   _controller.selection = TextSelection.fromPosition(
    //       TextPosition(offset: position - 1)); // 4) update the cursor
    // }
  }
}
