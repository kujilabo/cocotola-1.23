import 'package:flutter/material.dart';

class EnglishText {
  final String text;
  final bool isProblem;
  final TextEditingController? controller;
  final FocusNode? focusNode;
  final bool first;

  EnglishText(
    this.text, {
    this.isProblem = false,
    this.controller,
    this.focusNode,
    this.first = false,
  });
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
  final int index;
  final String englishText;
  final TextEditingController? controller;
  final FocusNode? focusNode;
  final bool first;
  final void Function(int) onCompleted;

  const EnglishBlankTextWidget({
    super.key,
    required this.index,
    required this.englishText,
    this.controller,
    this.focusNode,
    this.first = false,
    required this.onCompleted,
  });

  @override
  State<EnglishBlankTextWidget> createState() => _EnglishBlankTextWidgetState();
}

class _EnglishBlankTextWidgetState extends State<EnglishBlankTextWidget> {
  var readOnly = false;
  var color = Colors.black;

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
    print('readonly ' + readOnly.toString());
    if (widget.controller != null) {
      widget.controller!.addListener(() {
        if (widget.controller!.text == widget.englishText) {
          widget.onCompleted(widget.index);
          setState(() {
            readOnly = true;
            color = Colors.red;
          });
        }
        print('zzzzzzzzzzzzzzzzzzz');
        print(widget.controller!.text);
      });
    }
    print('build EnglishText');
    return Container(
      width: 100,
      child: TextField(
        autofocus: widget.first,
        focusNode: widget.focusNode,
        controller: widget.controller,
        readOnly: readOnly,
        style: TextStyle(color: color),
        // onChanged: (text) {
        //   if (text == widget.englishText) {
        //     print('yyyyyyyyyyyyyyyyyyy');
        //     setState(() {
        //       readOnly = true;
        //     });
        //   }
        // },
        // decoration: InputDecoration(
        //   hintText: widget.englishText,
        // ),
      ),
    );
  }
}
