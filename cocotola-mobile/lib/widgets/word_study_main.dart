import 'package:flutter/material.dart';

class WordStudyMain extends StatefulWidget {
  const WordStudyMain({super.key});

  @override
  State<WordStudyMain> createState() => _WordStudyMainState();
}

class _WordStudyMainState extends State<WordStudyMain> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Word Study'),
      ),
      body: SafeArea(
        child: Column(
          children: [
            const Center(
              child: Text(
                'Word Study',
                style: TextStyle(fontSize: 24),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
