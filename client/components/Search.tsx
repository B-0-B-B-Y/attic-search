import React, { useState } from 'react';
import { StyleSheet, Text, ScrollView, TextInput, View } from 'react-native';
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
    <View>
      <Text>What are you looking for?</Text>
      <TextInput
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
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
