import { Review } from '../types';

export const getReviews = async (source: string, selectedTopic?: string | null): Promise<Review[]> => {
  try {
    const searchParams = new URLSearchParams();
    searchParams.append('source', source);
    searchParams.append('limit', '1000');
    if (selectedTopic) {
      searchParams.append('topic_classification', selectedTopic);
    }

    const response = await fetch(`/api/reviews?${searchParams.toString()}`);
    const data = await response.json();
    console.log(data);
    const reviews = Array.isArray(data) ? data : [];

    return reviews;
  } catch (error) {
    console.error(error);

    return [];
  }
};
