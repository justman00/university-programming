import { useState, useEffect, useMemo } from 'react';
import { Review } from '../types';
import { getReviews } from '../services/reviews';
import { Body, Table, Title } from '@sumup/circuit-ui';
import Layout from '@/components/Layout';
import { css } from '@emotion/react';
import styled from '@emotion/styled';
import { Chart } from 'react-google-charts';
import { useRouter } from 'next/router';

export const data = [
  ['Task', 'Hours per Day'],
  ['Work', 11],
  ['Eat', 2],
  ['Commute', 2],
  ['Watch TV', 2],
  ['Sleep', 7],
];

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

const Analysis = () => {
  const [reviews, setReviews] = useState<Review[]>([]);
  const [selectedTopicIndex, setSelectedTopicIndex] = useState<number>(-1);
  const router = useRouter();

  const mappedReviews = useMemo(() => [['Topic', 'Count per topic classification'], ...mapReviews(reviews)], [reviews]);
  console.log(mappedReviews);

  useEffect(() => {
    if (selectedTopicIndex !== -1) {
      console.log('selectedTopicIndex', selectedTopicIndex);
      const topicClassification = mappedReviews[selectedTopicIndex + 1][0];

      router.push(`/${topicClassification}/sentiment_analysis`);
    }
  }, [mappedReviews, selectedTopicIndex, router]);

  useEffect(() => {
    const fetchReviews = async () => {
      const reviews = await getReviews('trustpilot');
      setReviews(reviews);
    };
    fetchReviews();
  }, []);

  const callback = ({ chartWrapper, google }: { chartWrapper: any; google: any }) => {
    const chart = chartWrapper.getChart();
    google.visualization.events.addListener(chart, 'click', (e: any) => {
      console.log('click click');
      let id = e.targetID?.split('#')[1];

      console.log('id', id);
      if (!id) {
        return;
      }

      setSelectedTopicIndex(Number(id));
    });
  };

  return (
    <Layout>
      <Container>
        <StyledTitle as="h1" size="three">
          Analysis
        </StyledTitle>
        <StyledBody>Analyze the reviews fetched from different sources and analyzed by our AI.</StyledBody>
        <Chart
          chartType="PieChart"
          data={mappedReviews}
          width={'100%'}
          height={'700px'}
          chartEvents={[
            {
              eventName: 'ready',
              callback,
            },
          ]}
        />
      </Container>
    </Layout>
  );
};

const mapReviews = (reviews: Review[]) => {
  // count the number of reviews based on topic classification
  const topicClassification = reviews.reduce((acc, review) => {
    const topics = review.topic_classification;
    topics.forEach((topic) => {
      if (!acc[topic]) {
        acc[topic] = 0;
      }
      acc[topic] += 1;
    });

    return acc;
  }, {} as Record<string, number>);

  return Object.entries(topicClassification).map(([topic, count]) => [topic, count]);
};

export default Analysis;
