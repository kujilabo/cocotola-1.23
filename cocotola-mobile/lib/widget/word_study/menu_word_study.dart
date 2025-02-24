import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/auth_repository.dart';
import 'package:mobile/provider/problem_provider.dart';
import 'package:mobile/provider/text_field_value_list_provider.dart';
import 'package:mobile/provider/word_study_status.dart';
import 'package:mobile/util/logger.dart';
import 'package:mobile/widget/word_study/word_study.dart';

class MenuWordStudy extends ConsumerWidget {
  const MenuWordStudy({super.key});

  double _calcWidth(String text, TextStyle style) {
    final textPainter = TextPainter(
      text: TextSpan(text: text, style: style),
      textDirection: TextDirection.ltr,
    )..layout();
    // textPainter.layout();
    logger.i('textPainter.size: ${textPainter.size}');
    return textPainter.size.width;
  }

  Widget _buildUserName(AsyncValue<User?> user) {
    switch (user) {
      case AsyncData(:final value):
        if (value == null) {
          return const Text(
            'No User',
            style: TextStyle(fontSize: 24),
          );
        }
        logger.i('user: $value');
        return Text(
          value.isAnonymous ? 'Anonymous' : value.displayName ?? 'No Name',
          style: const TextStyle(fontSize: 24),
        );
      case AsyncLoading():
        return const Text(
          'Loading...',
          style: TextStyle(fontSize: 24),
        );
      case AsyncError(:final error):
        logger.i('user: error $error');
        return const Text(
          'Error...',
          style: TextStyle(fontSize: 24),
        );
      default:
        logger.i('user: default');
        return const Text(
          'Error...',
          style: TextStyle(fontSize: 24),
        );
    }
  }

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final wordStudyStatusNotifier = ref.watch(wordStudyStatusProvider.notifier);
    final authRepositoryNotifier = ref.watch(authRepositoryProvider.notifier);
    final authRepository = ref.watch(authRepositoryProvider);

    final user = authRepository.whenData((value) => value.user);

    switch (user) {
      case AsyncData(:final value):
        logger.i('user: $value');
      case AsyncLoading():
        logger.i('user: loading');
      case AsyncError(:final error):
        logger.i('user: error $error');
      default:
        logger.i('user: default');
    }

    final width = _calcWidth('aaaaa', const TextStyle(fontSize: 24));
    logger.i('width: $width');
    // final textFieldValueListProvider = ref.watch(textFieldValueListProvider);
    return Scaffold(
      appBar: AppBar(
        title: const Text('Word Study'),
      ),
      body: SafeArea(
        child: Column(
          children: [
            _buildUserName(user),
            TextField(
              controller: TextEditingController(),
            ),
            const Center(
              child: Text(
                'Word StudyTOP',
                style: TextStyle(fontSize: 24),
              ),
            ),
            ElevatedButton(
              onPressed: authRepositoryNotifier.signOut,
              child: const Text('Sign Out'),
            ),
            ElevatedButton(
              onPressed: authRepositoryNotifier.signInAnonymously,
              child: const Text('Sign In Anonymously'),
            ),
            ElevatedButton(
              onPressed: authRepositoryNotifier.singInWithGoogle,
              child: const Text('Sign In With Google'),
            ),
            ElevatedButton(
              onPressed: () {
                wordStudyStatusNotifier.setQuestionStatus();
                ref
                  ..invalidate(textFieldValueListProvider)
                  ..invalidate(problemProvider);

                Navigator.of(context).push(
                  MaterialPageRoute<dynamic>(
                    builder: (context) => const WordStudy(),
                  ),
                );
              },
              child: const Text('Save Expense'),
            ),
          ],
        ),
      ),
    );
  }
}
