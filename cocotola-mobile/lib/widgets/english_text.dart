import 'package:flutter/material.dart';

import 'package:flutter_riverpod/flutter_riverpod.dart';

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

class EnglishBlankTextWidget extends ConsumerWidget {
  final int index;
  final String englishText;
  final TextEditingController? controller;
  final FocusNode? focusNode;
  final bool first;
  final bool completed;
  final void Function() onCompleted;

  const EnglishBlankTextWidget({
    super.key,
    required this.index,
    required this.englishText,
    this.controller,
    this.focusNode,
    this.first = false,
    this.completed = false,
    required this.onCompleted,
  });
  //    if (controller != null) {
  //     controller!.addListener(() {
  //       if (controller!.text ==englishText) {
  //         onCompleted(index);
  //         // setState(() {
  //         //   readOnly = true;
  //         //   color = Colors.red;
  //         // });
  //       }
  //       print(controller!.text);
  //     });
  //   }
  // }

  // var readOnly = false;
  // var color = Colors.black;

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    // print('readonly ' + readOnly.toString());

    // print('build EnglishText');
    return SizedBox(
      width: 100,
      child: TextField(
          autofocus: first,
          focusNode: focusNode,
          controller: controller,
          readOnly: completed //readOnly,
          // style: TextStyle(color: color),
          ),
    );
  }
}
