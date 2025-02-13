import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class ProblemTextField extends ConsumerWidget {
  final int index;
  final String englishText;
  final TextEditingController controller;
  final FocusNode focusNode;
  final bool first;
  final bool completed;

  const ProblemTextField({
    super.key,
    required this.index,
    required this.englishText,
    required this.controller,
    required this.focusNode,
    this.first = false,
    this.completed = false,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var color = completed ? Colors.red : Colors.black;
    // print('build EnglishText');
    return SizedBox(
      width: 100,
      child: Container(
        padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
        child: TextField(
          autofocus: first,
          focusNode: focusNode,
          controller: controller,
          readOnly: completed,
          style: TextStyle(color: color),
        ),
      ),
    );
  }
}
