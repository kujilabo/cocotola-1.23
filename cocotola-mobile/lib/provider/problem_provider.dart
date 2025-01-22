import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/gateway/problem_repository.dart';
import 'package:mobile/model/word_problem.dart';

final problemProvider =
    NotifierProvider<ProblemRepository, WordProblem>(ProblemRepository.new);
