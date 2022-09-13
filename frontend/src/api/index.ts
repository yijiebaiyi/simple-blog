export interface ApiMethodList {
  ArticlesList(): any;
  ArticlesDetail(id: number): any;
}

const api: ApiMethodList = {
  async ArticlesList() {
    const requestUrl: string = "http://localhost:8000/api/index/articles";
    return await (await fetch(requestUrl)).json();
  },

  async ArticlesDetail(id: number) {
    const requestUrl: string = "http://localhost:8000/api/index/articles/" + id;
    return await (await fetch(requestUrl)).json();
  },
};

export default api;
