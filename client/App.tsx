import { StatusBar } from 'expo-status-bar';
import React, { useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { Search } from './components/Search';
import { Result } from './components/Result';
import { SearchResult } from './util/api';

export default function App() {
  const [data, setData] = useState<SearchResult | null>(null);

  return (
    <View style={styles.container}>
      <View>
        <Search setData={setData} />
        <Result data={data} />
      </View>
      <StatusBar style='auto' />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#E8E8E8',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
