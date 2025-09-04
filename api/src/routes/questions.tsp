import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";
import "../models/common.tsp";
import "../models/question.tsp";
import "../models/user.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

namespace ShitsumonBako;

/**
 * 質問・回答関連のエンドポイント
 */
@tag("Questions")
@route("/api/questions")
interface Questions {
  /**
   * 質問送信
   */
  @post
  @operationId("createQuestion")
  @summary("新しい質問を送信")
  create(@body request: QuestionCreateRequest): {
    @statusCode _: 201;
    @body response: ApiResponse<Question>;
  } | {
    @statusCode _: 400;
    @body error: ApiResponse<ErrorDetails>;
  } | {
    @statusCode _: 403;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 受信した質問一覧取得
   */
  @get
  @route("/received")
  @operationId("getReceivedQuestions")
  @summary("受信した質問一覧取得")
  getReceived(
    @header authorization: string,
    ...PaginationQuery,
    ...QuestionFilter,
  ): {
    @statusCode _: 200;
    @body response: PaginatedResponse<QuestionDetail>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 送信した質問一覧取得
   */
  @get
  @route("/sent")
  @operationId("getSentQuestions")
  @summary("送信した質問一覧取得")
  getSent(
    @header authorization: string,
    ...PaginationQuery,
    ...SentQuestionFilter,
  ): {
    @statusCode _: 200;
    @body response: PaginatedResponse<QuestionDetail>;
  } | {
    @statusCode _: 401;
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 質問詳細取得
   */
  @get
  @route("/{questionId}")
  @operationId("getQuestionDetail")
  @summary("質問詳細情報取得")
  getQuestion(@header authorization: string, @path questionId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<QuestionDetail>;
      }
    | {
        @statusCode _: 401;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 404;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 403;
        @body error: ApiResponse<ErrorDetails>;
      };

  /**
   * 質問編集
   */
  @patch
  @route("/{questionId}")
  @operationId("updateQuestion")
  @summary("質問内容を編集")
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
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 404;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 403;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 400;
        @body error: ApiResponse<ErrorDetails>;
      };

  /**
   * 質問削除
   */
  @delete
  @route("/{questionId}")
  @operationId("deleteQuestion")
  @summary("質問を削除")
  deleteQuestion(@header authorization: string, @path questionId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<{}>;
      }
    | {
        @statusCode _: 401;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 404;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 403;
        @body error: ApiResponse<ErrorDetails>;
      };

  /**
   * 質問を既読にする
   */
  @patch
  @route("/{questionId}/read")
  @operationId("markQuestionAsRead")
  @summary("質問を既読にマーク")
  markAsRead(@header authorization: string, @path questionId: string):
    | {
        @statusCode _: 200;
        @body response: ApiResponse<Question>;
      }
    | {
        @statusCode _: 401;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 404;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 403;
        @body error: ApiResponse<ErrorDetails>;
      };

  /**
   * 質問に回答
   */
  @post
  @route("/{questionId}/answer")
  @operationId("createAnswer")
  @summary("質問に回答を作成")
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
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 404;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 403;
        @body error: ApiResponse<ErrorDetails>;
      }
    | {
        @statusCode _: 400;
        @body error: ApiResponse<ErrorDetails>;
      };

  /**
   * ユーザーの公開Q&A一覧取得
   */
  @get
  @route("/public/{userId}")
  @operationId("getPublicQAs")
  @summary("ユーザーの公開質問・回答一覧取得")
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
    @body error: ApiResponse<ErrorDetails>;
  };

  /**
   * 通報機能
   */
  @post
  @route("/report")
  @operationId("reportContent")
  @summary("不適切なコンテンツを通報")
  report(@body request: ReportRequest): {
    @statusCode _: 201;
    @body response: ApiResponse<{}>;
  } | {
    @statusCode _: 400;
    @body error: ApiResponse<ErrorDetails>;
  };
}
