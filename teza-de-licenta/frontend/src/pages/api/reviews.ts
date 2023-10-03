import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const searchParams = new URLSearchParams(req.query as Record<string, string>);

    const response = await fetch(`http://localhost:8080/api/reviews?${searchParams.toString()}`);
    const data = await response.json();
    const reviews = Array.isArray(data) ? data : [];

    return res.send(reviews);
  } catch (error) {
    console.error(error);

    return res.send([]);
  }
}
