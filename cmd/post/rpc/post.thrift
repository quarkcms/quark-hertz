namespace go post

struct Post {
    1: i64 id (vt.gt = "0")
    2: string title
    3: string name
    4: i64 category_id
    5: string content
}

struct PageRequest {
    1: i64 id (vt.gt = "0")
    2: optional string name
}

struct PageResponse {
    1: i64 id
}

struct ArticleDetailRequest {
    1: i64 id (vt.gt = "0")
    2: optional string name
}

struct ArticleDetailResponse {
    1: i64 id
}

struct ArticleListRequest {
    1: optional string search
    2: i64 page (vt.ge = "1")
    3: i64 page_size (vt.ge = "0")
    4: string order
    5: i64 category_id (vt.ge = "0")
}

struct ArticleListResponse {
    1: list<Post> items
    2: i64 total
}

// 内容服务
service PostService {

    // 获取单页内容
    PageResponse GetPage(1: PageRequest req)
  
    // 获取文章详情
    ArticleDetailResponse GetArticleDetail(1: ArticleDetailRequest req)

    // 获取文章列表
    ArticleListResponse GetArticleList(1: ArticleListRequest req)
}