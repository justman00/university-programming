package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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
	EntityID  string
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

type ChatgptResponse struct {
	Sentiment           string                   `json:"sentiment"`
	Emotion             string                   `json:"emotion"`
	TopicClassification TopicClassificationArray `json:"topic_classification"`
	Justification       string                   `json:"justification"`
	Translation         string                   `json:"translation"`
}

type TopicClassificationArray []string

func (t *TopicClassificationArray) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("failed to unmarshal topic classification: %w", err)
	}

	s = strings.ReplaceAll(s, " ", "")
	topicClassificationArray := strings.Split(s, ",")

	*t = topicClassificationArray

	return nil
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
		SystemChatMessage: &systemPrompt,
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

	// yolo
	marshaledResponse, err := json.Marshal(chatGPTResponse)
	if err != nil {
		return fmt.Errorf("marshal chatgpt response to obbey db: %w", err)
	}

	if err := processor.dbInstance.InsertReview(ctx, &db.InsertReviewParams{
		ID:              uuid.New(),
		Rating:          p.Rating,
		Source:          p.Source,
		Review:          p.Contents,
		Analysis:        string(marshaledResponse),
		OriginalPayload: string(t.Payload()),
		EntityID:        p.EntityID,
		ReviewCreatedAt: p.CreatedAt,
		ReviewUpdatedAt: p.CreatedAt,
	}); err != nil {
		return fmt.Errorf("insert review: %w", err)
	}

	return nil
}

var systemPrompt = `
You are an AI designed to analyze and categorize customer feedback based on sentiment, emotion, and topic. Below is a JSON structure that represents an example of how customer feedback is classified:
[
    {
        "example_feedback": "We bought 2 SumUp Solo machines for a village charity fund raising event just over a year ago. They were were only used on that day. We put them away, boxed, in working order and, this week, took them out to use for our village show. One is working, the other is blocked. It can't be unblocked and is out of warranty by 2 months according to SumUp customer services.They asked for a receipt yet all the information is on the reader! It is unbelievable that in this day and age and our understanding of technology that we must now throw a £60+ gadget away solely because it can't be unblocked. Will never buy from this company again.",
        "topic_classification": "solo_card_reader"
    },
    {
        "example_feedback": "After using the Air Sum Up machine quite successfully for one Pop Up Sale whiich I had. .it suddenly stopped working. When I checked with the Company, they wanted me to buy another machine!! I wouldn't trust them.!! I think that is probably what they do after a certain amount of sales are reached to boost their income.although they would miss out on the commission. I am going to try the Bank of Scotland, I hope more reliable. Thank you for reading, A Country Affair Scone Perth",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "Didn't understand what I was asking. I went onto chat twice and problem still occurs.\nWhenever I want to take a payment my Sumup Air automatically defaults to 20% VAT and I have to remember to change it to 0% VAT & Press Save. The next time it happens again it's on 20% VAT & have to repeat the same each time. It doesn't stay at 0% or even NO VAT when I press Save.",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "I bought a Pos on October 30th and also made an account. Billing email never came, also no POS and they have never activated my account. Since then I have sent 3 emails and no one has responded. I have requested a refund and also tried to contact via messenger without any response. FINALLY after a week they have responded to arrange my account and also have found the purchace.....",
        "topic_classification": "pos"
    },
    {
        "example_feedback": "Overall I've been very happy with the Sumup card reader and transaction service; occasional hiccups with the app but generally ok. However I was recently invited to use a new feature of QR codes and was sent all the necessary documentation etc. Unfortunately I cannot get it to work for my customers and have tried to reach out on three occasions over the last 2 weeks for support and have received zero response from customer support not even a professional etiquette acknowledgement.",
        "topic_classification": "qr_code"
    },
    {
        "example_feedback": "Useless support i ever seen. From the 1st day, i'm contacting again and again to their live chat about activation of payment scope for WooCommerce payments. But No and not at all, their internal team never contact you back and don't resolve the issue, don't know what are they doing. I don't recommend SumUp for any type of business. They will never resolve any issues instead of messages.",
        "topic_classification": "direct_integration"
    },
    {
        "example_feedback": "Not able to get a response to my issue\nI want to change my account name and bank details but not getting any response. I've asked them to let me know that if it's not possible that I want to then close my account and stop getting people to pay me via Sumup!! I've money sitting there but I can't access it.\n\nAlso I called their customer service who told me the reader was purchased last year which is ridiculous because I bought it a week ago. Are they selling refurbished readers??\n\nI'm feeling uneasy with Sumup and stuck because I can't even close my account!\n\nDate of experience: 12 August 2023",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "Multiple issues with the online store. Failed transactions and repeatedly told its the bank's fault, even though it's multiple banks and vendors. Customer service has been appalling and dismissive. I pointed out we are losing customers, and was told I would be contacted. 4 days later, multiple failed transactions, multiple customers lost... no contact. Have also had a payment in sales report not appear in the online store orders. So I can't contact this person, can't let them know I've received their money, and have no clue as to what product they have purchased. Don't offer something that doesn't work! Any prospective businesses, steer clear.\n",
        "topic_classification": "online_store"
    },
    {
        "example_feedback": "The Club I do the accounts for bought Sumup POS system to use in their bar and bitterly regretted it. Difficult to set up but most of all the Reports are very limited and the format is next to useless. It just exports info to Excel and leaves you do all the hard work to get any useful management info, made more difficult by the fact that it transposes the month and day of the date, which you then have to go through and correct before being able to sort by date.\n\nUtterly, utterly useless. Don't waste your money, buy Square instead. I use Square for my own business and it leaves Sumup standing, produces simple accessible reports and even has a booking system as well. Sumup you need to learn from your competitors....",
        "topic_classification": "pos"
    },
    {
        "example_feedback": "I have been a business customer with SumUp since 2018. Through word of mouth, and due to my love of their products (terminals, POS, invoices, links) I signed up two of my businesses to SumUp and referred many connections who bought terminals too.\n\nUnfortunately, now that the company has grown so rapidly, and expanded its product range, I have noticed a steep decline in the quality of customer service. I have not been able to resolve any issues in the past 12 months via live chat or phone call. It's an Indian or Chinese call centre who are robotic in their replies. Very little help.\n\nThe life of the terminals is not what it used to be either - especially the Solo - which recently died on me.\n\nI am now transitioning my businesses to Square.\n\nDate of experience: 22 July 2023",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "My invoices haven't been working now for over 2 weeks. This is impacting my business and I am unable to send invoices or edit existing ones. Get this sorted or you will lose custom! I'm looking at changing to another provider.\n\nDate of experience: 23 July 2023",
        "topic_classification": "invoices"
    },
    {
        "example_feedback": "Like others, I have had an awful experience with SumUp support. The live chat agents have been far from helpful when trying to get an application enabled for the payments scope but instead I have been directed to email the integrations team which I have done multiple times with no response. I have clients that use SupUp for their POS solution some of who now will be switching to our other payment integrations as SumUp have been terrible to deal with.",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "It saves a lot of hassle when getting payment. However, something is always broken and when it is a feature that you are paying for such as invoices, it's extremely frustrating. The machine also has connection issues at least once a week, which means you have to go through the five minutes of reset procedure in order to take a payment, and if you do the five step procedure in the wrong order you have to start again.\n\nDate of experience: 15 June 2023",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "These guys promise a lot but do nothing to help small business owners.\n\nFor starters you can forget about next day settlement payments as they literally free style when and what time you should receive your money.\n\nIt's been 4 days and my card company have cleared 4 payments to my sum up business account which only I've received one payment only!! I've called sum up support team multiple times there's no actual urgency to find out where the money is and how long it will take to receive my money. I've emailed them proofs that my card company have sent them my money still I've heard nothing.\n\nI don't understand how a company like this doesn't get properly regulated by the financial conduct authorities and I will be passing my case to them hopefully somethings gets done because I'm sure I'm not the only one.",
        "topic_classification": "business_account"
    },
    {
        "example_feedback": "They tell you it's next day if you transfer to your business account then when transfer to your bank account it's pending for up to 3 days ... Don't be lied too because it's cheap",
        "topic_classification": "business_account"
    },
    {
        "example_feedback": "APPALLING CUSTOMER SERVICE. I have been trying to withdraw money from my business account for over a week now. Every day I call and get fobbed off and sent automations on email. My account is blocked and they cannot tell me why, I am unable to transfer any of my funds that I need. Just a horrendous service.",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "Do not use SumUp and specifically Do NOT use “SumUp Pay”. Claims everywhere you can withdraw at any time but that's a lie! Not even an option and now many is held against my wishes with no choice but to spend it before my account can be closed. Terrible follow up service waiting days for a reply. My bank also treated it as a cash transaction so charged me up front and interest! No offer of help at all.",
        "topic_classification": "sumup_pay"
    },
    {
        "example_feedback": "This is an amazing set up for small business. Card machine is wireless, easy to use, compact and looks good.\nPayout frequently can be set, I prefer daily. Customer not present option is a help. Love that I can just txt an invoice with payment links. Only thing I wish they would resume is the 'cash advance' Brilliant, quick, cheap way of fast lending without all the hassle. You don't even notice the repayments. This machine is a definite asset to a small business. 100% recommended.",
        "topic_classification": "lending"
    },
    {
        "example_feedback": "It all looked very good on the packaging, promises of quick payouts etc. I raised an invoice via the SumUp app to which my customer promptly paid, when I looked on the application to find out if the money had landed in my bank account I had a message informing me that my account was restricted pending verification. I questioned this and was asked to send copy of the invoice to SumUp to validate, this by the way way the invoice raised on the SumUp app! I sent the invoice anyway and are still waiting for someone to validate this invoice, why on earth god only knows ! This delay in payment is unacceptable when they boast of quick payments (3 days now!) I'm having to ask my customer to cancel the transaction and find another payment method. I do notice that more than 25% of the reviews about SumUp are less than happy !",
        "topic_classification": "invoices"
    },
    {
        "example_feedback": "I've found Sumup to be fairly average for providing my small business card reader.\nI'm not sure if it's down to our connection or not, but about 5-10% of transactions on a daily basis fail to go through and require a re-try on the card reader, not sure why.\nThe fees at 1.69% are higher than most of their competitors. I find the reports on the dashboard are quite lacking in filters, for example there is no way to see the separate tip amounts from what I can see.\nInitial merchant verification was also a tad confusing.",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "Very disappointed. The new Sumup Solo (with printer) is absolutely rubbish and not fit for purpose. It's totally impractical in a busy retail environment. I reported an issue with mine as it would not automatically print receipts and I no longer had a printer icon on my screen to change the settings. I was told it was being referred to their technical department but that was over 10 days ago and I've heard nothing. I'm also concerned that I've heard nothing about fulfilling my PCI compliance.",
        "topic_classification": "solo_card_reader"
    },
    {
        "example_feedback": "taking payments via SumUp is ok, its works and the fee isnt too high. But the accounting side of the software is a nightmare. I have funds that come in from different areas for different services but the accounting reports doesnt show any reference number that relates to my sales, no customer name or customer email for me to balance my books. Spoke to SumUp who told me they couldn't give me this information due to GDPR which is ridiculous. Support team were very unhelpful with any kind of solution.",
        "topic_classification": "accounting"
    },
    {
        "example_feedback": "Not even managed to use it yet. My account has been blocked for verification for 5 days already. Can't even logon to use to chat to get help. Phoned customer service and all they did was say yes your account is being verified and yes we've received your id, you just have to wait. I totally get that things need to be checked but this is ridiculous Really not happy Gone to square and everything was sorted in a couple of hours.",
        "topic_classification": "verification"
    },
    {
        "example_feedback": "sumup have been absolutely useless at trying to help me recover money they have had in their account since September 2022. Each time I try to log in it tells me the account is being verified but no-one at sum up will tell me which phone this verification is going to as it does not come to me as the registered phone number!!! I have sent them an invoice, sent them original print out (that they sent to me!) with my name and address on it but will not answer me as my email address and phone is not the one they seem to have. I don't know what to do now except get legal people involved as sumup are withholding our money!",
        "topic_classification": "verification"
    },
    {
        "example_feedback": "I purchased machine and submitted my documents for verification. Although it is mentioned on the site that it takes 72 hours to review the application, mine has been under review more than 2 weeks by now and there is no definite answer when I will receive a reply. Very disappointed with the customer service. Looking for alternative POS.",
        "topic_classification": "verification"
    },
    {
        "example_feedback": "There were no problems at all in the beginning. My clients were able to pay with an EC card without any problems. But after just a short time, the receipt printer stopped working. The battery is empty within half a day!",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "The device is generally very easy to use. Unfortunately, there has been a contact error for a few weeks now, the device does not recognize the cash register roll/the receipts cannot be printed out. A replacement device would be useful. Unfortunately, I can't request a replacement device via “sumup”.",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "The device looks chic and is easy to use. It is incredibly well made. Unfortunately, the integrated connection fails regularly despite the super network - the reboot takes forever, customers have to wait. This is really annoying!",
        "topic_classification": "solo_card_reader"
    },
    {
        "example_feedback": "I have been a SumUp user for several years. This week we have had a problem. Someone infiltrated the SumUp security and changed the password on the account. It appears to be impossible to regain control. The infiltrator has not changed the bank account for receipts (and I guess cannot) but I cannot issue invoices or correlate payments received to outstanding invoices. I spent ages on phone yesterday trying to resolve - this ended with the agent saying she could not and would have to 'escalate' the issue. I was promised a call back which has not materialised. After a long time as a satisfied user of SumUp, I am now entertaining doubts as to whether I should continue to work with it. This review is, ultimately, an attempt to get SumUp to resolve the problem…",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "I've used sun up for years (about 6-7) and always been happy, no problems. So I decided to use it for my other business too, so I set up an account with the email for my other business & couldn't get past the setting up bank details, it just sent me round in circles. I emailed customer service and every time they told me to try something, it made no difference. Eventually they said to create a new account with a new email address (I had to create another email address just for this which was a pain) & it still did the same. It just won't accept my bank details. I emailed again to complain and this time they've just ignored my email completely.",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "Disastrous company. After ordering everything went fine. Device received within a few days. Account created and approved by Sumup. Trial payment made. After 7 days I received an email that I have to return it because my business cannot be verified in the account??? In my complaint email I was simply informed of the terms and conditions. Now I sent it back 8 days ago and have not received my money or a confirmation of receipt of the return. As a result - stay away from Sumup",
        "topic_classification": "verification"
    },
    {
        "example_feedback": "Changing to one star. Horrible connection could not be worse. Scanning for the unit, find it, but won't connect. So, I need to reset Bluetooth connection and do it again. It happens too often, to give this app decent feedback. My phone android software is updated. The SumUp is updated. But connection is awful. A couple of times, I couldn't accept the payments because the Bluetooth connection didn't work. The app on the phone couldn't connect to AIR SumUp. I still believe one star is too much.",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "Every single time I try to connect my Samsung s21 ultra to the reader, it keeps saying can't connect. I keep having to disconnect it completely from my Bluetooth. When my phone goes on standby, it disconnects again! I have an event in a few days, and I can't even trust it to work right!",
        "topic_classification": "card_reader"
    },
    {
        "example_feedback": "Awful customer support.  The system has broken down and there is just no support.  Why is this a popular POS in Berlin?",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "App not worked for sending invoices since the last update. Tried uninstalling, restarting and all the usual things but still won't add items to invoices. Used for two years and always seems to have some kind of glitch.",
        "topic_classification": "invoices"
    },
    {
        "example_feedback": "Great devices, support is high quality. I always reach someone. The support is incredibly nice. I'm completely satisfied and believe that SumUp will continue to grow a lot in the coming years. The support has meanwhile been significantly increased, for which great praise is due. I'm more than excited",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "Disastrous service. When the system fails, you're on your own. There is no hotline on the internet that you can call. No response to emails either. I've been waiting for an answer for over a week. The prices are way too high for this service. Absolutely unacceptable!",
        "topic_classification": "customer_service"
    },
    {
        "example_feedback": "Since updating, the new vat drop down on the checkout page won't save 0%. I change it every time I make a charge to 0%, hit save, then when I go back for a new charge it is back at 20%. It is driving me mad. How do I make it stay at 0%??? I have turned off vat on the settings, nothing works!",
        "topic_classification": "online_store"
    }
]
Here is your task:
- Translate the feedback into English for better understanding.
- Analyze the sentiment of each customer feedback and categorize it as "positive", "neutral", "mixed", or "negative".
- Identify the emotion conveyed in the feedback, categorizing it as "anger", "anxiety", "confused", "delight", "disappointed", "indifferent", "regrets", "safe", "skeptical", "stressed", "surprise", or "trust".
- Based on the example feedback, classify the given customer feedback into one of the available topic classifications: "solo_card_reader", "card_reader", "pos", "qr_code", "direct_integration", "online_store", "customer_service", "business_account", "sumup_pay", "verification", "accounting", "invoices", "lending".
- If the feedback can be classified into more than one category, list all categories separated by a comma. If none of the given topics fit, classify as "other".
- Provide a brief justification for your classifications.
Output a JSON object in the following format:
{
	"sentiment": "<sentiment_value>",
	"emotion": "<emotion_value>",
	"topic_classification": "<topic_classification_value(s)>",
    "justification": "<justification_string>",
    "translation": "<translation_string>"
}
Answer only with the JSON object, even if the provided feedback is in any other language than English.
`

func generateUserPrompt(r string) string {
	return fmt.Sprintf("Given the following customer review, run the analysis on it based on the instructions provided in the system prompt. The review is provided inside the backticks and should be trated as such: `%s`", r)
}
