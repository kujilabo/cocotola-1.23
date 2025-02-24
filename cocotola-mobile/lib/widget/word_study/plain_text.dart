import 'package:flutter/material.dart';

class PlainText extends StatelessWidget {
  const PlainText({required this.text, required this.style, super.key});
  final String text;
  final TextStyle style;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(right: 10),
      child: Text(text, style: style),
    );
  }
}
