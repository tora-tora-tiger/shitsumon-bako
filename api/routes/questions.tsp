import "@typespec/http";
import "@typespec/rest";
import "../models/common.tsp";
import "../models/question.tsp";
import "../models/user.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

namespace ShitsumonBako;

/**
 * 質問・回答関連のエンドポイント
 */
@route("/api/questions")
interface Questions {
  /**
   * 質問送信
   */
  @post
  create(@body request: QuestionCreateRequest): {
    @statusCode _: 201;
    @body response: ApiResponse<Question>;
  } | {
    @statusCode _: 400;
    @body error: ErrorResponse;
  } | {
    @statusCode _: 403;
    @body error: ErrorResponse;
  };

  /**
   * 受信した質問一覧取得
   */
  @get
  @route("/received")
  getReceived(
    @header authorization: string,
    ...PaginationQuery,
    ...QuestionFilter,
  ): {
    @statusCode _: 200;
    @body response: PaginatedResponse<QuestionDetail>;
  } | {
    @statusCode _: 401;
    @body error: ErrorResponse;
  };

  /**
   * 送信した質問一覧取得
   */
  @get
  @route("/sent")
  getSent(
    @header authorization: string,
    ...PaginationQuery,
    ...SentQuestionFilter,
  ): {
    @statusCode _: 200;
    @body response: PaginatedResponse<QuestionDetail>;
  } | {
    @statusCode _: 401;
    @body error: ErrorResponse;
  };

  /**
   * 質問詳細取得
   */
  @get
  @route("/{questionId}")
  getQuestion(@header authorization: string, @path questionId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<QuestionDetail>;
      }
    | {
        @statusCode _: 401;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 404;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 403;
        @body error: ErrorResponse;
      };

  /**
   * 質問編集
   */
  @patch
  @route("/{questionId}")
  updateQuestion(
    @header authorization: string,
    @path questionId: string,
    @body request: QuestionUpdateRequest,
  ):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<Question>;
      }
    | {
        @statusCode _: 401;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 404;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 403;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 400;
        @body error: ErrorResponse;
      };

  /**
   * 質問削除
   */
  @delete
  @route("/{questionId}")
  deleteQuestion(@header authorization: string, @path questionId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<{}>;
      }
    | {
        @statusCode _: 401;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 404;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 403;
        @body error: ErrorResponse;
      };

  /**
   * 質問を既読にする
   */
  @patch
  @route("/{questionId}/read")
  markAsRead(@header authorization: string, @path questionId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<Question>;
      }
    | {
        @statusCode _: 401;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 404;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 403;
        @body error: ErrorResponse;
      };

  /**
   * 質問に回答
   */
  @post
  @route("/{questionId}/answer")
  createAnswer(
    @header authorization: string,
    @path questionId: string,
    @body request: AnswerCreateRequest,
  ):
    | {
        @statusCode _: 201;
        @body response: ApiResponse<Answer>;
      }
    | {
        @statusCode _: 401;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 404;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 403;
        @body error: ErrorResponse;
      }
    | {
        @statusCode _: 400;
        @body error: ErrorResponse;
      };

  /**
   * ユーザーの公開Q&A一覧取得
   */
  @get
  @route("/public/{userId}")
  getPublicQAs(
    @path userId: string,
    ...PaginationQuery,
    @query sortBy?: "createdAt" | "updatedAt" = "createdAt",
    @query sortOrder?: SortOrder = SortOrder.desc,
  ): {
    @statusCode _: 200;
    @body response: PaginatedResponse<PublicQA>;
  } | {
    @statusCode _: 404;
    @body error: ErrorResponse;
  };

  /**
   * 通報機能
   */
  @post
  @route("/report")
  report(@body request: ReportRequest): {
    @statusCode _: 201;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode _: 400;
    @body error: ErrorResponse;
  };
}
