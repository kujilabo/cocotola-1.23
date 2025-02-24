import 'package:flutter/material.dart';

class PlainText extends StatelessWidget {
  final String text;
  final TextStyle style;

  const PlainText({required this.text, required this.style, super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      // color: Colors.purple,
      child: Padding(
        padding: const EdgeInsets.only(right: 10),
        child: Text(text, style: style),
      ),
    );
  }
}
