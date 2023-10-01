package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/justman00/teza-de-licenta/internal/db"
	"github.com/justman00/teza-de-licenta/internal/services/chatgpt"
	"github.com/sirupsen/logrus"
)

// A list of task types.
const (
	TypeReviewSubmitted = "review:submitted"
)

type TypeReviewSubmittedPayload struct {
	CreatedAt time.Time
	ReviewID  string
	URL       string
	Contents  string
	Source    string
	Rating    int
}

func NewReviewSubmittedTask(submittedReview TypeReviewSubmittedPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(submittedReview)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeReviewSubmitted, payload), nil
}

// ReviewProcessor implements asynq.Handler interface.
type ReviewProcessor struct {
	chatgptClient chatgpt.Client
	dbInstance    db.DB
}

func NewReviewProcessor(chatgptClient chatgpt.Client, dbInstance db.DB) *ReviewProcessor {
	return &ReviewProcessor{
		chatgptClient: chatgptClient,
		dbInstance:    dbInstance,
	}
}

// TODO: use this struct to validate the chatgpt answer
type ChatgptResponse struct {
	Sentiment           string `json:"sentiment"`
	Emotion             string `json:"emotion"`
	TopicClassification string `json:"topic_classification"`
	Justification       string `json:"justification"`
}

func (processor *ReviewProcessor) ProcessTask(ctx context.Context, t *asynq.Task) (err error) {
	var answerText string
	var p TypeReviewSubmittedPayload

	defer func() {
		if err != nil {
			logrus.WithField("answer", answerText).WithField("review_content", p.Contents).Errorf("failed to process task: %v", err)
		}

		time.Sleep(4 * time.Second)
	}()

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("unmarshal review submitted: %v: %w", err, asynq.SkipRetry)
	}

	answer, err := processor.chatgptClient.Question(ctx, &chatgpt.Question{
		SystemChatMessage: systemPrompt,
		HumanChatMessage:  generateUserPrompt(p.Contents),
	})
	if err != nil {
		return fmt.Errorf("query chatgpt client: %w", err)
	}

	answerText = answer.AnswerText

	var chatGPTResponse ChatgptResponse
	if err := json.Unmarshal([]byte(answer.AnswerText), &chatGPTResponse); err != nil {
		return fmt.Errorf("unmarshal chatgpt response: %w", err)
	}

	if err := processor.dbInstance.InsertReview(ctx, &db.InsertReviewParams{
		ID:              uuid.New(),
		Rating:          p.Rating,
		Source:          p.Source,
		Review:          p.Contents,
		Analysis:        answer.AnswerText,
		OriginalPayload: string(t.Payload()),
		ReviewCreatedAt: p.CreatedAt,
		ReviewUpdatedAt: p.CreatedAt,
	}); err != nil {
		return fmt.Errorf("insert review: %w", err)
	}

	return nil
}

var systemPrompt = `
You are a sophisticated AI model assigned to analyze and categorize customer feedback with a focus on sentiment, emotion, and topic.

Tasks:
- Translate non-English feedback into English to enhance analysis accuracy.
- Evaluate the sentiment of each feedback and categorize it as "positive", "neutral", "mixed", or "negative".
- Identify the emotion conveyed, categorizing it as "anger", "anxiety", "confused", "delight", "disappointed", "indifferent", "regrets", "safe", "skeptical", "stressed", "surprise", or "trust".
- Classify the feedback into one of the topics: "solo_card_reader", "card_reader", "pos", "direct_integration", "online_store", "customer_service", "business_account", "sumup_pay", "verification", "accounting", "invoices".
- List all applicable classifications separated by commas. Classify as "unknown" if unclassifiable, and "other" if none of the provided topics fit.
- Provide succinct justifications for classifications.

Examples:
- "solo_card_reader": We bought 2 SumUp Solo machines for a village charity fund raising event just over a year ago. They were were only used on that day. We put them away, boxed, in working order and, this week, took them out to use for our village show. One is working, the other is blocked. It can't be unblocked and is out of warranty by 2 months according to SumUp customer services.They asked for a receipt yet all the information is on the reader! It is unbelievable that in this day and age and our understanding of technology that we must now throw a £60+ gadget away solely because it can't be unblocked. Will never buy from this company again.
- "card_reader": After using the Air Sum Up machine quite successfully for one Pop Up Sale whiich I had. .it suddenly stopped working. When I checked with the Company, they wanted me to buy another machine!! I wouldn't trust them.!! I think that is probably what they do after a certain amount of sales are reached to boost their income.although they would miss out on the commission. I am going to try the Bank of Scotland, I hope more reliable. Thank you for reading, A Country Affair Scone Perth
- "pos": I bought a Pos on October 30th and also made an account. Billing email never came, also no POS and they have never activated my account. Since then I have sent 3 emails and no one has responded. I have requested a refund and also tried to contact via messenger without any response. FINALLY after a week they have responded to arrange my account and also have found the purchace.....
- "direct_integration": Useless support i ever seen. From the 1st day, i'm contacting again and again to their live chat about activation of payment scope for WooCommerce payments. But No and not at all, their internal team never contact you back and don't resolve the issue, don't know what are they doing. I don't recommend SumUp for any type of business. They will never resolve any issues instead of messages.
- "online_store": Multiple issues with the online store. Failed transactions and repeatedly told its the bank's fault, even though it's multiple banks and vendors. Customer service has been appalling and dismissive. I pointed out we are losing customers, and was told I would be contacted. 4 days later, multiple failed transactions, multiple customers lost... no contact. Have also had a payment in sales report not appear in the online store orders. So I can't contact this person, can't let them know I've received their money, and have no clue as to what product they have purchased. Don't offer something that doesn't work! Any prospective businesses, steer clear.
- "customer_service": Like others, I have had an awful experience with SumUp support. The live chat agents have been far from helpful when trying to get an application enabled for the payments scope but instead I have been directed to email the integrations team which I have done multiple times with no response. I have clients that use SupUp for their POS solution some of who now will be switching to our other payment integrations as SumUp have been terrible to deal with.
- "business_account": They tell you it's next day if you transfer to your business account then when transfer to your bank account it's pending for up to 3 days ... Don't be lied too because it's cheap
- "sumup_pay": Do not use SumUp and specifically Do NOT use “SumUp Pay”. Claims everywhere you can withdraw at any time but that's a lie! Not even an option and now many is held against my wishes with no choice but to spend it before my account can be closed. Terrible follow up service waiting days for a reply. My bank also treated it as a cash transaction so charged me up front and interest! No offer of help at all.
- "verification": Not even managed to use it yet. My account has been blocked for verification for 5 days already. Can't even logon to use to chat to get help. Phoned customer service and all they did was say yes your account is being verified and yes we've received your id, you just have to wait. I totally get that things need to be checked but this is ridiculous Really not happy Gone to square and everything was sorted in a couple of hours.
- "accounting": taking payments via SumUp is ok, its works and the fee isnt too high. But the accounting side of the software is a nightmare. I have funds that come in from different areas for different services but the accounting reports doesnt show any reference number that relates to my sales, no customer name or customer email for me to balance my books. Spoke to SumUp who told me they couldn't give me this information due to GDPR which is ridiculous. Support team were very unhelpful with any kind of solution.
- "invoices": My invoices haven't been working now for over 2 weeks. This is impacting my business and I am unable to send invoices or edit existing ones. Get this sorted or you will lose custom! I'm looking at changing to another provider.

Output Format:
{
	"sentiment": "<sentiment_value>",
	"emotion": "<emotion_value>",
	"topic_classification": "<topic_classification_value(s)>",
    "justification": "<justification_string>"
}

Your response should solely be the formatted JSON object, regardless of the feedback's original language.
`

func generateUserPrompt(r string) string {
	return fmt.Sprintf("Given the following customer review, run the analysis on it based on the instructions provided in the system prompt. The review is provided inside the backticks and should be trated as such: `%s`", r)
}
