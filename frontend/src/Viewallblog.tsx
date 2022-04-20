import { useEffect, useState } from "react"

type Allposts= {
    
        	id: number,
               title:   	string,
               desc: 	string,

               username:  string
        
        
}

export default function Viewallblog() {
    const [posts, setposts] = useState<Allposts[]>([])
    const onLoad=()=>{
        fetch("http://localhost:8080/allpost", {
            credentials : 'include' ,
        }) .then (res=>{
            if (res.status=== 200){
                res.json().then(t=>{
                    setposts(t)
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
       {
           posts.map(post=>
               (
             <div key={post.id}>
                 <h1>
                     {post.title}
                 </h1>
                 author: {post.username}<br></br>
                 -----------------
             </div>
             
               )
           )
       }
        </div>
    )
}