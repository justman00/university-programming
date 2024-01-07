import { useState, useEffect } from 'react';
import { Review } from '../types';
import { getReviews } from '../services/reviews';
import { Body, Table, Title, useModal } from '@sumup/circuit-ui';
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

const formatDate = (date: string) => {
  const d = new Date(date);
  return `${d.getDate()}/${d.getMonth() + 1}/${d.getFullYear()}`;
};

// TODO: add pagination
const Reviews = () => {
  const { setModal } = useModal();
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
          onRowClick={(row) => {
            console.log('row clicked', row);
            const review = reviews[row];

            setModal({
              closeButtonLabel: 'Close',
              children: (
                <div>
                  <Body variant="highlight" size="one">
                    Justification provided by our AI
                  </Body>
                  <Body>{review.justification}</Body>

                  <Body
                    style={{
                      marginTop: 32,
                      display: 'block',
                    }}
                    variant="highlight"
                    size="one"
                  >
                    Translation of the review if applicable
                  </Body>
                  <Body variant="quote">{review.translation}</Body>
                </div>
              ),
              variant: 'contextual',
            });
          }}
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
              review.contents,
              review.sentiment,
              review.topic_classification.join(', '),
              review.rating,
              formatDate(review.created_at),
            ],
          }))}
          className="table-auto w-full"
        />
      </Container>
    </Layout>
  );
};

export default Reviews;
