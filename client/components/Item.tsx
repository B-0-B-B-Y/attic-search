import React from 'react';
import { StyleSheet, View, Text } from 'react-native';
import { capitalise } from '../util';
import { DataObject } from '../util/api';

interface ItemProps {
  item: DataObject | null;
}

export const Item = ({ item }: ItemProps) => {
  return (
    <View style={styles.container}>
      <Text>Item: {item?.item}</Text>
      <Text>Container: {item?.container}</Text>
      <Text>Position: {item?.position}</Text>
      <Text>Side: {capitalise(`${item?.side}`)}</Text>
      <Text>Description: {item?.description}</Text>
      <Text>Frequently used: {`${item?.frequent}`}</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'flex-start',
    justifyContent: 'center',
    textAlign: 'left',
    backgroundColor: 'white',
    opacity: 0.9,
    padding: 8,
    marginTop: 24,
    width: '90%',
    maxWidth: '90%',
    borderRadius: 8,
  },
});
