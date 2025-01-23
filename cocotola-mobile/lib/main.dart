import 'package:flutter/material.dart';
import 'package:mobile/widget/word_study/word_study_top.dart';
import 'package:mobile/widget/word_study/word_study_main.dart';
// import 'package:mobile/widgets/editor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

void main() {
  runApp(
    ProviderScope(
      child: MaterialApp(
        // home: const WordStudyTop(),
        // home: WordStudyMain(),
        home: WordStudyMain(),
      ),
    ),
  );
}
