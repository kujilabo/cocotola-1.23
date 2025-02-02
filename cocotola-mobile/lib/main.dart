import 'package:flutter/material.dart';
import 'package:mobile/widget/word_study/menu_word_study.dart';
import 'package:mobile/widget/word_study/word_study.dart';
// import 'package:mobile/widgets/editor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';
import 'firebase_options.dart';

Future<void> main() async {
  print('aaa');
  WidgetsBinding widgetsBinding = WidgetsFlutterBinding.ensureInitialized();
  print('bbb');
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  print('ccc');
  await FirebaseAuth.instance.useAuthEmulator('localhost', 9099);
  print('ddd');
// void main() {
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
