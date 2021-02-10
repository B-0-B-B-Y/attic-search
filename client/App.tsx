import { StatusBar } from 'expo-status-bar';
import React, { useState } from 'react';
import { StyleSheet, ScrollView, View } from 'react-native';
import { Search } from './components/Search';
import { Result } from './components/Result';
import { SearchResult } from './util/api';

export default function App() {
  const [data, setData] = useState<SearchResult | null>(null);

  return (
    <View style={styles.container}>
      <ScrollView contentContainerStyle={styles.container}>
        <Search setData={setData} />
        <Result data={data} />
      </ScrollView>
      <StatusBar style='auto' />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
