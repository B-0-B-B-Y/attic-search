import React from 'react';
import { ScrollView, Text } from 'react-native';
import { SearchResult } from '../util/api';
import { Item } from './Item';

interface ResultProps {
  data: SearchResult | null;
}

export const Result = ({ data }: ResultProps) => {
  console.log(data);

  return (
    <ScrollView>
      {data?.items &&
        data.items.map((item) => <Item key={item.position} item={item} />)}
      {!data && <Text>No items were found unfortunately :(</Text>}
    </ScrollView>
  );
};
