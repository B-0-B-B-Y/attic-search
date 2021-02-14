import React, { useState } from 'react';
import { StyleSheet, Text, TextInput, View } from 'react-native';
import { search } from '../util/api';

interface SearchProps {
  setData: Function;
}

export const Search = ({ setData }: SearchProps) => {
  const [keyword, setKeyword] = useState<string>('');

  const getData = async () => {
    const data = await search(keyword);

    setData(data);
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>What are you looking for? 🕵️‍♂️</Text>
      <TextInput
        style={styles.input}
        value={keyword}
        onChangeText={(keyword) => setKeyword(keyword)}
        onSubmitEditing={getData}
        placeholder='Type your keyword in here...'
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    alignItems: 'center',
    justifyContent: 'center',
    marginTop: 72,
  },
  title: {
    fontSize: 24,
    marginTop: 16,
    marginBottom: 16,
    marginRight: 0,
    marginLeft: 0,
    textAlign: 'center',
  },
  input: {
    fontSize: 16,
    padding: 8,
    width: '90%',
  },
});
