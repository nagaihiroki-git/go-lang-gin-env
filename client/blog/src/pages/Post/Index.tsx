import React from 'react';
import '../..//App.scss';
import axios from "axios";
import {useParams} from "react-router-dom";
import ReactMarkdown from "react-markdown";

export const Post = () => {
    React.useEffect(() => {
        document.title = "記事ページ";
    });
    interface Post{
        title:string,
        body:string
    }
    const [post, setPost] = React.useState<Post>({
        title:"",
        body:""
    });
    const {id}=useParams();
    React.useEffect(()=>{
        axios.get<Post>(`http://localhost:3000/post?id=${id}`)
            .then(({data})=>{
                setPost(data)
            })
    },[])
    return (
        <>
            <div className={'index'}>
                <h1>TITLE::{post.title}</h1>
                <ReactMarkdown>{post.body}</ReactMarkdown>
            </div>
        </>
    );
}

