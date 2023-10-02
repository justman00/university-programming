CREATE INDEX idxgin_reviews_analysis_topic_classification ON reviews USING gin((analysis->'topic_classification') jsonb_path_ops);
CREATE INDEX idx_sentiment ON reviews((analysis->>'sentiment'));
CREATE INDEX idx_emotion ON reviews((analysis->>'emotion'));
