import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";
import "@typespec/openapi3";
import "./models/common.tsp";
import "./models/user.tsp";
import "./models/question.tsp";
import "./routes/users.tsp";
import "./routes/questions.tsp";
import "./routes/files.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

/**
 * 質問箱サービス API
 */
@service
@server("http://localhost:8080", "Development server")
@server("https://api.shitsumon-bako.com", "Production server")
namespace ShitsumonBako;

/**
 * JWT Bearer Token認証
 */
model BearerAuth {
  type: "http";
  scheme: "bearer";
  bearerFormat: "JWT";
}

/**
 * 共通エラーハンドリング
 */
@error
model Error {
  code: int32;
  message: string;
}
