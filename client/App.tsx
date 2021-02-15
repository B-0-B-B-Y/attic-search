import { StatusBar } from 'expo-status-bar';
import React, { useState, useEffect } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import auth, { FirebaseAuthTypes } from '@react-native-firebase/auth';
import firestore, {
  FirebaseFirestoreTypes,
} from '@react-native-firebase/firestore';
import { Login } from './components/Login';
import { Search } from './components/Search';
import { Result } from './components/Result';
import { SearchResult } from './util/api';

export default function App() {
  const [initialising, setInitialising] = useState<boolean>(true);
  const [userPermissions, setUserPermissions] = useState<
    FirebaseFirestoreTypes.DocumentData | undefined
  >(undefined);
  const [user, setUser] = useState<FirebaseAuthTypes.User | null>(null);
  const [data, setData] = useState<SearchResult | null>(null);

  const getUserPermissions = async () => {
    const userPermissionsPromise = await firestore()
      .collection('users')
      .doc(user?.uid)
      .get();
    const userPermissions = await userPermissionsPromise.data();

    setUserPermissions(userPermissions);
    if (userPermissions?.allowed) {
    }
  };

  const onAuthStateChanged = (user: FirebaseAuthTypes.User | null) => {
    setUser(user);
    if (initialising) setInitialising(false);
  };

  useEffect(() => {
    const subscriber = auth().onAuthStateChanged(onAuthStateChanged);

    return subscriber;
  }, []);

  useEffect(() => {
    if (user) {
      getUserPermissions();
    }
  }, [user]);

  if (!user) {
    return (
      <View style={styles.container}>
        <Login />
      </View>
    );
  }

  if (user && userPermissions?.allowed) {
    return (
      <View style={styles.container}>
        <Search setData={setData} />
        <Result data={data} />
        <StatusBar style='dark' />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <Text style={styles.text}>
        You don't have permission to use this app.
      </Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#E8E8E8',
    alignItems: 'center',
    justifyContent: 'center',
    height: '100%',
    width: '100%',
  },
  text: {
    fontFamily: 'sans-serif',
    fontSize: 20,
    fontWeight: 'bold',
    width: '90%',
    marginBottom: 32,
  },
});
