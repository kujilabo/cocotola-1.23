import 'package:flutter/material.dart';
import 'package:mobile/widget/word_study/menu_word_study.dart';
import 'package:mobile/widget/word_study/word_study.dart';
// import 'package:mobile/widgets/editor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

void main() {
  runApp(
    ProviderScope(
      child: MaterialApp(
        home: const MenuWordStudy(),
        // home: WordStudyMain(),
        // home: WordStudyMain(),
      ),
    ),
  );
}
