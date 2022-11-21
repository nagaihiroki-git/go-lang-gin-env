import React from "react";
import axios from "axios";
import ReactMarkdown from "react-markdown";

export const CreatePost = () => {
  const [title, setTitle] = React.useState<string>("");
  const [body, setBody] = React.useState<string>("");

  const savePost = () => {
      const params = new URLSearchParams();
        params.append("title", title);
        params.append("body", body);
    axios.post("http://localhost:3000/create/post", params)
      .then((r) => alert("Post saved!"))
      .catch((e) => alert("Error!"));
  };
  return (
    <div>
      <h1>Create Post</h1>
      <div>TITLE</div>
      <input
        type="text"
        placeholder={"title"}
        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
          setTitle(event.target.value);
        }}
      />
      <div>BODY</div>
      <textarea
        placeholder={"body"}
        onChange={(event: React.ChangeEvent<HTMLTextAreaElement>) => {
          setBody(event.target.value);
        }}
      />
        <div>
            <ReactMarkdown>{body}</ReactMarkdown>
        </div>
      <button onClick={savePost}>Save</button>
    </div>
  );
};
