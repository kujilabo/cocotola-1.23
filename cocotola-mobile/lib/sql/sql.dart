// import 'package:mobile/widgets/editor.dart';
import 'package:mobile/util/logger.dart';
import 'package:path/path.dart';
import 'package:sqflite/sqflite.dart';

const int dbVersion = 1;

Map<String, List<String>> scripts = <String, List<String>>{
  '1': [
    '''
create table `word_sentence_pair` (
 `id` integer primary key
,`document_id` varchar(36)
,`workbook_id` int not null
,`src_sentence_number` int not null
,`src_lang2` varchar(2) not null
,`src_text` varchar(1000) not null
,`src_author` varchar(20) not null
,`dst_sentence_number` int not null
,`dst_lang2` varchar(2) not null
,`dst_text` varchar(1000) not null
,`dst_author` varchar(20) not null
,`created_at` datetime not null
,`updated_at` datetime not null
);
''',
<<<<<<< HEAD
    '''
CREATE INDEX document_id_word_sentence_pair_idx ON word_sentence_pair(document_id);
''',
    '''
CREATE INDEX workbook_id_word_sentence_pair_idx ON word_sentence_pair(workbook_id);
    ''',
    '''
CREATE INDEX src_lang2_word_sentence_pair_idx ON word_sentence_pair(src_lang2);
''',
    '''
CREATE INDEX dst_lang2_word_sentence_pair_idx ON word_sentence_pair(dst_lang2);
''',
  ],
  '2': [
    '''
CREATE INDEX document_id_word_sentence_pair_idx ON word_sentence_pair(document_id);
''',
=======
    'CREATE INDEX document_id_word_sentence_pair_idx ON word_sentence_pair(document_id);',
    'CREATE INDEX workbook_id_word_sentence_pair_idx ON word_sentence_pair(workbook_id);',
    'CREATE INDEX src_lang2_word_sentence_pair_idx ON word_sentence_pair(src_lang2);',
    'CREATE INDEX dst_lang2_word_sentence_pair_idx ON word_sentence_pair(dst_lang2);',
  ],
  '2': [
    'CREATE INDEX document_id_word_sentence_pair_idx ON word_sentence_pair(document_id);',
>>>>>>> c876418 (sqlite)
  ],
};

class DB {
  factory DB() => _instance;
  DB._();
  static final DB _instance = DB._();

  Future<Database> get db async => _db ??= await initDB();
  Database? _db;

  // Future<Database?> get db async {
  //   if (_db != null) {
  //     return _db;
  //   }
  //   _db = await initDB();
  //   return _db;
  // }

  // Future<void> init() async {
  //   _db = await initDB();
  // }

  Future<Database> initDB() async {
    const filePath = 'cocotola.db';
    final dbPath = await getDatabasesPath();
    final path = join(dbPath, filePath);
    // await deleteDatabase(path);
    logger.i(path);

    return openDatabase(
      path,
      version: dbVersion,
      onCreate: (db, version) async {
        for (var i = 1; i <= version; i++) {
          final queries = scripts[i.toString()];
          if (queries != null) {
            for (final query in queries) {
              await db.execute(query);
            }
          }
        }
      },
      onUpgrade: (db, oldVersion, newVersion) async {
        for (var i = oldVersion + 1; i <= newVersion; i++) {
          final queries = scripts[i.toString()];
          if (queries != null) {
            for (final query in queries) {
              await db.execute(query);
            }
          }
        }
      },
    );
  }
}
