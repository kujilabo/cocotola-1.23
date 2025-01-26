import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter/material.dart';
import 'package:mobile/widget/keyboard.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/widget/word_study/word_study_answer.dart';
import 'package:mobile/widget/word_study/word_study_question.dart';
import 'package:mobile/widget/word_study/problem_card.dart';

class WordStudy extends ConsumerWidget {
  const WordStudy({super.key});

  Widget _buildMain(WordStudyStatus status) {
    switch (status) {
      case WordStudyStatus.question:
        return WordStudyQuestion();
      case WordStudyStatus.answer:
        return WordStudyAnswer();
      default:
        return Text('aaaaa');
    }
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    print('WordStudy build');
    final wordStudyStatus = ref.watch(wordStudyStatusProvider);

    final header = Text('aaaaaaaaaaaaaaa');
    final main = _buildMain(wordStudyStatus);
    return Scaffold(
      appBar: AppBar(
        title: Text('WordStudy'),
      ),
      body: Container(
        child: Column(
          children: [
            header,
            Expanded(
              child: main, // Expandedウィジェットでmainを包む
            ),
          ],
        ),
      ),
      // body: Column(
      //   children: [
      //     main,
      //   ],
      // ),
      // body: Center(
      //   child: Column(
      //     mainAxisAlignment: MainAxisAlignment.center,
      //     children: [
      //       TextButton(
      //         style: TextButton.styleFrom(
      //           minimumSize: Size(20, 20), // 最小サイズを設定
      //           padding: EdgeInsets.zero, // パディングをゼロ
      //           // fixedSize: const Size(20, 20),
      //           fixedSize: const Size(5, 5),
      //           foregroundColor: Colors.white,
      //           backgroundColor: Colors.blue,
      //         ),
      //         onPressed: () => print('a'),
      //         child: Text('a'),
      //       ),
      //       Text('WordStudy'),
      //       SingleChildScrollView(child: main),
      //     ],
      //   ),
      // ),
    );
  }
}
