import 'package:flutter/material.dart';

class EnglishText {
  final String text;
  final bool isProblem;
  final TextEditingController? controller;
  final FocusNode? focusNode;

  EnglishText(this.text,
      {this.isProblem = false, this.controller, this.focusNode});
}

class EnglishBlankText {
  final String text;

  EnglishBlankText({required this.text});
}

class EnglishPlainTextWidget extends StatelessWidget {
  final String englishText;

  const EnglishPlainTextWidget({super.key, required this.englishText});

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: EdgeInsets.only(
          left: 0,
          right: 10,
        ),
        child: Text(
          englishText,
          style: TextStyle(fontSize: 24),
        ));
  }
}

class EnglishBlankTextWidget extends StatefulWidget {
  final String englishText;
  final TextEditingController? controller;
  final FocusNode? focusNode;
  final void Function() onCompleted;

  const EnglishBlankTextWidget({
    super.key,
    required this.englishText,
    this.controller,
    this.focusNode,
    required this.onCompleted,
  });

  @override
  State<EnglishBlankTextWidget> createState() => _EnglishBlankTextWidgetState();
}

class _EnglishBlankTextWidgetState extends State<EnglishBlankTextWidget> {
  // late TextEditingController _controller;
  // late TextSelection _selection;

  @override
  void initState() {
    super.initState();
    // _controller = TextEditingController()..addListener(_onSelectionChanged);
    // _selection = _controller.selection;
  }

  @override
  void dispose() {
    // _controller.removeListener(_onSelectionChanged);
    super.dispose();
  }

  // void _onSelectionChanged() {
  //   setState(() {
  //     // update selection on change (updating position too)
  //     _selection = _controller.selection;
  //   });
  //   print('Cursor position: ${_selection.base.offset}'); // print position
  // }

  @override
  Widget build(BuildContext context) {
    if (widget.controller != null) {
      widget.controller!.addListener(() {
        print('zzzzzzzzzzzzzzzzzzz');
        print(widget.controller!.text);
      });
    }
    print('xxxxxxxxxxxxxxxxxxxxxxx');
    return Container(
      width: 100,
      child: TextField(
        focusNode: widget.focusNode,
        controller: widget.controller,
        // decoration: InputDecoration(
        //   hintText: widget.englishText,
        // ),
      ),
    );
  }
}
