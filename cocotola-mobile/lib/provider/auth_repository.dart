import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'package:google_sign_in/google_sign_in.dart';

class UserState {
  const UserState({required this.user, required this.isSignedIn});
  final User? user;
  final bool isSignedIn;
}

// 認証リポジトリクラス
class AuthRepository extends AsyncNotifier<UserState> {
  final FirebaseAuth firebaseAuth = FirebaseAuth.instance;
  GoogleAuthProvider provider = GoogleAuthProvider();
  @override
  Future<UserState> build() async {
    provider.setCustomParameters({
      'prompt': 'select_account',
    });
    return _signInAnonymously();
    // return UserState(user: null, isSignedIn: false);
  }

  Future<UserState> _signInAnonymously() async {
    await firebaseAuth.signInAnonymously();
    return UserState(user: firebaseAuth.currentUser, isSignedIn: true);
  }

  Future<UserState> _signInWithGoogle() async {
    final googleLogin = GoogleSignIn(
      scopes: [
        'email',
      ],
    );
    // await firebaseAuth.signInWithPopup(provider);
    // await firebaseAuth.signInWithRedirect(provider);
    final signinAccount = await googleLogin.signIn();
    if (signinAccount == null) throw ('aaa');
    final auth = await signinAccount.authentication;
    final credential = GoogleAuthProvider.credential(
      idToken: auth.idToken,
      accessToken: auth.accessToken,
    );
    // 認証情報をFirebaseに登録
    final user =
        (await FirebaseAuth.instance.signInWithCredential(credential)).user;
    if (user == null) throw ('bbb');
    return UserState(user: user, isSignedIn: true);
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

  Future<void> singInWithGoogle() async {
    state = await AsyncValue.guard(() async {
      return _signInWithGoogle();
    });
  }

  Future<void> signOut() async {
    state = await AsyncValue.guard(() async {
      await firebaseAuth.signOut();
      return const UserState(user: null, isSignedIn: false);
    });
  }
}

// AuthRepositoryを提供し、ref.readを渡してアクセスできるようにする
final authRepositoryProvider =
    AsyncNotifierProvider<AuthRepository, UserState>(AuthRepository.new);
