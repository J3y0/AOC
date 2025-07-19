use diesel::prelude::Queryable;

#[allow(non_snake_case, unused)]
#[derive(Debug, Queryable)]
pub struct Cookies {
    pub id: Option<i32>,
    pub originAttributes: String,
    pub name: Option<String>,
    pub value: Option<String>,
    pub host: Option<String>,
    pub path: Option<String>,
    pub expiry: Option<i32>,
    pub lastAccessed: Option<i32>,
    pub creationTime: Option<i32>,
    pub isSecure: Option<i32>,
    pub isHttpOnly: Option<i32>,
    pub inBrowserElement: Option<i32>,
    pub sameSite: Option<i32>,
    pub schemeMap: Option<i32>,
    pub isPartitionedAttributeSet: Option<i32>,
}
