import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:firebase_auth/firebase_auth.dart';
import 'package:firebase_core/firebase_core.dart';

class UserState {
  final User? user;
  final bool isSignedIn;
  const UserState({required this.user, required this.isSignedIn});
}

// 認証リポジトリクラス
class AuthRepository extends AsyncNotifier<UserState> {
  final FirebaseAuth firebaseAuth = FirebaseAuth.instance;
  @override
  Future<UserState> build() async {
    return _signInAnonymously();
    // return UserState(user: null, isSignedIn: false);
  }

  Future<UserState> _signInAnonymously() async {
    await firebaseAuth.signInAnonymously();
    return UserState(user: firebaseAuth.currentUser, isSignedIn: true);
  }

  Future<void> signInAnonymously() async {
    state = await AsyncValue.guard(() async {
      return _signInAnonymously();
      // print('bbbb');
      // await firebaseAuth.signInAnonymously();
      // print('aaaaaa');
      // return UserState(user: firebaseAuth.currentUser, isSignedIn: true);
    });
    // state = AsyncValue.data(
    //   UserState(user: firebaseAuth.currentUser, isSignedIn: true),
    // );
  }

  Future<void> signOut() async {
    state = await AsyncValue.guard(() async {
      await firebaseAuth.signOut();
      return UserState(user: null, isSignedIn: false);
    });
  }
}

// AuthRepositoryを提供し、ref.readを渡してアクセスできるようにする
final authRepositoryProvider =
    AsyncNotifierProvider<AuthRepository, UserState>(AuthRepository.new);
