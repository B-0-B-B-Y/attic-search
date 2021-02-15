import { StatusBar } from 'expo-status-bar';
import React, { useState, useEffect } from 'react';
import { StyleSheet, View } from 'react-native';
import auth, { FirebaseAuthTypes } from '@react-native-firebase/auth';
import { Login } from './components/Login';
import { Search } from './components/Search';
import { Result } from './components/Result';
import { SearchResult } from './util/api';

export default function App() {
  const [initialising, setInitialising] = useState<boolean>(true);
  const [user, setUser] = useState<FirebaseAuthTypes.User | null>(null);
  const [data, setData] = useState<SearchResult | null>(null);

  const onAuthStateChanged = (user: FirebaseAuthTypes.User | null) => {
    setUser(user);
    if (initialising) setInitialising(false);
  };

  useEffect(() => {
    const subscriber = auth().onAuthStateChanged(onAuthStateChanged);

    return subscriber;
  }, []);

  if (!user) {
    return (
      <View style={styles.container}>
        <Login />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <Search setData={setData} />
      <Result data={data} />
      <StatusBar style='dark' />
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
});
