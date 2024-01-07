export type Review = {
  id: number;
  source: string;
  topic_classification: string[];
  justification: string;
  translation: string;
  emotion: string;
  sentiment: string;
  contents: string;
  rating: number;
  created_at: string;
};
