import React from 'react';
import { StyleSheet, View, Text } from 'react-native';
import auth from '@react-native-firebase/auth';
import {
  GoogleSignin,
  GoogleSigninButton,
} from '@react-native-community/google-signin';
import firestoreCreds from '../android/app/google-services.json';

export const Login = () => {
  GoogleSignin.configure({
    webClientId: firestoreCreds.client[0].oauth_client[1].client_id,
  });

  const handleGoogleSignIn = async () => {
    const { idToken } = await GoogleSignin.signIn();
    const googleCredential = auth.GoogleAuthProvider.credential(idToken);

    return auth().signInWithCredential(googleCredential);
  };

  return (
    <View style={styles.container}>
      <Text style={styles.text}>Please sign in with Google to continue</Text>
      <GoogleSigninButton onPress={handleGoogleSignIn} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  text: {
    fontFamily: 'sans-serif',
    fontSize: 20,
    fontWeight: 'bold',
    width: '90%',
    marginBottom: 32,
  },
});
