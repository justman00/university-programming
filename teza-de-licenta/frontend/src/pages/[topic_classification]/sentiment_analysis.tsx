import { useState, useEffect, useMemo } from 'react';
import { Review } from '../../types';
import { getReviews } from '../../services/reviews';
import { Body, Table, Title } from '@sumup/circuit-ui';
import Layout from '@/components/Layout';
import { css } from '@emotion/react';
import styled from '@emotion/styled';
import { Chart } from 'react-google-charts';
import { useRouter } from 'next/router';

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

const TopicClassification = () => {
  const router = useRouter();
  const [reviews, setReviews] = useState<Review[]>([]);

  const selectedTopic = router.query.topic_classification as string;

  useEffect(() => {
    const fetchReviews = async () => {
      const reviews = await getReviews('trustpilot', selectedTopic);
      setReviews(reviews);
    };
    fetchReviews();
  }, [selectedTopic]);

  const mappedReviews = useMemo(() => [['Topic', 'Count per sentiment analysis'], ...mapReviews(reviews)], [reviews]);
  console.log(mappedReviews);

  return (
    <Layout>
      <Container>
        <StyledTitle as="h1" size="three">
          Topic Classification - {selectedTopic}
        </StyledTitle>
        <StyledBody>Analyze the reviews fetched from different sources and analyzed by our AI.</StyledBody>
        <Chart chartType="PieChart" data={mappedReviews} width={'100%'} height={'700px'} />
        <Table
          headers={[
            { children: 'Source' },
            { children: 'Description' },
            { children: 'Sentiment Analysis' },
            { children: 'Emotion' },
            { children: 'Topic Classification' },
            { children: 'Rating' },
            { children: 'Created At' },
          ]}
          rows={reviews.map((review) => ({
            cells: [
              review.source,
              review.contents,
              review.sentiment,
              review.emotion,
              review.topic_classification.join(', '),
              review.rating,
              review.created_at,
            ],
          }))}
        />
      </Container>
    </Layout>
  );
};

const mapReviews = (reviews: Review[]) => {
  const sentimentAnalysis = reviews.reduce((acc, review) => {
    const sentimentAnalysis = review.sentiment;
    if (!acc[sentimentAnalysis]) {
      acc[sentimentAnalysis] = 0;
    }
    acc[sentimentAnalysis] += 1;
    return acc;
  }, {} as Record<string, number>);

  return Object.entries(sentimentAnalysis).map(([sentimentAnalysis, count]) => [sentimentAnalysis, count]);
};

export default TopicClassification;
