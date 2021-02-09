import got from 'got'

const { API_HOST } = process.env

type DataObject = {
  side: string,
  position: number,
  item: string,
  container: string,
  description: string,
  frequent: boolean,
  keywords: Array<string>
}

export const search = async (keyword: string): Promise<DataObject | null> => {
  try {
    const request = await got.get(`${API_HOST}/search/word/${keyword}`)
    const data = await JSON.parse(request.body)

    return data
  } catch (err) {
    console.error(err)
  }

  return null
}
