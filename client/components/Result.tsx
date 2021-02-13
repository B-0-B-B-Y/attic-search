import React from 'react';
import { StyleSheet, ScrollView, View, Text } from 'react-native';
import { SearchResult } from '../util/api';
import { Item } from './Item';

interface ResultProps {
  data: SearchResult | null;
}

export const Result = ({ data }: ResultProps) => {
  return (
    <ScrollView contentContainerStyle={styles.container}>
      <View style={styles.searchResultCounter}>
        {data && (
          <Text style={styles.counter}>Found: {data?.items?.length || 0}</Text>
        )}
      </View>
      {data?.items &&
        data.items.map((item) => <Item key={item.position} item={item} />)}
      {!data && <Text style={styles.text}>No items were found.</Text>}
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  text: {
    marginTop: 32,
  },
  searchResultCounter: {
    flex: 1,
    alignItems: 'flex-end',
    width: '90%',
    marginTop: 16,
  },
  counter: {
    fontWeight: 'bold',
  },
});
