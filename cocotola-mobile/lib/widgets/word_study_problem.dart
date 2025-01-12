import 'package:flutter/material.dart';
import 'package:mobile/widgets/english_text.dart';

class WordStudyProblem extends StatefulWidget {
  final List<EnglishText> englishTexts;
  final List<String> japaneseTexts;
  final void Function() onCompleteWord;

  const WordStudyProblem({
    super.key,
    required this.englishTexts,
    required this.japaneseTexts,
    required this.onCompleteWord,
  });

  @override
  State<WordStudyProblem> createState() => _WordStudyProblemState();
}

class _WordStudyProblemState extends State<WordStudyProblem> {
  late List<Widget> englishTexts;

  @override
  Widget build(BuildContext context) {
    englishTexts = widget.englishTexts.map((englishText) {
      if (englishText.isProblem) {
        return EnglishBlankTextWidget(
          englishText: englishText.text,
          controller: englishText.controller,
          focusNode: englishText.focusNode,
          onCompleted: widget.onCompleteWord,
        );
      }
      return EnglishPlainTextWidget(englishText: englishText.text);
    }).toList();
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
