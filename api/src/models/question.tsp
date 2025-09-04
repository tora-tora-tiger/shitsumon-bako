import "@typespec/http";
import "@typespec/rest";
import "./common.tsp";
import "./user.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * 質問ステータス
 */
enum QuestionStatus {
  unread: "unread",
  read: "read",
  answered: "answered",
}

/**
 * 質問
 */
model Question {
  id: string;
  status: QuestionStatus;

  // 受信者ID
  recipientId: string;

  // 送信者ID ログインユーザーの場合のみ
  senderId?: string;

  // 質問本文
  content: string;

  attachedImageList?: ImageFile[];

  // 回答
  answer?: Answer;

  isAnonymous: boolean;
  ...Timestamps;
}

// /**
//  * 質問詳細（受信者視点）
//  */
// model QuestionDetail extends Question {
//   recipient: PublicUserInfo;
//   sender?: PublicUserInfo; // ログインユーザーの場合のみ
// }

/**
 * 回答
 */
model Answer {
  id: string;
  questionId: string;
  content: string;
  attachedImageList?: ImageFile[];
  isPublic: boolean;
  ...Timestamps;
}

/**
 * 公開された質問・回答セット
 */
model PublicQA {
  id: string;
  question: {
    content: string;
    attachedImageList?: ImageFile[];
    isAnonymous: boolean;
    createdAt: utcDateTime;
  };
  answer: {
    content: string;
    attachedImageList?: ImageFile[];
    createdAt: utcDateTime;
  };
  answeredBy: PublicUserInfo;
}

/**
 * 質問送信リクエスト
 */
model QuestionCreateRequest {
  content: string;
  attachedImageIdList?: string[];
  recipientId: string;
}

/**
 * 回答作成リクエスト
 */
model AnswerCreateRequest {
  content: string;
  attachedImageId?: string;
  isPublic: boolean;
}

/**
 * 質問編集リクエスト
 */
model QuestionUpdateRequest {
  content: string;
  attachedImageId?: string;
}

/**
 * 質問一覧のフィルター
 */
model QuestionFilter {
  @query status?: QuestionStatus;
  @query sortBy?: "createdAt" | "updatedAt" = "createdAt";
  @query sortOrder?: SortOrder = SortOrder.desc;
}

/**
 * 送信済み質問一覧のフィルター
 */
model SentQuestionFilter {
  @query hasAnswer?: boolean;
  @query sortBy?: "createdAt" | "updatedAt" = "createdAt";
  @query sortOrder?: SortOrder = SortOrder.desc;
}

/**
 * 通報リクエスト
 */
model ReportRequest {
  targetType: "question" | "user";
  targetId: string;
  reason: string;
  details?: string;
}
