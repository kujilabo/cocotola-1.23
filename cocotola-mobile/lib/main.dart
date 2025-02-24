import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
// import 'package:mobile/widgets/editor.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/firebase_options.dart';
import 'package:mobile/sql/sql.dart';
import 'package:mobile/util/logger.dart';
import 'package:mobile/widget/word_study/menu_word_study.dart';

Future<void> main() async {
  logger.i('aaa');
  WidgetsFlutterBinding.ensureInitialized();
  logger.i('bbb');
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  logger.i('ccc');
  // await FirebaseAuth.instance.useAuthEmulator('localhost', 9099);
  // logger.i('ddd');

  final db = DB();
  final x = await db.db;

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
