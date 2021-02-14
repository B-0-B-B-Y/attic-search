import axios from 'axios';

const { API_HOST } = process.env;

export type SearchResult = {
  items: Array<DataObject> | null;
};

export type DataObject = {
  side: string;
  position: number;
  item: string;
  container: string;
  description: string;
  frequent: boolean;
  keywords: Array<string>;
};

export const search = async (keyword: string): Promise<SearchResult | null> => {
  try {
    const request = await axios.get(
      `${API_HOST}/search/word/${keyword}`
    );
    const data = await request.data;

    return data;
  } catch (err) {
    console.error(err);
  }

  return null;
};
