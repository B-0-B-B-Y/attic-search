import React from 'react';
import { View, Text } from 'react-native';
import { capitalise } from '../util';
import { DataObject } from '../util/api';

interface ItemProps {
  item: DataObject | null;
}

export const Item = ({ item }: ItemProps) => {
  return (
    <View>
      <Text>Item: {item?.item}</Text>
      <Text>Position: {item?.position}</Text>
      <Text>Side: {capitalise(`${item?.side}`)}</Text>
      <Text>Description: {item?.description}</Text>
      <Text>Frequently used: {`${item?.frequent}`}</Text>
    </View>
  );
};
