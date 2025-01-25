import 'package:flutter/material.dart';

class PlainText extends StatelessWidget {
  final String text;

  const PlainText({super.key, required this.text});

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: EdgeInsets.only(
          left: 0,
          right: 10,
        ),
        child: Text(
          text,
          style: TextStyle(fontSize: 24),
        ));
  }
}
