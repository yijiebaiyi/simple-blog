export interface commentAttributes {
  articleId: number;
  Content: string;
  Name: string;
  Url?: string;
  Email?: string;
  Tel?: string;
}
export interface ApiMethodList {
  ArticlesList(): any;
  ArticlesCommentsList(articleId: number): any;
  ArticlesCommentsCreate(comment: commentAttributes): any;
  ArticlesDetail(id: number): any;
}

const api: ApiMethodList = {
  async ArticlesList() {
    const requestUrl: string = "http://localhost:8000/api/index/articles";
    return await (await fetch(requestUrl)).json();
  },

  async ArticlesCommentsList(articleId: number) {
    const requestUrl: string =
      "http://localhost:8000/api/index/comments/" + articleId;
    return await (await fetch(requestUrl)).json();
  },

  async ArticlesCommentsCreate(requestData: commentAttributes) {
    const requestOptions = {
      method: "POST",
      mode: "no-cors",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(requestData),
    };
    const requestUrl: string = "http://localhost:8000/api/index/comments";

    fetch(requestUrl, requestOptions as RequestInit);
    return true;
  },

  async ArticlesDetail(id: number) {
    const requestUrl: string = "http://localhost:8000/api/index/articles/" + id;
    return await (await fetch(requestUrl)).json();
  },
};

export default api;
