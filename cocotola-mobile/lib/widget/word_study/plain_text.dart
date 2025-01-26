import 'package:flutter/material.dart';

class PlainText extends StatelessWidget {
  final String text;
  final TextStyle style;

  const PlainText({super.key, required this.text, required this.style});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.purple,
      child: Padding(
        padding: EdgeInsets.only(left: 0, right: 10),
        child: Text(text, style: style),
      ),
    );
  }
}
