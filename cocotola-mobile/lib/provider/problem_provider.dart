import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:mobile/provider/problem_set_provider.dart';
import 'package:mobile/model/word_problem.dart';

class ProblemWithStatus {
  final List<WordProblem> problemSet;
  final WordProblem currentProblem;
  final int index;
  const ProblemWithStatus(
      {required this.problemSet,
      required this.currentProblem,
      required this.index});
  bool hasNext() {
    return index < problemSet.length - 1;
  }
}

class ProblemRepository extends Notifier<ProblemWithStatus> {
  @override
  ProblemWithStatus build() {
    final problemSet = ref.watch(problemSetProvider);
    return ProblemWithStatus(
        problemSet: problemSet, currentProblem: problemSet[0], index: 0);
  }

  void next() {
    state = ProblemWithStatus(
      problemSet: state.problemSet,
      currentProblem: state.problemSet[state.index + 1],
      index: state.index + 1,
    );
  }
}

final problemProvider = NotifierProvider<ProblemRepository, ProblemWithStatus>(
    ProblemRepository.new);
