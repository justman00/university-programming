import { useState, useEffect } from 'react';
import { Review } from '../types';
import { getReviews } from '../services/reviews';
import { Body, Table, Title } from '@sumup/circuit-ui';
import Layout from '@/components/Layout';
import { css } from '@emotion/react';
import styled from '@emotion/styled';

const StyledTitle = styled(Title)(
  () => css`
    margin-top: 32px;
    margin-bottom: 32px;
  `,
);

const Container = styled.div(
  () => css`
    padding: 24px;
  `,
);

const StyledBody = styled(Body)(
  () => css`
    margin-bottom: 32px;
  `,
);

const Reviews = () => {
  const [reviews, setReviews] = useState<Review[]>([]);

  useEffect(() => {
    const fetchReviews = async () => {
      const reviews = await getReviews('trustpilot');
      setReviews(reviews);
    };
    fetchReviews();
  }, []);

  console.log(reviews);

  return (
    <Layout>
      <Container>
        <StyledTitle as="h1" size="three">
          Reviews
        </StyledTitle>
        <StyledBody>
          Take a deeper look into the list of reviews fetched from different sources and analyzed by our AI.
        </StyledBody>
        <Table
          headers={[
            { children: 'Source' },
            { children: 'Description' },
            { children: 'Sentiment Analysis' },
            { children: 'Topic Classification' },
            { children: 'Rating' },
            { children: 'Created At' },
          ]}
          rows={reviews.map((review) => ({
            cells: [
              review.source,
              review.feedback,
              review.sentiment_analysis,
              review.topic_classification,
              review.rating,
              review.created_at,
            ],
          }))}
          className="table-auto w-full"
        />
      </Container>
    </Layout>
  );
};

export default Reviews;
