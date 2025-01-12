import 'package:flutter/material.dart';
import 'package:mobile/widgets/english_text.dart';

class WordStudyProblem extends StatefulWidget {
  final List<EnglishText> englishTexts;
  final List<String> japaneseTexts;
  final void Function(int) onCompletedWord;

  const WordStudyProblem({
    super.key,
    required this.englishTexts,
    required this.japaneseTexts,
    required this.onCompletedWord,
  });

  @override
  State<WordStudyProblem> createState() => _WordStudyProblemState();
}

class _WordStudyProblemState extends State<WordStudyProblem> {
  late List<Widget> englishTexts;

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    englishTexts = [];
    var index = 0;
    var length = widget.englishTexts.length;
    print('length: $length');
    for (var i = 0; i < length; i++) {
      print('i: $i');
      final englishText = widget.englishTexts[i];
      if (englishText.isProblem) {
        englishTexts.add(EnglishBlankTextWidget(
          index: index,
          englishText: englishText.text,
          controller: englishText.controller,
          focusNode: englishText.focusNode,
          onCompleted: widget.onCompletedWord,
          first: englishText.first,
        ));
        index++;
      } else {
        englishTexts.add(EnglishPlainTextWidget(
          englishText: englishText.text,
        ));
      }
    }
    return Card(
      child: Container(
        alignment: Alignment.topLeft,
        // height: 100.0,
        width: double.infinity,
        // color: Colors.red,
        padding: EdgeInsets.all(15),
        child: Column(
          children: [
            Wrap(children: englishTexts),
            SizedBox(height: 10),
            Wrap(children: englishTexts),
          ],
        ),
      ),
    );
  }
}
