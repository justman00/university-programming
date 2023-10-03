import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const url = new URL('http://localhost:8080/api/reviews');
    // url.searchParams.append('source', source);
    // url.searchParams.append('limit', '1000');
    // if (selectedTopic) {
    //   url.searchParams.append('topic_classification', selectedTopic);
    // }

    const response = await fetch(url.toString());
    const data = await response.json();
    const reviews = Array.isArray(data) ? data : [];

    return res.send(reviews);
  } catch (error) {
    console.error(error);

    return res.send([]);
  }
}
