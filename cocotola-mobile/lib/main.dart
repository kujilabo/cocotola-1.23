import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
// import 'package:mobile/widgets/editor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/firebase_options.dart';
import 'package:mobile/widget/word_study/menu_word_study.dart';

Future<void> main() async {
  print('aaa');
  WidgetsFlutterBinding.ensureInitialized();
  print('bbb');
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  print('ccc');
  // await FirebaseAuth.instance.useAuthEmulator('localhost', 9099);
  print('ddd');
// void main() {
  runApp(
    const ProviderScope(
      child: MaterialApp(
        home: MenuWordStudy(),
        // home: WordStudyMain(),
        // home: WordStudyMain(),
      ),
    ),
  );
}
