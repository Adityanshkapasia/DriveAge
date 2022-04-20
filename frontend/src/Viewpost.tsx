import { useEffect, useState } from "react"
import { useParams } from "react-router-dom";

type Post= {
    
        	id: number,
               title:   	string,
               desc: 	string,

               username:  string
        
        
}

export default function Viewallblog() {
    const [post, setpost] = useState<Post>()
    const { id } = useParams();
    const onLoad=()=>{
        fetch("http://localhost:8080/post/"+id, {
            credentials : 'include' ,
        }) .then (res=>{
            if (res.status=== 200){
                res.json().then(t=>{
                    setpost(t)
                })
            }
            else {
               
            }
        })
    }
    useEffect(() => {
      onLoad()
    
    }, [])
    
    return(
        <div>
      
               
             <div>
                 <h1>
                     {post&&post.title}
                 </h1>
                 <p>
                     {post && post.desc}
                 </p>
                 author: {post&&post.username}<br></br>
                 -----------------
             </div>
             
        </div>
    )
}