use diesel;

diesel::table! {
    moz_cookies (id) {
        id -> Nullable<Integer>,
        originAttributes -> Text,
        name -> Nullable<Text>,
        value -> Nullable<Text>,
        host -> Nullable<Text>,
        path -> Nullable<Text>,
        expiry -> Nullable<Integer>,
        lastAccessed -> Nullable<Integer>,
        creationTime -> Nullable<Integer>,
        isSecure -> Nullable<Integer>,
        isHttpOnly -> Nullable<Integer>,
        inBrowserElement -> Nullable<Integer>,
        sameSite -> Nullable<Integer>,
        schemeMap -> Nullable<Integer>,
        isPartitionedAttributeSet -> Nullable<Integer>,
    }
}
