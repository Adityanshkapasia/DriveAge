import { useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom";

type Post= {
    
        	id: number,
               title:   	string,
               desc: 	string,

               username:  string
        
        
}

export default function Viewallblog() {
    const [post, setpost] = useState<Post>()
    const { id } = useParams();
    const navigate=useNavigate()
    const onDelete=()=>{
        fetch("http://localhost:8080/delete/"+id, {
            credentials : 'include' ,
        }) .then (res=>{
            if (res.status=== 200){
                res.json().then(t=>{
                    navigate("/blog")
                })
            }
            else {
               
            }
        })
    }
    useEffect(() => {
      onLoad()
    
    }, [])

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
                 <button onClick={e=>onDelete()}>Delete</button>
                 -----------------
             </div>
             
        </div>
    )
}