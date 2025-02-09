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

class ProblemRepository extends AsyncNotifier<ProblemWithStatus> {
  @override
  Future<ProblemWithStatus> build() async {
    final problemSet = await ref.watch(problemSetProvider.future);
    return ProblemWithStatus(
        problemSet: problemSet, currentProblem: problemSet[0], index: 0);
  }

  void next() {
    final currentState = state.value!;
    state = AsyncValue.data(
      ProblemWithStatus(
        problemSet: currentState.problemSet,
        currentProblem: currentState.problemSet[currentState.index + 1],
        index: currentState.index + 1,
      ),
    );
  }
}

final problemProvider =
    AsyncNotifierProvider<ProblemRepository, ProblemWithStatus>(
        ProblemRepository.new);
