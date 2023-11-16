export interface commentAttributes {
  articleId: number;
  content: string;
  name: string;
  url?: string;
  email?: string;
  tel?: string;
}

export const CODE_STATUS_OK = 200;
export interface CommonResponser {
  code: number;
  message: string;
  data: any;
}
export interface ApiMethodList {
  ArticlesList(): Promise<CommonResponser>;
  ArticlesCommentsList(articleId: number): Promise<CommonResponser>;
  ArticlesCommentsCreate(comment: commentAttributes): Promise<CommonResponser>;
  ArticlesDetail(id: number): Promise<CommonResponser>;
}

const api: ApiMethodList = {
  async ArticlesList(): Promise<CommonResponser> {
    const requestUrl: string = "http://localhost:8000/api/index/articles";
    return (await fetch(requestUrl)).json();
  },

  async ArticlesCommentsList(articleId: number): Promise<CommonResponser> {
    const requestUrl: string =
      "http://localhost:8000/api/index/comments/" + articleId;
    return (await fetch(requestUrl)).json();
  },

  async ArticlesCommentsCreate(requestData: commentAttributes): Promise<CommonResponser> {
    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(requestData),
    };

    const requestUrl: string = "http://localhost:8000/api/index/comments";
    return (await fetch(requestUrl, requestOptions as RequestInit).catch(err => {
      console.log("错误错误")
      console.log(err);
      return err;
    })).json();
  },

  async ArticlesDetail(id: number): Promise<CommonResponser> {
    const requestUrl: string = "http://localhost:8000/api/index/articles/" + id;
    return await (await fetch(requestUrl)).json();
  },
};

export default api;
