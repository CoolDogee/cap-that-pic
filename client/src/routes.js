import React from "react";
import { Redirect } from "react-router-dom";

import { CaptionsPage } from "./components/ChooseCaption";
import { PostPage } from "./components/Post";
import { AllPostsPage } from "./components/AllPosts";
import { FileUpload } from "./components/FileUpload";
import EmptyLayout from "./components/EmptyLayout";

export default [
    {
        path: "/",
        exact: true,
        layout: EmptyLayout,
        component: () => <Redirect to="/home" />
    },
    {
        path: "/home",
        layout: EmptyLayout,
        component: FileUpload
    },
    {
        path: "/choose-caption",
        layout: EmptyLayout,
        component: CaptionsPage
    },
    {
        path: "/i/:id",
        layout: EmptyLayout,
        component: PostPage
    },
    {
        path: "/posts",
        layout: EmptyLayout,
        component: AllPostsPage
    }
];