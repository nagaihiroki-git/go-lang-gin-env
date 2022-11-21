import {BrowserRouter, Route, Routes} from "react-router-dom";
import {Home} from "../pages/Home/Index";
import {Post} from "../pages/Post/Index";
import {CreatePost} from "../pages/CreatePost/Index";

export const Router=()=>{
    return(
        <BrowserRouter>
            <Routes>
                <Route index element={<Home/>}/>
                <Route path="post/:id"  element={<Post/>} />
                <Route path="create_post" element={<CreatePost/>}/>
            </Routes>
        </BrowserRouter>
    )
}