import 'package:flutter/material.dart';

class Keyboard extends StatelessWidget {
  const Keyboard({
    required this.onPresskey,
    required this.onPressBackspace,
    super.key,
  });
  final void Function(String) onPresskey;
  final void Function() onPressBackspace;

  @override
  Widget build(BuildContext context) {
    final screenWidth = MediaQuery.of(context).size.width;
    final screenHeight = MediaQuery.of(context).size.height;
    final size = Size(screenWidth / 12, screenHeight / 15);
    return ColoredBox(
      color: Colors.lightGreenAccent,
      child: Padding(
        padding: const EdgeInsets.all(8),
        child: Column(
          spacing: 5,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              spacing: 5,
              children: _buildButtons(
                ['q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'],
                size,
              ),
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              spacing: 5,
              children: _buildButtons(
                ['a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'],
                size,
              ),
            ),
            Row(
              // mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              mainAxisAlignment: MainAxisAlignment.center,
              spacing: 5,
              children: _buildButtons(
                    ['z', 'x', 'c', 'v', 'b', 'n', 'm'],
                    size,
                  ) +
                  [_buildButton('⌫', size, onPressed: onPressBackspace)],
            ),
          ],
        ),
      ),
    );
  }

  List<Widget> _buildButtons(List<String> characters, Size size) {
    return characters.map((c) => _buildButton(c, size)).toList();
  }

  Widget _buildButton(String text, Size size, {VoidCallback? onPressed}) {
    return TextButton(
      style: TextButton.styleFrom(
        padding: EdgeInsets.zero,
        fixedSize: size,
        minimumSize: const Size(20, 20), // 最小サイズを設定
        // maximumSize: Size(20, 20), // 最大サイズを設定
        foregroundColor: Colors.white,
        backgroundColor: Colors.blue,
        tapTargetSize: MaterialTapTargetSize.shrinkWrap, // タップターゲットサイズを縮小
        // visualDensity: VisualDensity.compact,
      ),
      onPressed: onPressed ?? () => onPresskey(text),
      child: Text(
        text,
        // style: TextStyle(fontSize: 20),
      ),
    );
  }
}
