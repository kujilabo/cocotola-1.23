import 'package:flutter/material.dart';

class Keyboard extends StatelessWidget {
  final void Function(String) onPresskey;
  final void Function() onPressBackspace;

  const Keyboard({
    super.key,
    required this.onPresskey,
    required this.onPressBackspace,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: Column(
        spacing: 5,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            spacing: 5,
            children: _buildButtons(
              ['q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'],
            ),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            spacing: 5,
            children: _buildButtons(
              ['a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'],
            ),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            spacing: 5,
            children: _buildButtons(
                  ['z', 'x', 'c', 'v', 'b', 'n', 'm'],
                ) +
                [_buildButton('⌫', onPressed: onPressBackspace)],
          ),
        ],
      ),
    );
  }

  List<Widget> _buildButtons(List<String> characters) {
    return characters.map((c) => _buildButton(c)).toList();
  }

  Widget _buildButton(String text, {VoidCallback? onPressed}) {
    return TextButton(
      style: TextButton.styleFrom(
        fixedSize: const Size(20, 20),
        foregroundColor: Colors.white,
        backgroundColor: Colors.blue,
      ),
      onPressed: onPressed ?? () => onPresskey(text),
      child: Text(text),
    );
  }
}
